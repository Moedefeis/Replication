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
	operationHashSet map[string]*Operation
	executionOrder   chan *proto.OperationId
	newOperation     chan bool
}

type Leader struct {
	broadcast          sync.Mutex
	broadcasted        map[*Operation]bool
	broadcastOperation chan *Operation
}

type Operation struct {
	opid    *proto.OperationId
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
		operationHashSet: make(map[string]*Operation),
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
	op := newOperation(bid.Opid)
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

func (a *Auction) Result(ctx context.Context, opid *proto.OperationId) (*proto.Amount, error) {
	op := newOperation(opid)
	node.addOperation(op)
	<-op.execute
	highestBid := a.highestBid
	op.done <- true
	return &proto.Amount{Amount: highestBid}, nil
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
	log.Printf("Received opid: %v", opid.Id)
	n.executionOrder <- opid
	return &proto.Void{}, nil
}

func (l *Leader) broadcastOperationExecutionOrder() {
	for {
		op := <-l.broadcastOperation
		l.broadcast.Lock()
		if _, contains := l.broadcasted[op]; !contains {
			log.Printf("Broadcasting opid: %v", op.opid.Id)
			l.broadcasted[op] = true
			for _, conn := range conns {
				conn.AddOperationToExecutionOrder(ctx, op.opid)
			}
			node.AddOperationToExecutionOrder(ctx, op.opid)
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
		op := n.getOperation(opid)
		if isLeader() {
			delete(leader.broadcasted, op)
		}
		op.execute <- true
		log.Printf("Executing %v", opid.Id)
		<-op.done
	}
}

func (n *ServerNode) getOperation(opid *proto.OperationId) *Operation {
	for {
		n.operations.Lock()
		op, contains := n.operationHashSet[opid.Id]
		if contains {
			n.operationOrder.Remove(op)
			delete(n.operationHashSet, opid.Id)
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
	n.operations.Lock()
	n.operationOrder.Add(op)
	n.operationHashSet[op.opid.Id] = op
	n.operations.Unlock()
	if len(n.newOperation) == 0 {
		n.newOperation <- true
	}
}

func (a *Auction) StartAuction(ctx context.Context, startTime *timestamp.Timestamp) (*proto.Void, error) {
	endTime := startTime.AsTime().Add(time.Minute)
	go a.endAuctionAt(&endTime)
	return &proto.Void{}, nil
}

func (a *Auction) endAuctionAt(endTime *time.Time) {
	duration := endTime.Sub(time.Now())
	time.Sleep(duration)
	opid := &proto.OperationId{Id: "end auction"}
	op := newOperation(opid)
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

func newOperation(opid *proto.OperationId) *Operation {
	log.Printf("Received op: %v", opid.Id)
	return &Operation{
		opid:    opid,
		execute: make(chan bool),
		done:    make(chan bool),
	}
}

func equal(op1 *proto.OperationId, op2 *proto.OperationId) bool {
	return op1.Id == op2.Id
}
