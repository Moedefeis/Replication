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
	"time"

	proto "github.com/Moedefeis/Replication/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var messageAmountLock sync.Mutex
var messageAmount int = 0
var clientId string
var conns = make(map[int]proto.AuctionClient)
var reader *bufio.Scanner
var ctx = context.Background()

func main() {
	clientId = os.Args[1]
	reader = bufio.NewScanner(os.Stdin)
	ports := []int{5000, 5001}
	joinTime := timestamppb.New(time.Now())
	for _, port := range ports {
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		client := proto.NewAuctionClient(conn)
		conns[port] = client
		go client.StartAuction(ctx, joinTime)
	}
	//stressTest()
	handleInput()
}

func stressTest() {
	for i := 0; i <= 2000; i++ {
		go bid(i)
	}
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
		Opid:   &proto.OperationId{Id: fmt.Sprintf("bid %d", amount) + messageIdAffix()},
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
		log.Printf("Bidded %d unsuccessfully, %v", amount, wrapper.reponse.Message)
	}
}

func bidHandler(port int, conn proto.AuctionClient, bid *proto.Amount, wg *sync.WaitGroup, wrapper *response) {
	defer wg.Done()

	response, err := conn.Bid(ctx, bid)
	if err != nil {
		//log.Printf(err.Error())
		handleCrashedServer(port)
	} else {
		wrapper.reponse = response
	}
}

func queryResult() {
	var result int32 = 0
	opid := &proto.OperationId{Id: fmt.Sprintf("result") + messageIdAffix()}
	for port, conn := range conns {
		highestBid, err := conn.Result(ctx, opid)
		if err != nil {
			handleCrashedServer(port)
		} else {
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

func amount() int {
	messageAmountLock.Lock()
	amount := messageAmount
	messageAmount++
	messageAmountLock.Unlock()
	return amount
}

func messageIdAffix() string {
	return " " + clientId + " " + strconv.Itoa(amount())
}

type response struct {
	reponse *proto.Response
}
