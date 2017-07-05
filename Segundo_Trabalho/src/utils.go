package main

import (
     "fmt"
     "sync"
     "errors"
    // "strconv"
    "SD/Segundo_Trabalho/gen-go/graphdb"
)

var rw sync.RWMutex
const maxDistance = 1000000.0

func IsVertexInGraph(graph *graphdb.Graph, name graphdb.Int) bool {
    for i := 0; i < len(graph.Vertices); i++ {
        if name == graph.Vertices[i].Name {
            return true
        }
    }
    return false
}

func isEdgeInGraph(graph *graphdb.Graph, e *graphdb.GraphEdge) bool {
    for i:=0; i < len(graph.Edges); i++ {
        if e.Equals(graph.Edges[i]) {
            return true
        }
    }
    return false
}

func PrintVertex(vertex *graphdb.GraphVertex) {
    fmt.Print("Name: ", vertex.Name, "\n" )
    fmt.Print("Description: ", vertex.Description, "\n")
    fmt.Println("=============")
}


func PrintVertices(vertices []*graphdb.GraphVertex) {
    for i := 0; i < len(vertices); i++ {
        PrintVertex(vertices[i])
    }
}

func PrintEdge(edge *graphdb.GraphEdge) {
    fmt.Print("FirstVertex: ", edge.FirstVertex, "\n")
    fmt.Print("SecondVertex: ", edge.SecondVertex, "\n")
    fmt.Print("Description: ", edge.Description, "\n")
    fmt.Print("IsBidirectional: ", edge.IsBidirectional, "\n")
    fmt.Println("=============")
}

func PrintEdges(edges []*graphdb.GraphEdge) {
    for i := 0; i < len(edges); i++ {
        PrintEdge(edges[i])
    }
}

func InsertVertex(vertices []*graphdb.GraphVertex, v *graphdb.GraphVertex) []*graphdb.GraphVertex {
    rw.Lock()
    defer rw.Unlock()
    vertices = append(vertices, v)
    return vertices
}

func FetchVertex(vertices []*graphdb.GraphVertex, name graphdb.Int) (*graphdb.GraphVertex, error){
   rw.RLock()
   defer rw.RUnlock()
   for i := 0; i < len(vertices); i++ {
        if name == vertices[i].Name {
            return  vertices[i], nil
        }
    }
    return nil, errors.New("Vértice não encontrado")
}

func DeleteVertex(graph *graphdb.Graph, name graphdb.Int) *graphdb.Graph {
    rw.Lock()
    // newVertices := make([]*graphdb.GraphVertex, len(graph.Vertices))
    // copy(newVertices, graph.Vertices)
    defer rw.Unlock()
    for i := 0; i < len(graph.Vertices); i++ {
        if name == graph.Vertices[i].Name {
            graph.Vertices = append(graph.Vertices[:i], graph.Vertices[i+1:]...)
            graph = DeleteEdgesOfVertex(graph , name)
            i = i - 1
        }
    }
    // graph.Vertices = newVertices
    return graph
}

func UpdateVertex(vertices []*graphdb.GraphVertex, vertex *graphdb.GraphVertex) []*graphdb.GraphVertex{
    rw.Lock()
    defer rw.Unlock()
    for i := 0; i < len(vertices); i++ {
        if vertices[i].Name == vertex.Name {
            vertices[i] = vertex
            break
        }
    }

    return vertices
}

func InsertEdge(edges []*graphdb.GraphEdge, e *graphdb.GraphEdge) []*graphdb.GraphEdge{
    rw.Lock()
    defer rw.Unlock()
    edges = append(edges, e)
    return edges
}

func FetchEdge(edges []*graphdb.GraphEdge , nameFirstVertex graphdb.Int, nameSecondVertex graphdb.Int) (*graphdb.GraphEdge, error){
    rw.RLock()
    defer rw.RUnlock()
    fmt.Println("No pega aresta:     ", nameFirstVertex, nameSecondVertex)
    for i:=0; i < len(edges); i++ {
        if edges[i].FirstVertex.Name == nameFirstVertex && edges[i].SecondVertex.Name  == nameSecondVertex {
            return edges[i], nil
        } else if  edges[i].IsBidirectional && edges[i].FirstVertex.Name == nameSecondVertex && edges[i].SecondVertex.Name  == nameFirstVertex {
            return edges[i], nil
        }
    }

    return nil, errors.New("Aresta não encontrada")
}

func DeleteEdgesOfVertex(graph *graphdb.Graph, name graphdb.Int)  *graphdb.Graph{

    // newEdges := make([]*graphdb.GraphEdge, len(graph.Edges))
    // copy(newEdges, graph.Edges)
    for i := 0; i < len(graph.Edges); i++ {
        if graph.Edges[i].FirstVertex.Name == name || graph.Edges[i].SecondVertex.Name == name {
            graph.Edges = append(graph.Edges[:i], graph.Edges[i+1:]...)
            i = i - 1
        }
    }

    // graph.Edges = newEdges
    return graph
}

func DeleteEdge(edges []*graphdb.GraphEdge, e *graphdb.GraphEdge) ([]*graphdb.GraphEdge){
    rw.Lock()
    defer rw.Unlock()
    for i := 0; i < len(edges); i++ {
        if e.Equals(edges[i]) {
            edges = append(edges[:i], edges[i+1:]...)
            i = i - 1
        }
    }
    return edges
}

func UpdateEdge(edges []*graphdb.GraphEdge, e *graphdb.GraphEdge) ([]*graphdb.GraphEdge) {
    rw.Lock()
    defer rw.Unlock()
    for i := 0; i < len(edges); i++ {
        if e.Equals(edges[i]) {
            edges[i] = e
            break
        }
    }

    return edges
}

func FindEdgesOfVertex(edges []*graphdb.GraphEdge, name graphdb.Int) (edgesOfVertex []*graphdb.GraphEdge) {
    rw.RLock()
    defer rw.RUnlock()
    for i := 0; i < len(edges); i++ {
        if edges[i].FirstVertex.Name == name || edges[i].SecondVertex.Name == name {
            edgesOfVertex = append(edgesOfVertex, edges[i])
        }
    }

    return
}

func FindNeighbours(edges []*graphdb.GraphEdge, name graphdb.Int) (neighbours []*graphdb.GraphVertex) {
    rw.RLock()
    defer rw.RUnlock()
    for i := 0; i < len(edges); i++ {
        if name == edges[i].FirstVertex.Name {
            neighbours = append(neighbours, edges[i].SecondVertex)
        } else if  name == edges[i].SecondVertex.Name{
            neighbours = append(neighbours, edges[i].FirstVertex)
        }
    }

    return
}


func removeDuplicatesVertices(elements []*graphdb.GraphVertex) []*graphdb.GraphVertex {
    // Use map to record duplicates as we find them.
    encountered := map[graphdb.Int]bool{}
    result := make([]*graphdb.GraphVertex, 0)

    for _, v := range elements {
        if encountered[v.Name] == true {
            // Do not add duplicate.
        } else {
            // Record this element as an encountered element.
            encountered[v.Name] = true
            // Append to result slice.
            result = append(result, v)
        }
    }
    // Return the new slice.
    return result
}

func removeDuplicatesEdges(elements []*graphdb.GraphEdge) []*graphdb.GraphEdge {
    // Use map to record duplicates as we find them.
    encountered := map[string]bool{}
    result := make([]*graphdb.GraphEdge, 0)

    for _, e := range elements {
        if encountered[e.Description] == true {
            // Do not add duplicate.
        } else {
            // Record this element as an encountered element.
            encountered[e.Description] = true
            // Append to result slice.
            result = append(result, e)
        }
    }
    // Return the new slice.
    return result
}

func Djikstra(graph *graphdb.Graph, source graphdb.Int, target graphdb.Int) (float64, error) {
    visitedVertices := make(map[int]bool)
    distances := make(map[int]float64)

    for _, v := range graph.Vertices {
        visitedVertices[int(v.Name)] = false
        distances[int(v.Name)] = maxDistance
    }

    distances[int(source)] = 0

    u := minDist(distances, visitedVertices)
    for u > 0 && u != int(target) {
        visitedVertices[u] = true
        /*fmt.Println("u: ", u)
        fmt.Println("dist: ", distances)
        fmt.Println("visited: ", visitedVertices)*/
        for _, v := range FindNeighbours(graph.Edges, graphdb.Int(u)) {
            fmt.Println("vizinho:  ", v)
            fmt.Println("u:  " ,u)
            if e, err := FetchEdge(graph.Edges, graphdb.Int(u), v.Name); err == nil {
                fmt.Println("aresta", e)
                PrintEdges(graph.Edges)
                if alt := distances[u] + e.Weight; alt < distances[int(v.Name)] {
                    distances[int(v.Name)] = alt
                }
            }
        }

        u = minDist(distances, visitedVertices)
        fmt.Println("u: ", u)
        fmt.Println("dist: ", distances)
        fmt.Println("visited: ", visitedVertices)
    }

    if distances[int(target)] == maxDistance {
        return 0, errors.New("Não existe caminho entre os vértices")
    } else {
        return distances[int(target)], nil
    }
}

func minDist(distances map[int]float64, visitedVertices map[int]bool) (d int) {
    min := maxDistance
    d = -100
    for k, v := range distances {
        if v < min && !visitedVertices[k] {
            fmt.Println(k,v)
            fmt.Println("visited", visitedVertices)
            fmt.Println("dist:", distances)
            min = v
            d = k
        }
    }

    fmt.Println("d: ", d)
    return
}

func Mock(serverId int) graphdb.Graph {

    switch (serverId) {
        case 0:
            return Mock0()
        case 1:
            return Mock1()
        case 2:
            return Mock2()
        default:
            return Mock0()
    }
}

func Mock0() graphdb.Graph {
    var vertices []*graphdb.GraphVertex
    var edges []*graphdb.GraphEdge

    var v1  = graphdb.GraphVertex {
        Name:3,
        Color:3,
        Weight:15.0,
        Description:"terceiro vertice"}

    var v2 = graphdb.GraphVertex {
        Name:6,
        Color:6,
        Weight:20.0,
        Description:"sexto vertice"}

    var v3 = graphdb.GraphVertex {
        Name:1,
        Color:1,
        Weight:10.0,
        Description:"primeiro vertice"}

    var v4 = graphdb.GraphVertex {
        Name:5,
        Color:5,
        Weight:10.0,
        Description:"quinto vertice"}

    var e1 = graphdb.GraphEdge {
        FirstVertex:&v2,
        SecondVertex:&v1,
        Weight:10.0,
        IsBidirectional:false,
        Description:"aresta 6-3"}

    var e2 = graphdb.GraphEdge {
        FirstVertex:&v1,
        SecondVertex:&v3,
        Weight:10.0,
        IsBidirectional:true,
        Description:"aresta 3-1"}

    var e3 = graphdb.GraphEdge {
        FirstVertex:&v1,
        SecondVertex:&v4,
        Weight:15.0,
        IsBidirectional:false,
        Description:"aresta 3-5"}

    vertices = append(vertices, &v1)
    vertices = append(vertices, &v2)
    edges = append(edges, &e1)
    edges = append(edges, &e2)
    edges = append(edges, &e3)

    return graphdb.Graph{Vertices: vertices, Edges: edges}
}

func Mock1() graphdb.Graph {
    var vertices []*graphdb.GraphVertex
    var edges []*graphdb.GraphEdge

    var v1  = graphdb.GraphVertex {
        Name:4,
        Color:4,
        Weight:15.0,
        Description:"quarto vertice"}

    var v2 = graphdb.GraphVertex {
        Name:1,
        Color:1,
        Weight:10.0,
        Description:"primeiro vertice"}

    var v3 = graphdb.GraphVertex {
        Name:3,
        Color:3,
        Weight:10.0,
        Description:"terceiro vertice"}

    var e1 = graphdb.GraphEdge {
        FirstVertex:&v1,
        SecondVertex:&v2,
        Weight:20.0,
        IsBidirectional:false,
        Description:"aresta 4-1"}

    var e2 = graphdb.GraphEdge {
        FirstVertex:&v2,
        SecondVertex:&v3,
        Weight:10.0,
        IsBidirectional:true,
        Description:"aresta 3-1"}

    vertices = append(vertices, &v1)
    vertices = append(vertices, &v2)
    edges = append(edges, &e1)
    edges = append(edges, &e2)

    return graphdb.Graph{Vertices: vertices, Edges: edges}
}

func Mock2() graphdb.Graph {
    var vertices []*graphdb.GraphVertex
    var edges []*graphdb.GraphEdge

    var v1  = graphdb.GraphVertex {
        Name:5,
        Color:5,
        Weight:10.0,
        Description:"quinto vertice"}

    var v2 = graphdb.GraphVertex {
        Name:2,
        Color:2,
        Weight:20.0,
        Description:"segundo vertice"}

    var v3 = graphdb.GraphVertex {
        Name:6,
        Color:6,
        Weight:20.0,
        Description:"sexto vertice"}

    var e1 = graphdb.GraphEdge {
        FirstVertex:&v1,
        SecondVertex:&v2,
        Weight:13.0,
        IsBidirectional:true,
        Description:"aresta 5-2"}

    var e2 = graphdb.GraphEdge {
        FirstVertex:&v2,
        SecondVertex:&v3,
        Weight:10.0,
        IsBidirectional:false,
        Description:"aresta 2-6"}

    vertices = append(vertices, &v1)
    vertices = append(vertices, &v2)
    edges = append(edges, &e1)
    edges = append(edges, &e2)

    return graphdb.Graph{Vertices: vertices, Edges: edges}
}
