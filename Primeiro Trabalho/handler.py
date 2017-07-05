import sys
sys.path.append('gen-py')

from graphdb.ttypes import GraphVertex, GraphEdge, Graph,  ElementNotFoundException
from threading import Lock

mutex = Lock()

def readFile():
    setOfVertex = []
    setOfEdges = []

    mutex.acquire()
    with open("vertex_base", "r") as file:
        for line in file:
            line = line.replace("\n","")
            args = line.split(",")
            vertex = GraphVertex(int(args[0]), int(args[1]), float(args[2]), args[3])
            setOfVertex.append(vertex)

    with(open("edge_base","r")) as file:
        for line in file:
            line = line.replace("\n","")
            args = line.split(",")
            firstVertex = searchVertex(int(args[0]),setOfVertex)
            secondVertex = searchVertex(int(args[1]),setOfVertex)
            if(firstVertex != None and secondVertex != None):
                edge = GraphEdge(firstVertex, secondVertex, float(args[2]), bool(args[3]), args[4])
                setOfEdges.append(edge)

    mutex.release()
    return Graph(setOfVertex, setOfEdges)

def searchVertex(index, setOfVertex):
    for v in setOfVertex:
        if v.name == index:
            return v
    return None

def searchEdge(firstV, secondV, setOfEdges):
    for e in setOfEdges:
            if((e.firstVertex.name == firstV and e.secondVertex.name == secondV) or (e.secondVertex.name == firstV and e.firstVertex.name == secondV)):
                return e
    return None

def insertNewVertex(vertex):
    with open("vertex_base","a") as f:
        f.write(vertexToString(vertex))

def insertNewEdge(edge):
    with open("edge_base", "a") as f:
        f.write(edgeToString(edge))

def deleteVertexFromGraph(name, graph):
    setOfVertex = graph.setOfVertex
    setOfEdges = graph.setOfEdges
    auxForVertex = []
    auxForEdges = []
    for i in range(len(setOfVertex)):
        if(setOfVertex[i].name == name):
            auxForVertex = setOfVertex[:i]+setOfVertex[i+1:]

    for edge in setOfEdges:
        if(edge.firstVertex.name != name and edge.secondVertex.name != name):
            auxForEdges.append(edge)
    writeGraphToFile(Graph(auxForVertex, auxForEdges))

def deleteEdgeFromGraph(firstVertex, secondVertex, setOfEdges):
    auxForEdges = []
    printEdges(setOfEdges)
    for e in setOfEdges:
        if((e.firstVertex.name != firstVertex or e.secondVertex.name != secondVertex)):
            if(e.isBidirectional):
                if(e.secondVertex.name != firstVertex or e.firstVertex.name != secondVertex):
                    auxForEdges.append(e)
            else:
                auxForEdges.append(e)
    writeEdges(auxForEdges)


def updateVertex(vertex, graph):
    setOfVertex = graph.setOfVertex
    auxForVertex = []
    hasFound = False
    for i in range(len(setOfVertex)):
        if(vertex.name == setOfVertex[i].name):
            auxForVertex = setOfVertex[:i]+setOfVertex[i+1:]
            auxForVertex.append(vertex)
            hasFound = True
            break

    if(hasFound):
        writeVertex(auxForVertex)

def updateEdge(edge, setOfEdges):
    auxForEdges = []
    hasFound = False
    for i in range(len(setOfEdges)):
        if((setOfEdges[i].firstVertex.name == edge.firstVertex.name and setOfEdges[i].secondVertex.name == edge.secondVertex.name) or (setOfEdges[i].secondVertex.name == edge.firstVertex.name and setOfEdges[i].firstVertex.name == edge.secondVertex.name)):
            auxForEdges = setOfEdges[:i] + setOfEdges[i+1:]
            auxForEdges.append(edge)
            hasFound = True
            break
    if(hasFound):
        writeEdges(auxForEdges)

def vertexToString(vertex):
    vertex_string = ""
    vertex_string += str(vertex.name)+","
    vertex_string += str(vertex.color)+","
    vertex_string += str(vertex.weight)+","
    vertex_string += vertex.description+"\n"
    return vertex_string

def edgeToString(edge):
    edge_string = ""
    edge_string += str(edge.firstVertex.name)+","
    edge_string += str(edge.secondVertex.name)+","
    edge_string += str(edge.weight)+","
    edge_string += str(edge.isBidirectional)+","
    edge_string += edge.description+"\n"
    return edge_string

def printVertexs(setOfVertex):
    for vertex in setOfVertex:
        print("Name:", vertex.name)
        print("Color:", vertex.color)
        print("Weight:", vertex.weight)
        print("Description:", vertex.description)
        print("--------------")

def printEdges(setOfEdges):
    for edge in setOfEdges:
        print("FirstVertex:", edge.firstVertex)
        print("SecondVertex:", edge.secondVertex)
        print("Weight:", edge.weight)
        print("IsBidirectional", edge.isBidirectional)
        print("Description", edge.description)
        print("--------------")

def writeGraphToFile(graph):
    writeVertex(graph.setOfVertex)
    writeEdges(graph.setOfEdges)

def writeVertex(setOfVertex):
    vertexData = ""
    for vertex in setOfVertex:
        vertexData += vertexToString(vertex)
    with open("vertex_base", "w") as f:
        f.write(vertexData)

def writeEdges(setOfEdges):
    edgesData = ""
    for edge in setOfEdges:
        edgesData += edgeToString(edge)
    with open("edge_base", "w") as f:
        f.write(edgesData)

def findEdgesOfVertex(vertex, setOfEdges):
    auxForEdges = []
    for e in setOfEdges:
        if(vertex.name == e.firstVertex.name or vertex.name == e.secondVertex.name):
            auxForEdges.append(e)

    return auxForEdges

def findNeighbours(vertex, setOfEdges):
    auxForNeighbours = []
    for e in setOfEdges:
        if(vertex.name == e.firstVertex.name):
            auxForNeighbours.append(e.secondVertex)
        elif(vertex.name == e.secondVertex.name):
            auxForNeighbours.append(e.firstVertex)
    return auxForNeighbours

class Handler:

    def __init__(self):
        pass

    def createVertex(self, vertex):
        """
        Parameters:
        - vertex
        """
        graph = readFile()
        mutex.acquire()
        if(searchVertex(vertex.name, graph.setOfVertex) == None):
            print("Inserindo vertice")
            insertNewVertex(vertex)
        else:
            print("Vertice ja existe")
        mutex.release()

    def readVertex(self, name):
        """
        Parameters:
        - name
        """
        graph = readFile()
        v = searchVertex(name, graph.setOfVertex)
        if(v != None):
            return v
        else:
            e = ElementNotFoundException()
            raise e


    def deleteVertex(self, name):
        """
        Parameters:
        -name
        """
        graph = readFile()
        deleteVertexFromGraph(name, graph)

    def updateVertex(self, vertex):
        """
        Parameters:
        - vertex
        """
        graph = readFile()
        updateVertex(vertex, graph)

    def createEdge(self,  firstVertex, secondVertex, weight, isBidirectional, description):
        """
        Parameters:
        - edge
        """
        graph = readFile()
        mutex.acquire()
        firstV = searchVertex(firstVertex, graph.setOfVertex)
        secondV = searchVertex(secondVertex, graph.setOfVertex)
        e = searchEdge(firstVertex, secondVertex, graph.setOfEdges)
        if(firstV != None and secondV != None and e == None ):
            insertNewEdge(GraphEdge(firstV, secondV, weight, isBidirectional, description))
        mutex.release()

    def readEdge(self, firstVertex, secondVertex):
        """
        Parameters:
        - firstVertex
        - secondVertex
        """
        graph = readFile()
        e = searchEdge(firstVertex, secondVertex, graph.setOfEdges)
        if (e != None):
            return e
        else:
            e = ElementNotFoundException()
            raise e

    def deleteEdge(self, firstVertex, secondVertex):
        """
        Parameters:
        - firstVertex
        - secondVertex
        """
        graph = readFile()
        deleteEdgeFromGraph(firstVertex, secondVertex, graph.setOfEdges)

    def updateEdge(self, firstVertex, secondVertex, weight, isBidirectional, description):
        """
        Parameters:
        - firstVertex
        - secondVertex
        """
        graph = readFile()
        firstV = searchVertex(firstVertex, graph.setOfVertex)
        secondV = searchVertex(secondVertex, graph.setOfVertex)
        if(firstV != None and secondV != None):
            edge = GraphEdge(firstV, secondV, weight, isBidirectional, description)
            updateEdge(edge, graph.setOfEdges)


    def findEdgesOfVertex(self, vertex):
        """
        Parameters:
        - vertex
        """
        graph = readFile()
        v = searchVertex(vertex, graph.setOfVertex)
        if(v != None):
            return findEdgesOfVertex(v, graph.setOfEdges)
        else:
            e = ElementNotFoundException()
            raise e

    def findNeighbours(self, vertex):
        """
        Parameters:
        - vertex
        """
        graph = readFile()
        v = searchVertex(vertex, graph.setOfVertex)
        if(v != None):
            return findNeighbours(v, graph.setOfEdges)
        else:
            e = ElementNotFoundException()
            raise e