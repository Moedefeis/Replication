# Distributed Auction System with Replication

To run the program, first navigate to this directory, and start two servers by executing these commands:
    "go run server/server.go 5000"
    "go run server/server.go 5001"

Next run as many clients as you wish by executing the following command: 
    "go run client/client.go"

You can now use the command "bid *amount*" to bid to the auction, or the command "result" to see the currently highest bid. The auction ends one minute after the first client has joined.