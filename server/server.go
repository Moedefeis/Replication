package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	proto "github.com/Moedefeis/Replication/grpc"
	"google.golang.org/grpc"
)

type Auction struct {
	proto.UnimplementedAuctionServer

	highestBid int32
	bidLock    sync.Mutex
}

type ServerNode struct {
	proto.UnimplementedServerNodeServer

	lock       sync.Mutex
	operations []Operation
}

type Operation struct {
	id string
	op chan bool
}

var ctx context.Context = context.Background()
var node *ServerNode
var ownPort int
var leaderPort int
var conns = map[int]proto.ServerNodeClient{5000: nil, 5001: nil, 5002: nil}

func main() {
	ownPort, _ = strconv.Atoi(os.Args[1])
	delete(conns, ownPort)
	election()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", ownPort, err)
	}

	defer listener.Close()
	node = &ServerNode{operations: make([]Operation, 0)}
	a := &Auction{highestBid: 0}

	server := grpc.NewServer()
	proto.RegisterAuctionServer(server, a)
	proto.RegisterServerNodeServer(server, node)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (a *Auction) Bid(ctx context.Context, bid *proto.Amount) (*proto.Response, error) {
	op := make(chan bool)
	node.lock.Lock()
	node.operations = append(node.operations, Operation{
		id: bid.Op.Id,
		op: op,
	})
	node.lock.Unlock()
	log.Printf("Received bid operation, waiting for leader's notice")
	if isLeader() {
		go node.executeOperations()
	}
	<-op
	a.bidLock.Lock()
	var status bool
	if a.highestBid < bid.Amount {
		a.highestBid = bid.Amount
		status = true
	} else {
		status = false
	}
	a.bidLock.Unlock()
	return &proto.Response{Status: status}, nil
}

func (a *Auction) Result(ctx context.Context, void *proto.Void) (*proto.Amount, error) {
	return &proto.Amount{Amount: a.highestBid}, nil
}

func (a *Auction) Crashed(ctx context.Context, id *proto.ServerId) (*proto.Void, error) {
	port := int(id.Port)
	delete(conns, port)
	if port == leaderPort {
		election()
		if isLeader() {
			node.executeOperations()
		}
	}
	return &proto.Void{}, nil
}

func (n *ServerNode) ExecuteOperation(ctx context.Context, opid *proto.OperationId) (*proto.Void, error) {
	log.Printf("Received notice to execute operation: %v", opid.Id)
	var removed Operation
	n.operations, removed = remove(n.operations, opid.Id)
	removed.op <- true
	return &proto.Void{}, nil
}

func (n *ServerNode) executeOperations() {
	n.lock.Lock()
	log.Printf("Broadcasting operations to execute")
	for len(n.operations) > 0 {
		op := n.operations[0]
		opId := &proto.OperationId{Id: op.id}
		for _, conn := range conns {
			conn.ExecuteOperation(ctx, opId)
		}
		n.ExecuteOperation(ctx, opId)
	}
	n.lock.Unlock()
	log.Printf("Broadcasting operations to execute - done")
}

func isLeader() bool {
	return ownPort == leaderPort
}

func election() {
	log.Printf("Electing leader")
	leaderPort = ownPort
	for port := range conns {
		if port > leaderPort {
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

func remove(slice []Operation, id string) ([]Operation, Operation) {
	var i int
	var op Operation
	for i, op = range slice {
		if op.id == id {
			break
		}
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1], op
}
