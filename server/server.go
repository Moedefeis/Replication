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
	"google.golang.org/grpc"
)

type Auction struct {
	proto.UnimplementedAuctionServer

	highestBid int32
	bidLock    sync.Mutex
	ended      bool
}

type ServerNode struct {
	proto.UnimplementedServerNodeServer

	election          sync.Mutex
	operations        sync.Mutex
	pendingOperations []*Operation
	executionOrder    []*proto.OperationId
}

type Operation struct {
	id      *proto.OperationId
	execute chan bool
	done    chan bool
}

var broadcast sync.Mutex
var broadcasted = make(map[*Operation]bool)
var ctx context.Context = context.Background()
var node *ServerNode
var ownPort int
var leaderPort int
var conns = map[int]proto.ServerNodeClient{5000: nil, 5001: nil}

func main() {
	ownPort, _ = strconv.Atoi(os.Args[1])
	delete(conns, ownPort)
	election()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", ownPort, err)
	}

	defer listener.Close()
	node = &ServerNode{pendingOperations: make([]*Operation, 0)}
	a := &Auction{highestBid: 0, ended: false}

	server := grpc.NewServer()
	proto.RegisterAuctionServer(server, a)
	proto.RegisterServerNodeServer(server, node)

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
	node.operations.Lock()
	log.Printf("Received operation: %v", bid.Opid.Id)
	node.pendingOperations = append(node.pendingOperations, op)
	node.operations.Unlock()
	if isLeader() {
		go node.broadcastOperationExecution(op)
	} else {
		go node.executePendingOperations()
	}
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
			node.broadcastExecutionOfAllPendingOperations()
		}
	}
	node.election.Unlock()
	return &proto.Void{}, nil
}

func (n *ServerNode) AddOperationToExecutionOrder(ctx context.Context, opid *proto.OperationId) (*proto.Void, error) {
	n.operations.Lock()
	log.Printf("Notice to execute operation: %v", opid.Id)
	n.executionOrder = append(n.executionOrder, opid)
	n.operations.Unlock()
	go n.executePendingOperations()
	return &proto.Void{}, nil
}

func (n *ServerNode) broadcastExecutionOfAllPendingOperations() {
	log.Printf("Broadcasting order of all pending operations")
	operations := n.pendingOperations
	for _, op := range operations {
		n.broadcastOperationExecution(op)
	}
}

func (n *ServerNode) broadcastOperationExecution(op *Operation) {
	broadcast.Lock()
	if _, contains := broadcasted[op]; !contains {
		log.Printf("Broadcasting operation: %v", op.id.Id)
		broadcasted[op] = true
		for _, conn := range conns {
			conn.AddOperationToExecutionOrder(ctx, op.id)
		}
		n.AddOperationToExecutionOrder(ctx, op.id)
	}
	broadcast.Unlock()
}

func (n *ServerNode) executePendingOperations() {
	n.operations.Lock()
	hasNextOperation := true
	for len(n.executionOrder) > 0 && hasNextOperation {
		opid := n.executionOrder[0]
		hasNextOperation = false
		for i, op := range n.pendingOperations {
			if equal(opid, op.id) {
				hasNextOperation = true
				n.executionOrder = n.executionOrder[1:]
				n.pendingOperations = remove(n.pendingOperations, i)
				if isLeader() {
					delete(broadcasted, op)
				}
				op.execute <- true
				<-op.done
				break
			}
		}
	}
	n.operations.Unlock()
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
	node.operations.Lock()
	node.pendingOperations = append(node.pendingOperations, op)
	node.operations.Unlock()
	if isLeader() {
		go node.broadcastOperationExecution(op)
	}
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
	}
}

func remove(slice []*Operation, i int) []*Operation {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func equal(op1 *proto.OperationId, op2 *proto.OperationId) bool {
	return op1.Id == op2.Id
}
