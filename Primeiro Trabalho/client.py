import sys

sys.path.append('gen-py')

from graphdb import GraphCRUD
from graphdb.ttypes import GraphVertex, GraphEdge, Graph

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

class Client:
    def __init__(self):
        pass

    def run(self, hostname = 'localhost', portNumber = 9090):
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

        client.createVertex(GraphVertex(7,1,10,"novo vertice"))
        try:
            ops = client.readVertex(20)
        except:
            print("Vertice nao encontrado")
        client.createEdge(8,6,10,True,"aresta inserida")
        client.deleteEdge(8,6)
        client.updateEdge(1,4,39.0,True,"aresta alterada")
        print(client.findEdgesOfVertex(1))
        print(client.findNeighbours(1))

if __name__ == "__main__":
    client = Client()
    client.run()