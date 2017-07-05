package main


import (
    "flag"
    "fmt"
    "git.apache.org/thrift.git/lib/go/thrift"
)

func main() {

    server := flag.Bool("server", false, "if its the server or client")
    serverId := flag.Int("serverId", 0, "wich endpoint is the server")
    addr := []string{"localhost:9090", "localhost:9080", "localhost:9070"}
    secure := false

    flag.Parse()

    var protocolFactory thrift.TProtocolFactory
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

    var transportFactory thrift.TTransportFactory
    transportFactory = thrift.NewTBufferedTransportFactory(8192)

     if *server {
            if err := runServer(transportFactory, protocolFactory, addr[*serverId], secure, *serverId); err != nil {
            fmt.Println("error running server:", err)
        }
    } else {
        if err := runClient(transportFactory, protocolFactory, addr[*serverId], secure, handleClient); err != nil {
            fmt.Println("error running client:", err)
        }
    }
}