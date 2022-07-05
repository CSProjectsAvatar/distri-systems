from __future__ import print_function
from distutils.cygwinccompiler import Mingw32CCompiler
import imp

import logging
from os import stat
from random import randint
from traceback import print_tb
from importlib_metadata import files

from numpy import byte

import grpc
import middleware_pb2 as mid
import middleware_pb2_grpc as mid_grpc


class grpcNode:
    def __init__(self) -> None:
        self.tourStats = {} 
        self.remote_ip = 'localhost:50051'
        pass

    def upload_tournment(self, tour_name, tourn_type, file_list):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print("Uploading tournament: " + tour_name)

            tourReq = mid.TournamentReq(name=tour_name, tour_type=tourn_type, files=file_list)
            response = stub.UploadTournament(tourReq)
            # print(response)
            tourId = response.tourId
            # tourId = randint(0, 100)
            self.tourStats[tourId] = []
        # print("Client received: " + response.tourId)

    def get_stats(self, tour_id):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print('Calling GetStats')

            req = mid.StatsReq(tourId=tour_id)
            resp = stub.GetStats(req)
            self.tourStats[tour_id] = resp
            return resp

    def get_all_ids(self):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print('Calling GetAllIds')

            req = mid.AllIdsReq()
            resp: mid.AllIdsResp = stub.GetAllIds(req)
            resp.tourIds

            for id in resp.tourIds:
                self.tourStats[id] = []

    def get_all_stats(self):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print('Calling GetAllStats')

            self.get_all_ids()
            for id in self.tourStats:
                req = mid.StatsReq(tourId=id)
                # resp = stub.GetStats(req)
                # self.tourStats[id] = resp
                # mock
                mockStat: mid.StatsResp = mid.StatsResp(
                    tourName='test',
                    winner='Player' + str(randint(0, 100)),
                    bestPlayer='Player' + str(randint(0, 100)),
                )
                self.tourStats[id].append(mockStat)
                print(mockStat)
                self.tourStats[id] = {'name': 'tour' + str(id), 'stats': {'score': str(randint(0,100)), 'time': str(randint(0,100))}}
                print(self.tourStats[id])
                # print(resp)




if __name__ == '__main__':
    logging.basicConfig()
    node = grpcNode()
    # list = mid.FileList(files=[mid.File(name='test.py', data=b'print("Hello World!")')])
    # node.upload_tournment(list)
