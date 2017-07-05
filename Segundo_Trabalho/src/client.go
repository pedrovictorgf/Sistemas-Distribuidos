package main

import (
    "crypto/tls"
    "fmt"
    "git.apache.org/thrift.git/lib/go/thrift"
    "SD/Segundo_Trabalho/gen-go/graphdb"
)

type ClientOperation func (client *graphdb.GraphCRUDClient) (err error)

func handleClient(client *graphdb.GraphCRUDClient) (err error) {
    /*v := graphdb.NewGraphVertex()
    v.Name = 7
    v.Description = "setimo vertice"
    if e := client.CreateVertex(v); e == nil {
        fmt.Println("Inseriu")
    } else {
        fmt.Println("Não inseriu")
    }

    if v1, err := client.ReadVertex(3); err == nil {
        PrintVertex(v1)
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }

    if v1, err := client.ReadVertex(4); err == nil {
        PrintVertex(v1)
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }*/

    // if err := client.DeleteVertex(1); err == nil {
    //     fmt.Println("Deleção efetuada com sucesso")
    // } else {
    //     fmt.Print("Erro:")
    //     fmt.Println(err)
    // }

    /*v2 := graphdb.NewGraphVertex()
    v2.Name = 5
    v2.Color = 5
    v2.Description = "alterado"
    v2.Weight = 10.0
    if err:= client.UpdateVertex(v2); err == nil {
        fmt.Println("Update efetuado com sucesso")
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }

    v3 := graphdb.NewGraphVertex()
    v4 := graphdb.NewGraphVertex()
    v3.Name = 1
    v3.Color = 1
    v3.Weight = 10.0
    v3.Description = "primeiro vertice"
    v4.Name = 2
    v4.Color = 2
    v4.Weight = 20.0
    v4.Description = "segundo vertice"
    e1 := graphdb.NewGraphEdge()
    e1.FirstVertex = v3
    e1.SecondVertex = v4
    e1.Weight = 15.0
    e1.IsBidirectional = true
    e1.Description = "aresta 1-2"

    if err:= client.CreateEdge(e1); err == nil {
        fmt.Println("Aresta criada com sucesso")
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }*/

    /*var edgeToDelete *graphdb.GraphEdge
    if ee, err := client.ReadEdge(1,2); err == nil {
        fmt.Println("Aresta encontrada:")
        PrintEdge(ee)
        edgeToDelete = ee
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }

    edgeToDelete.Description = "alterada"
    if err := client.UpdateEdge(edgeToDelete); err == nil {
        fmt.Println("Update aresta com sucesso")
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }*/

    /*if edges, err := client.FindEdgesOfVertex(1); err == nil {
        fmt.Println("Procura arestas vértice sucesso")
        PrintEdges(edges)
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }

    if vertices, err := client.FindNeighbours(1); err == nil {
        fmt.Println("Vizinhos do vértice")
        PrintVertices(vertices)
    } else {
        fmt.Print("Erro:", )
        fmt.Println(err)
    }*/

    if d, err := client.ShortestPath(6, 2); err == nil {
        fmt.Println(d)
    }else {
        fmt.Print("Erro:", )
        fmt.Println(err)
    }

    return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, fn ClientOperation) error {
    var transport thrift.TTransport
    var err error
    if secure {
        cfg := new(tls.Config)
        cfg.InsecureSkipVerify = true
        transport, err = thrift.NewTSSLSocket(addr, cfg)
    } else {
        transport, err = thrift.NewTSocket(addr)
    }
    if err != nil {
        fmt.Println("Error opening socket:", err)
        return err
    }
    transport, _ = transportFactory.GetTransport(transport)
    defer transport.Close()
    if err := transport.Open(); err != nil {
        return err
    }

    fmt.Println("Starting the client... on ", addr)
    // go func(){

    return fn(graphdb.NewGraphCRUDClientFactory(transport, protocolFactory))

}

func getClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) (*graphdb.GraphCRUDClient, error) {
    var transport thrift.TTransport
    var err error
    if secure {
        cfg := new(tls.Config)
        cfg.InsecureSkipVerify = true
        transport, err = thrift.NewTSSLSocket(addr, cfg)
    } else {
        transport, err = thrift.NewTSocket(addr)
    }
    if err != nil {
        fmt.Println("Error opening socket:", err)
        return nil, err
    }
    transport, _ = transportFactory.GetTransport(transport)
    defer transport.Close()
    if err := transport.Open(); err != nil {
        return nil, err
    }

    fmt.Println("Starting the client... on ", addr)
    // go func(){

    return graphdb.NewGraphCRUDClientFactory(transport, protocolFactory), nil
}

func handler(client *graphdb.GraphCRUDClient) (err error) {
    fmt.Println("Pau no cu")

    if v1, err := client.ReadVertex(4); err == nil {
        PrintVertex(v1)
    } else {
        fmt.Print("Erro:")
        fmt.Println(err)
    }

    return nil
}