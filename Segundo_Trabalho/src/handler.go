package main

import (
    "fmt"
    "errors"
    // "strconv"
    "git.apache.org/thrift.git/lib/go/thrift"
    "SD/Segundo_Trabalho/gen-go/graphdb"
)

type GraphHandler struct {
    Graph *graphdb.Graph
    ServerId int
    NumServers  int
    IpList  []string
    TransportFactory thrift.TTransportFactory
    ProtocolFactory thrift.TProtocolFactory
}

func NewGraphHandler(graph graphdb.Graph) *GraphHandler {
    return &GraphHandler{Graph: &graph}
}

func (p *GraphHandler) CreateVertex(v *graphdb.GraphVertex) (err error){
    if correctServer := int(v.Name) % p.NumServers; correctServer == p.ServerId {
        if !IsVertexInGraph(p.Graph, v.Name) {
            fmt.Println("Sucesso CreateVertex, server: ", p.ServerId)
            p.Graph.Vertices = InsertVertex(p.Graph.Vertices, v)
            PrintVertices(p.Graph.Vertices)
            return nil
        } else {
            return errors.New("Vértice já existe")
        }
    } else {
        err = runClient(p.TransportFactory, p.ProtocolFactory, p.IpList[correctServer], false, func(client *graphdb.GraphCRUDClient) (err error){
            return client.CreateVertex(v)
            })

        return err
    }
}

func (p *GraphHandler) ReadVertex(name graphdb.Int) (*graphdb.GraphVertex, error){
    if correctServer := int(name) % p.NumServers; correctServer == p.ServerId {
        if v, e := FetchVertex(p.Graph.Vertices, name); e == nil {
            fmt.Println("Sucesso no ReadVertex, server: ", p.ServerId)
            return v, e
        } else {
            return nil, e
        }
    }  else {
        var vv *graphdb.GraphVertex
        runClient(p.TransportFactory, p.ProtocolFactory, p.IpList[correctServer], false, func(client *graphdb.GraphCRUDClient) (err error){
                if v, err := client.ReadVertex(name); err == nil {
                    vv = v
                    return nil
                } else {
                    return err
                }
            })
        if vv != nil {
            return vv, nil
        } else {
            return nil, errors.New("Não foi possível encontrar o vértice")
        }
     }
}

func (p *GraphHandler) DeleteVertex(name graphdb.Int) (err error){
    if correctServer := int(name) % p.NumServers; correctServer == p.ServerId {
        if IsVertexInGraph(p.Graph, name) {
            fmt.Println("Sucesso na deleção")
            p.Graph = DeleteVertex(p.Graph, name)
            PrintVertices(p.Graph.Vertices)
            PrintEdges(p.Graph.Edges)
            return nil
        } else {
            return errors.New("Vértice não encontrado")
        }
    } else {
        err = runClient(p.TransportFactory, p.ProtocolFactory, p.IpList[correctServer], false, func(client *graphdb.GraphCRUDClient) (err error){
                return client.DeleteVertex(name)
            })

        if err == nil {
            p.Graph = DeleteEdgesOfVertex(p.Graph, name)
            PrintEdges(p.Graph.Edges)
            for i := 0; i < p.NumServers; i++ {
                if i != correctServer && i != p.ServerId {
                    addr := p.IpList[i]
                    go func(){
                        runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                            return client.DeleteEdgesOfVertex(name)
                        })
                    }()
                }
            }
        }
        return err
    }
}

func (p *GraphHandler) UpdateVertex(v *graphdb.GraphVertex) (err error){
    if correctServer := int(v.Name) % p.NumServers; correctServer == p.ServerId {
        if IsVertexInGraph(p.Graph, v.Name) {
            fmt.Println("Update sucesso")
            p.Graph.Vertices = UpdateVertex(p.Graph.Vertices, v)
            PrintVertices(p.Graph.Vertices)
            } else {
                return errors.New("Vértice não encontrado")
            }
        return nil
    } else {
        err = runClient(p.TransportFactory, p.ProtocolFactory, p.IpList[correctServer], false, func(client *graphdb.GraphCRUDClient) (err error){
                return client.UpdateVertex(v)
            })

        return err
    }
}

func (p *GraphHandler) CreateEdge(e *graphdb.GraphEdge) (err error){
   if correctServer := int(e.FirstVertex.Name) % p.NumServers; correctServer == p.ServerId {
        if (IsVertexInGraph(p.Graph, e.FirstVertex.Name) || IsVertexInGraph(p.Graph, e.SecondVertex.Name)) && !isEdgeInGraph(p.Graph, e){
            fmt.Println("Sucesso criaçao aresta")
            p.Graph.Edges = InsertEdge(p.Graph.Edges, e)
            PrintEdges(p.Graph.Edges)
        } else {
            return errors.New("Um dos vértices não existe no grafo, ou aresta já existe")
        }
        return nil
    } else if correctServer = int(e.SecondVertex.Name) % p.NumServers; e.IsBidirectional && correctServer == p.ServerId {
        if (IsVertexInGraph(p.Graph, e.FirstVertex.Name) || IsVertexInGraph(p.Graph, e.SecondVertex.Name)) && !isEdgeInGraph(p.Graph, e){
            fmt.Println("Sucesso criaçao aresta")
            p.Graph.Edges = InsertEdge(p.Graph.Edges, e)
            PrintEdges(p.Graph.Edges)
        } else {
            return errors.New("Um dos vértices não existe no grafo, ou aresta já existe")
        }
        return nil
    } else {
        for i := 0; i < p.NumServers; i++ {
            if i != p.ServerId {
                addr := p.IpList[i]
                go func(){
                    err = runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                        return client.CreateEdge(e)
                    })
                }()
            }
        }

        return err
    }
}

func (p *GraphHandler) ReadEdge(firstVertexName  graphdb.Int, secondVertexName  graphdb.Int) (*graphdb.GraphEdge,  error) {
    if correctServer := int(firstVertexName) % p.NumServers; correctServer == p.ServerId {
        fmt.Println("Aresta lida com sucesso")
        if e, err := FetchEdge(p.Graph.Edges ,firstVertexName, secondVertexName); err == nil {
            return e, nil
        } else {
            return nil, err
        }
    } else if correctServer := int(secondVertexName) % p.NumServers; correctServer == p.ServerId {
           fmt.Println("Aresta lida com sucesso")
            if e, err := FetchEdge(p.Graph.Edges ,firstVertexName, secondVertexName); err == nil {
                return e, nil
            } else {
                return nil, err
            }
        } else {
            correctServer = int(firstVertexName) % p.NumServers
            var ee *graphdb.GraphEdge
            err := runClient(p.TransportFactory, p.ProtocolFactory, p.IpList[correctServer], false, func(client *graphdb.GraphCRUDClient) (err error){
                if e, err := client.ReadEdge(firstVertexName, secondVertexName); err == nil {
                    ee = e
                    return nil
                } else {
                    return err
                }
            })

            if ee != nil {
                return ee, nil
            } else {
                return nil, err
            }
        }
}

func (p *GraphHandler) DeleteEdge(e *graphdb.GraphEdge) (err error){
    if correctServer := int(e.FirstVertex.Name) % p.NumServers; correctServer == p.ServerId{

        if IsVertexInGraph(p.Graph, e.FirstVertex.Name) && isEdgeInGraph(p.Graph, e) {
            fmt.Println("Remoção aresta bem sucedida")
            p.Graph.Edges = DeleteEdge(p.Graph.Edges, e)
            PrintEdges(p.Graph.Edges)

            if e.IsBidirectional {
                secondServer := int(e.SecondVertex.Name) % p.NumServers
                addr := p.IpList[secondServer]

                go func(){
                    runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                        return client.DeleteEdgeColateral(e)
                    })
                }()
            }
            return nil

        } else {
            return errors.New("Aresta não existe")
            }

    } else if correctServer = int(e.SecondVertex.Name) % p.NumServers; e.IsBidirectional && correctServer == p.ServerId {

        if IsVertexInGraph(p.Graph, e.SecondVertex.Name) && isEdgeInGraph(p.Graph, e) {

            fmt.Println("Remoção aresta bem sucedida")
            p.Graph.Edges = DeleteEdge(p.Graph.Edges, e)
            PrintEdges(p.Graph.Edges)

            secondServer := int(e.FirstVertex.Name) % p.NumServers
            addr := p.IpList[secondServer]

            go func(){
                runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                    return client.DeleteEdgeColateral(e)
                })
            }()

            return nil
        } else {
            return errors.New("Aresta não existe")
            }
        } else {
            for i := 0; i < p.NumServers; i++ {
                if i != p.ServerId {
                    addr := p.IpList[i]
                    go func(){
                        err = runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                            return client.DeleteEdge(e)
                        })
                    }()
                }
            }

            return err
        }
}

func (p *GraphHandler) UpdateEdge(e *graphdb.GraphEdge) (err error) {
    if correctServer := int(e.FirstVertex.Name) % p.NumServers; correctServer == p.ServerId {
        if IsVertexInGraph(p.Graph, e.FirstVertex.Name) && isEdgeInGraph(p.Graph, e) {
            fmt.Println("Update aresta bem sucedido")
            p.Graph.Edges = UpdateEdge(p.Graph.Edges, e)
            PrintEdges(p.Graph.Edges)

            if e.IsBidirectional {
                secondServer := int(e.SecondVertex.Name) % p.NumServers
                addr := p.IpList[secondServer]

                go func(){
                err = runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                        return client.UpdateEdgeColateral(e)
                    })
                }()
            }

            return nil
        } else {
                return errors.New("Aresta não existe")
        }
    } else if correctServer = int(e.SecondVertex.Name) % p.NumServers; e.IsBidirectional && correctServer == p.ServerId {
            if IsVertexInGraph(p.Graph, e.SecondVertex.Name) && isEdgeInGraph(p.Graph, e) {
                fmt.Println("Update aresta bem sucedido")
                p.Graph.Edges = UpdateEdge(p.Graph.Edges, e)
                PrintEdges(p.Graph.Edges)

                secondServer := int(e.FirstVertex.Name) % p.NumServers
                addr := p.IpList[secondServer]

                go func(){
                err = runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                        return client.UpdateEdgeColateral(e)
                    })
                }()

                return nil
        } else  {
            return errors.New("Aresta não existe")
        }
    } else {
        for i := 0; i < p.NumServers; i++ {
            if i != p.ServerId {
                addr := p.IpList[i]
                go func(){
                    err = runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                        return client.UpdateEdge(e)
                    })
                }()
            }
        }

        return err
    }
}

func (p *GraphHandler) FindEdgesOfVertex(name graphdb.Int) ([]*graphdb.GraphEdge,  error){
    if correctServer := int(name) % p.NumServers; correctServer == p.ServerId {
        if IsVertexInGraph(p.Graph, name) {
            fmt.Println("Procura arestas vertice sucesso")
            edgesOfVertex := FindEdgesOfVertex(p.Graph.Edges, name)
            return edgesOfVertex, nil
        } else {
                return nil, errors.New("Vértice não existe")
        }
    } else {
        var edges []*graphdb.GraphEdge
        runClient(p.TransportFactory, p.ProtocolFactory, p.IpList[correctServer], false, func(client *graphdb.GraphCRUDClient) (err error){
                if e, err := client.FindEdgesOfVertex(name); err == nil {
                    edges = e
                    return nil
                } else {
                    return err
                }
            })

        if edges != nil {
            return edges, nil
        } else {
            return nil, errors.New("Vértice não existe, ou não possui arestas")
        }
    }
}

func (p *GraphHandler) FindNeighbours(name graphdb.Int) ([]*graphdb.GraphVertex, error){
    fmt.Println("Vai procurar vizinhos")
    neighbours := FindNeighbours(p.Graph.Edges, name)
    for i := 0; i < p.NumServers; i++ {
        if i != p.ServerId {
            addr := p.IpList[i]
            runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                n, err := client.FindNeighboursRemote(name)
                neighbours = append(neighbours, n...)
                return err
            })
        }
    }

    neighbours = removeDuplicatesVertices(neighbours)
    if neighbours != nil {
        return neighbours, nil
    } else {
        return neighbours, errors.New("Vizinhos não encontrados")
    }
}

func (p *GraphHandler) DeleteEdgesOfVertex(name graphdb.Int) (err error) {
    p.Graph = DeleteEdgesOfVertex(p.Graph, name)
    PrintEdges(p.Graph.Edges)
    return nil
}

func (p *GraphHandler) DeleteEdgeColateral(e *graphdb.GraphEdge) (err error) {
    fmt.Println("Remoção aresta bem sucedida")
    p.Graph.Edges = DeleteEdge(p.Graph.Edges, e)
    PrintEdges(p.Graph.Edges)
    return nil
}

func (p *GraphHandler) UpdateEdgeColateral(e *graphdb.GraphEdge) (err error) {
    fmt.Println("Update aresta bem sucedido")
    p.Graph.Edges = UpdateEdge(p.Graph.Edges, e)
    PrintEdges(p.Graph.Edges)
    return nil
}

func (p *GraphHandler) FindNeighboursRemote(name graphdb.Int) ([]*graphdb.GraphVertex, error) {
    fmt.Println("Vai procurar vizinhos")
    if neighbours := FindNeighbours(p.Graph.Edges, name); neighbours != nil {
        return neighbours, nil
    } else {
        return nil, errors.New("Vizinhos não encontrados")
    }
}

func (p *GraphHandler) GetGraph() (*graphdb.Graph,  error) {
    fmt.Println("Graph requested")
    return p.Graph, nil
}

func (p *GraphHandler) ShortestPath(source graphdb.Int, target graphdb.Int) (float64, error) {
    graph := p.Graph
    for i := 0; i < p.NumServers; i++ {
        if i != p.ServerId {
            addr := p.IpList[i]
            runClient(p.TransportFactory, p.ProtocolFactory, addr, false, func(client *graphdb.GraphCRUDClient) (err error){
                g, err := client.GetGraph()
                graph.Vertices = append(graph.Vertices, g.Vertices...)
                graph.Edges = append(graph.Edges, g.Edges...)
                return err
            })
        }
    }

    graph.Vertices = removeDuplicatesVertices(graph.Vertices)
    graph.Edges = removeDuplicatesEdges(graph.Edges)

    return Djikstra(graph, source, target)
}
