package main

import (
    "crypto/tls"
    "fmt"
    "git.apache.org/thrift.git/lib/go/thrift"
    "SD/Segundo_Trabalho/gen-go/graphdb"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, serverId int) error {
    var transport thrift.TServerTransport
    var err error
    ipList := []string{"localhost:9090", "localhost:9080", "localhost:9070"}
    if secure {
        cfg := new(tls.Config)
        if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
            cfg.Certificates = append(cfg.Certificates, cert)
        } else {
            return err
        }
        transport, err = thrift.NewTSSLServerSocket(addr, cfg)
    } else {
        transport, err = thrift.NewTServerSocket(addr)
    }

    if err != nil {
        return err
    }
    fmt.Printf("%T\n", transport)
    handler := NewGraphHandler(Mock(serverId))
    handler.TransportFactory = transportFactory
    handler.ProtocolFactory = protocolFactory
    handler.IpList = ipList
    handler.NumServers = len(handler.IpList)
    handler.ServerId = serverId
    processor := graphdb.NewGraphCRUDProcessor(handler)
    server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

    fmt.Println("Starting the simple server... on ", addr)
    return server.Serve()
}

