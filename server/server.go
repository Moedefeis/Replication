package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	proto "github.com/Moedefeis/Replication/grpc"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/goombaio/orderedset"
	"google.golang.org/grpc"
)

type Auction struct {
	proto.UnimplementedAuctionServer

	highestBid int32
	ended      bool
}

type ServerNode struct {
	proto.UnimplementedServerNodeServer

	election         sync.Mutex
	operations       sync.Mutex
	operationOrder   *orderedset.OrderedSet
	operationHashSet map[string]chan *Operation
	executionOrder   chan *proto.OperationId
	newOperation     chan bool
}

type Leader struct {
	broadcast          sync.Mutex
	broadcasted        map[*Operation]bool
	broadcastOperation chan *Operation
}

type Operation struct {
	id      *proto.OperationId
	execute chan bool
	done    chan bool
}

var ctx = context.Background()
var node *ServerNode
var ownPort int
var leaderPort int
var conns = map[int]proto.ServerNodeClient{5000: nil, 5001: nil}

var leader *Leader

func main() {
	ownPort, _ = strconv.Atoi(os.Args[1])
	delete(conns, ownPort)
	election()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", ownPort, err)
	}

	defer listener.Close()
	node = &ServerNode{
		operationOrder:   orderedset.NewOrderedSet(),
		operationHashSet: make(map[string]chan *Operation),
		executionOrder:   make(chan *proto.OperationId, 1000),
		newOperation:     make(chan bool, 10),
	}
	a := &Auction{highestBid: 0, ended: false}

	server := grpc.NewServer()
	proto.RegisterAuctionServer(server, a)
	proto.RegisterServerNodeServer(server, node)

	go node.executePendingOperations()
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (a *Auction) Bid(ctx context.Context, bid *proto.Amount) (*proto.Response, error) {
	op := &Operation{
		id:      bid.Opid,
		execute: make(chan bool),
		done:    make(chan bool),
	}
	node.addOperation(op)
	<-op.execute
	var status bool
	var message string
	if a.ended {
		status = false
		message = "auction has ended"
	} else if a.highestBid < bid.Amount {
		a.highestBid = bid.Amount
		status = true
	} else {
		status = false
		message = "higher bid exists"
	}
	op.done <- true
	return &proto.Response{Status: status, Message: message}, nil
}

func (a *Auction) Result(ctx context.Context, void *proto.Void) (*proto.Amount, error) {
	return &proto.Amount{Amount: a.highestBid}, nil
}

func (a *Auction) Crashed(ctx context.Context, id *proto.ServerId) (*proto.Void, error) {
	port := int(id.Port)
	delete(conns, port)
	node.election.Lock()
	if port == leaderPort {
		election()
		if isLeader() {
			leader.broadcastExecutionOfAllPendingOperations()
		}
	}
	node.election.Unlock()
	return &proto.Void{}, nil
}

func (n *ServerNode) AddOperationToExecutionOrder(ctx context.Context, opid *proto.OperationId) (*proto.Void, error) {
	log.Printf("Notice to execute operation: %v", opid.Id)
	n.executionOrder <- opid
	return &proto.Void{}, nil
}

func (l *Leader) broadcastOperationExecutionOrder() {
	for {
		op := <-l.broadcastOperation
		l.broadcast.Lock()
		if _, contains := l.broadcasted[op]; !contains {
			log.Printf("Broadcasting operation: %v", op.id.Id)
			l.broadcasted[op] = true
			for _, conn := range conns {
				conn.AddOperationToExecutionOrder(ctx, op.id)
			}
			node.AddOperationToExecutionOrder(ctx, op.id)
		}
		l.broadcast.Unlock()
	}
}

func (l *Leader) broadcastExecutionOfAllPendingOperations() {
	log.Printf("Broadcasting order of all pending operations")
	for _, op := range node.operationOrder.Values() {
		l.broadcastOperation <- op.(*Operation)
	}
}

func (n *ServerNode) executePendingOperations() {
	for {
		opid := <-n.executionOrder
		log.Printf("A1")
		op := n.getOperation(opid)
		log.Printf("A6")
		if isLeader() {
			delete(leader.broadcasted, op)
		}
		op.execute <- true
		<-op.done
	}
}

func (n *ServerNode) getOperation(opid *proto.OperationId) *Operation {

	for {
		//log.Printf("A2")
		n.operations.Lock()
		log.Printf("A3")
		opChan, contains := n.operationHashSet[opid.Id]
		if contains {
			//log.Printf("A4")
			op := <-opChan
			log.Printf("A5")
			if len(opChan) == 0 {
				delete(n.operationHashSet, opid.Id)
			}
			n.operationOrder.Remove(op)
			n.operations.Unlock()
			return op
		}
		n.operations.Unlock()
		<-n.newOperation
	}
}

func (n *ServerNode) addOperation(op *Operation) {
	if isLeader() {
		leader.broadcastOperation <- op
	}
	//log.Printf("B1")
	n.operations.Lock()
	//log.Printf("B2")
	n.operationOrder.Add(op)
	//log.Printf("B3")
	opChan, contains := n.operationHashSet[op.id.Id]
	if !contains {
		n.operationHashSet[op.id.Id] = make(chan *Operation, 2)
	} else if len(opChan) == cap(opChan) {
		temp := make(chan *Operation, 2*cap(opChan))
		for len(opChan) > 0 {
			temp <- <-opChan
		}
		n.operationHashSet[op.id.Id] = temp
	}
	n.operationHashSet[op.id.Id] <- op
	log.Printf("B4")
	n.operations.Unlock()
	if len(n.newOperation) < 1 {
		n.newOperation <- true
	}
	log.Printf("B5")
}

func (a *Auction) StartAuction(ctx context.Context, startTime *timestamp.Timestamp) (*proto.Void, error) {
	endTime := startTime.AsTime().Add(time.Minute)
	go a.endAuctionAt(&endTime)
	return &proto.Void{}, nil
}

func (a *Auction) endAuctionAt(endTime *time.Time) {
	duration := endTime.Sub(time.Now())
	time.Sleep(duration)
	op := &Operation{
		id:      &proto.OperationId{Id: "end auction"},
		execute: make(chan bool),
		done:    make(chan bool),
	}
	node.addOperation(op)
	<-op.execute
	a.ended = true
	op.done <- true
}

func isLeader() bool {
	return ownPort == leaderPort
}

func election() {
	log.Printf("Electing leader")
	leaderPort = ownPort
	for port := range conns {
		if port < leaderPort {
			leaderPort = port
		}
	}

	if isLeader() {
		for port := range conns {
			conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Could not connect: %v", err)
			}
			conns[port] = proto.NewServerNodeClient(conn)
		}
		leader = &Leader{
			broadcasted:        make(map[*Operation]bool),
			broadcastOperation: make(chan *Operation, 100),
		}
		go leader.broadcastOperationExecutionOrder()
	}
}

/* func remove(slice []*Operation, op *Operation) []*Operation {
	make()
	for i, _op := range slice {
		if _op == op {
			break
		}
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
} */

func equal(op1 *proto.OperationId, op2 *proto.OperationId) bool {
	return op1.Id == op2.Id
}
