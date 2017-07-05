import sys

sys.path.append('gen-py')

from graphdb import GraphCRUD
from graphdb.ttypes import GraphVertex, GraphEdge, Graph
from handler import Handler

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer



class Server(object):
    def __init__(self):
        pass

    def run(self, portNumber = 9090):
        handler = Handler()
        processor = GraphCRUD.Processor(handler)
        transport = TSocket.TServerSocket(port=portNumber)
        tfactory = TTransport.TBufferedTransportFactory()
        pfactory = TBinaryProtocol.TBinaryProtocolFactory()

        server = TServer.TThreadPoolServer(processor, transport, tfactory, pfactory)
        server.serve()

        # You could do one of these for a multithreaded server
        # server = TServer.TThreadedServer(processor, transport, tfactory, pfactory)
        # server = TServer.TThreadPoolServer(processor, transport, tfactory, pfactory)
        # server.serve()

if __name__ == '__main__':
    print("Starting the server...")
    server = Server()
    server.run()