package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	proto "github.com/Moedefeis/Replication/grpc"
	"google.golang.org/grpc"
)

var conns = make(map[int]proto.AuctionClient)
var reader *bufio.Scanner
var ctx = context.Background()

func main() {
	reader = bufio.NewScanner(os.Stdin)
	ports := []int{5000, 5001, 5002}
	for _, port := range ports {
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		conns[port] = proto.NewAuctionClient(conn)
	}
	//stressTest()
	handleInput()
}

func stressTest() {
	for i := 0; i <= 200; i++ {
		go bid(i)
	}
	<-make(chan bool)
}

func handleInput() {
	for {
		reader.Scan()
		input := strings.ToLower(reader.Text())
		var subInput string = "Not a number"
		if len(input) > 4 {
			subInput = input[4:]
		}
		if amount, err := strconv.Atoi(subInput); err == nil && strings.Contains(input, "bid") {
			bid(amount)
		} else if strings.Contains(input, "result") {
			queryResult()
		} else {
			log.Printf("Invaild input")
		}
	}
}

func bid(amount int) {
	bid := &proto.Amount{
		Amount: int32(amount),
		Opid:   &proto.OperationId{Id: fmt.Sprintf("Bid %d", amount)},
	}
	var wg sync.WaitGroup
	wrapper := &response{}
	for port, conn := range conns {
		wg.Add(1)
		go bidHandler(port, conn, bid, &wg, wrapper)
	}
	wg.Wait()
	if wrapper.reponse.Status {
		log.Printf("Bidded %d successfully", amount)
	} else {
		log.Printf("Bidded %d unsuccessfully, as higher bid existed", amount)
	}
}

func bidHandler(port int, conn proto.AuctionClient, bid *proto.Amount, wg *sync.WaitGroup, wrapper *response) {
	defer wg.Done()

	response, err := conn.Bid(ctx, bid)
	if err != nil {
		log.Printf(err.Error())
		handleCrashedServer(port)
	} else {
		wrapper.reponse = response
	}
}

func queryResult() {
	var result int32 = 0
	for port, conn := range conns {
		highestBid, err := conn.Result(ctx, &proto.Void{})
		if err != nil {
			handleCrashedServer(port)
		} else if highestBid.Amount > result {
			result = highestBid.Amount
		}
	}
	log.Printf("Highest bid is: %d", result)
}

func handleCrashedServer(port int) {
	if _, contains := conns[port]; contains {
		log.Printf("Server at port: %d crashed", port)
		delete(conns, port)
		serverid := &proto.ServerId{Port: int32(port)}
		for _, conn := range conns {
			go conn.Crashed(ctx, serverid)
		}
	}
}

type response struct {
	reponse *proto.Response
}
