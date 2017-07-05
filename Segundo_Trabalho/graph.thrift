namespace go graphdb
namespace py graphdb

typedef i32 int

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
            1:list<GraphVertex> vertices,
            2:list<GraphEdge> edges,
}

service GraphCRUD {

        void createVertex(1:GraphVertex vertex),
        GraphVertex readVertex(1:int name),
        void deleteVertex(1:int name),
        void updateVertex(1:GraphVertex vertex),
        void createEdge(1:GraphEdge edge),
        GraphEdge readEdge(1:int firstVertex, 2:int secondVertex),
        void deleteEdge(1:GraphEdge edge),
        void updateEdge(1:GraphEdge edge),
        list<GraphEdge> findEdgesOfVertex(1:int name),
        list<GraphVertex> findNeighbours(1:int name),
        void deleteEdgesOfVertex(1:int name),
        void deleteEdgeColateral(1:GraphEdge edge),
        void updateEdgeColateral(1:GraphEdge edge),
        list<GraphVertex> findNeighboursRemote(1:int name)
        Graph getGraph(),
        double shortestPath(1:int source, 2:int target)
}