package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	auction "github.com/Moedefeis/Replication/grpc"
	"google.golang.org/grpc"
)

type Auction struct {
	auction.UnimplementedAuctionServer

	highestBid int32
	bidLock    sync.Mutex
}

func main() {
	port, _ := strconv.Atoi(os.Args[1])

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	defer listener.Close()
	a := &Auction{
		highestBid: 0,
	}

	server := grpc.NewServer()
	auction.RegisterAuctionServer(server, a)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (a *Auction) Bid(ctx context.Context, bid *auction.Amount) (*auction.Response, error) {
	var status bool
	a.bidLock.Lock()
	if a.highestBid < bid.Amount {
		a.highestBid = bid.Amount
		status = true
	} else {
		status = false
	}
	a.bidLock.Unlock()
	return &auction.Response{Status: status}, nil
}

func (a *Auction) Result(ctx context.Context, void *auction.Void) (*auction.Amount, error) {
	return &auction.Amount{Amount: a.highestBid}, nil
}
