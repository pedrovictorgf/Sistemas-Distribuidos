import sys
import Queue
import threading

sys.path.append('gen-py')

from graphdb import GraphCRUD
from graphdb.ttypes import GraphVertex, GraphEdge, Graph

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

q = Queue.Queue()

def firstClient(hostname = 'localhost', portNumber = 9090):
        # Make socket
        transport = TSocket.TSocket(hostname, portNumber)
        # Buffering is critical. Raw sockets are very slow
        transport = TTransport.TBufferedTransport(transport)
        # Wrap in a protocol
        protocol = TBinaryProtocol.TBinaryProtocol(transport)
        # Create a client to use the protocol encoder
        client = GraphCRUD.Client(protocol)
        # Connect!
        transport.open()

        #client.deleteVertex(3)
        print(client.findNeighbours(6))
        '''client.createVertex(GraphVertex(9,1,10,"novo vertice"))
        try:
            v = client.readVertex(20)
        except:
            print("Vertice nao encontrado")
        client.createEdge(8,6,10,True,"aresta inserida")
        client.deleteEdge(2,4)
        client.updateEdge(1,4,20.0,True,"aresta alterada")
        try:
            print("Arestas:" , client.findEdgesOfVertex(6))
        except:
            print("Arestas nao encontradas")
        try:
            print("Vizinhos:", client.findNeighbours(3))
        except:
            print("Vizinhos nao encontrados")
'''
        q.put("Thread finalizada")

def secondClient( hostname = 'localhost', portNumber = 9090):
        # Make socket
        transport = TSocket.TSocket(hostname, portNumber)
        # Buffering is critical. Raw sockets are very slow
        transport = TTransport.TBufferedTransport(transport)
        # Wrap in a protocol
        protocol = TBinaryProtocol.TBinaryProtocol(transport)
        # Create a client to use the protocol encoder
        client = GraphCRUD.Client(protocol)
        # Connect!
        transport.open()
        '''client.createVertex(GraphVertex(7,1,10,"novo vertice"))
        client.createVertex(GraphVertex(7,1,10,"novo vertice"))
        try:
            v = client.readVertex(4)
            print(v)
        except:
            print("Vertice nao encontrado")
        client.updateVertex(GraphVertex(1,50,50,"vertice alterado"))
        client.deleteVertex(3)
        try:
            print("Arestas:" , client.findEdgesOfVertex(1))
        except:
            print("Arestas nao encontradas")

        try:
            print("Vizinhos:", client.findNeighbours(1))
        except:
            print("Vizinhos nao encontrados")
'''
        q.put("Thread finalizada")


class Client:
    def __init__(self):
        pass

    def run(self):
        t1 = threading.Thread(target=firstClient)
        t2 = threading.Thread(target=secondClient)
        t1.daemon = True
        t2.daemon = True
        t1.start()
        #t2.start()

        q.get()

if __name__ == "__main__":
    client = Client()
    client.run()