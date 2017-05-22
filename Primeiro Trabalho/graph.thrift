namespace go graphdb
namespace py graphdb

typedef i32 int

exception ElementNotFoundException{ }

struct GraphVertex {
            1:int name,
            2:int color,
            3:double weight,
            4:string description
}

struct GraphEdge {
            1:GraphVertex firstVertex,
            2:GraphVertex secondVertex,
            3:double weight,
            4:bool isBidirectional,
            5:string description
}

struct Graph {
            1:list<GraphVertex> setOfVertex,
            2:list<GraphEdge> setOfEdges,
}

service GraphCRUD {

        void createVertex(1:GraphVertex vertex),
        GraphVertex readVertex(1:int name)  throws (1:ElementNotFoundException e),
        void deleteVertex(1:int name),
        void updateVertex(1:GraphVertex vertex),
        void createEdge(1:int firstVertex, 2:int secondVertex, 3:double weight, 4:bool isBidirectional, 5:string description),
        GraphEdge readEdge(1:int firstVertex, 2:int secondVertex)  throws (1:ElementNotFoundException e),
        void deleteEdge(1:int firstVertex, 2:int secondVertex),
        void updateEdge(1:int firstVertex, 2:int secondVertex, 3:double weight, 4:bool isBidirectional, 5:string description),
        list<GraphEdge> findEdgesOfVertex(1:int vertex)  throws (1:ElementNotFoundException e),
        list<GraphVertex> findNeighbours(1:int vertex)  throws (1:ElementNotFoundException e)
}