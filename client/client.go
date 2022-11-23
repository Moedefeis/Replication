package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	auction "github.com/Moedefeis/Replication/grpc"
	"google.golang.org/grpc"
)

var conns = make(map[int]auction.AuctionClient)
var reader *bufio.Scanner
var ctx = context.Background()

func main() {
	reader = bufio.NewScanner(os.Stdin)

	for i := 1; i < len(os.Args); i++ {
		port, _ := strconv.Atoi(os.Args[i])

		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		conns[port] = auction.NewAuctionClient(conn)
	}

	handleInput()
}

func handleInput() {
	for {
		reader.Scan()
		input := strings.ToLower(reader.Text())
		if amount, err := strconv.Atoi(input[4:]); err != nil && strings.Contains(input, "bid") {
			bid(amount)
		} else if strings.Contains(input, "result") {
			queryResult()
		}
	}
}

func bid(amount int) {
	success := true
	bid := &auction.Amount{Amount: int32(amount)}
	for port, conn := range conns {
		response, err := conn.Bid(ctx, bid)
		if err != nil {
			handleCrashedServer(port)
		} else {
			success = success && response.Status
		}
	}
	if success {
		log.Printf("Bidded %d successfully", amount)
	} else {
		log.Printf("Bidded %d unsuccessfully, as higher bid existed", amount)
	}
}

func queryResult() {
	var result int32 = 0
	for port, conn := range conns {
		highestBid, err := conn.Result(ctx, &auction.Void{})
		if err != nil {
			handleCrashedServer(port)
		} else if highestBid.Amount > result {
			result = highestBid.Amount
		}
	}
	log.Printf("Highest bid is: %d", result)
}

func handleCrashedServer(port int) {
	log.Printf("Server at port: %d crashed", port)
	delete(conns, port)
}
