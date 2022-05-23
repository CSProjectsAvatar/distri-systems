from __future__ import print_function
from distutils.cygwinccompiler import Mingw32CCompiler
import imp

import logging
from traceback import print_tb
from importlib_metadata import files

from numpy import byte

import grpc
import middleware_pb2 as mid
import middleware_pb2_grpc as mid_grpc


class grpcNode:
    def __init__(self) -> None:
        self.tournament_dict = {} # {tournament_name: (tournType, [file_list])}
        pass

    def upload_tournment(self, tour_name, tourn_type, file_list):
        with grpc.insecure_channel('localhost:50051') as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print("Uploading tournament: " + tour_name)

            tourReq = mid.TournamentReq(name=tour_name, tour_type=tourn_type, files=file_list)
            response = stub.UploadTournament(tourReq)
        print("Client received: " + response.msg)

    
if __name__ == '__main__':
    logging.basicConfig()
    node = grpcNode()
    # list = mid.FileList(files=[mid.File(name='test.py', data=b'print("Hello World!")')])
    # node.upload_tournment(list)
