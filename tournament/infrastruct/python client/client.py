from __future__ import print_function
from distutils.cygwinccompiler import Mingw32CCompiler
import imp

import logging
from os import stat
from random import randint
from traceback import print_tb
from importlib_metadata import files

from numpy import byte
from typing import List

import grpc
import middleware_pb2 as mid
import middleware_pb2_grpc as mid_grpc


def ip_switch(method):
    def decorator(self, *args, **kwargs):
        start = self.ips_idx
        for offset in range(len(self.ips)):
            self.ips_idx = (start + offset) % len(self.ips)
            self.remote_ip = self.ips[self.ips_idx]
            try:
                print('Trying with IP:', self.remote_ip)
                return method(self, *args, **kwargs)
            except grpc.RpcError:
                pass

        raise Exception('All IPs failed')

    return decorator


class grpcNode:
    def __init__(self) -> None:
        self.tourStats = {}
        self.remote_ip = 'localhost:8082'
        # self.remote_ip = '192.168.122.219:8082'
        self.ips = self.get_remote_ips()
        self.ips_idx = 0  # self.ips[self.ips_idx] is self.remote_ip and that holds thanks to ip_switch decorator

    def get_remote_ips(self) -> List[str]:
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            resp: mid.IPsResp = stub.GetIPs(mid.IpsReq())
            return resp.ips

    @ip_switch
    def upload_tournment(self, tour_name, tourn_type, file_list):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print("Uploading tournament: " + tour_name + " to IP: " + self.remote_ip)

            tourReq = mid.TournamentReq(name=tour_name, tour_type=tourn_type, files=file_list)
            response = stub.UploadTournament(tourReq)
            # print(response)
            tourId = response.tourId
            # tourId = randint(0, 100)
            self.tourStats[tourId] = []
        # print("Client received: " + response.tourId)

    @ip_switch
    def get_stats(self, tour_id):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print('Calling GetStats')

            req = mid.StatsReq(tourId=tour_id)
            resp = stub.GetStats(req)
            self.tourStats[tour_id] = resp
            return resp

    @ip_switch
    def get_all_ids(self):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print('Calling GetAllIds')

            req = mid.AllIdsReq()
            resp: mid.AllIdsResp = stub.GetAllIds(req)
            resp.tourIds

            for id in resp.tourIds:
                self.tourStats[id] = []

    @ip_switch
    def get_all_stats(self):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print('Calling GetAllStats')

            self.get_all_ids()
            for id in self.tourStats:
                req = mid.StatsReq(tourId=id)
                resp = stub.GetStats(req)
                self.tourStats[id] = resp
                # mock
                # mockStat: mid.StatsResp = mid.StatsResp(
                #     tourName='test',
                #     winner='Player' + str(randint(0, 100)),
                #     bestPlayer='Player' + str(randint(0, 100)),
                # )
                self.tourStats[id] = resp
                # print(mockStat)
                # self.tourStats[id] = {'name': 'tour' + str(id), 'stats': {'score': str(randint(0,100)), 'time': str(randint(0,100))}}
                # print(self.tourStats[id])
                # print(resp)

    @ip_switch
    def get_rand_stats(self):
        with grpc.insecure_channel(self.remote_ip) as channel:
            stub = mid_grpc.MiddlewareStub(channel)
            print('Calling GetRandStat')

            req = mid.StatsReq(tourId='1')
            resp = stub.GetRndStats
            print(resp)
            resp = stub.GetRndStats(req)
            self.tourStats[1] = resp
            print(resp)
            return resp


if __name__ == '__main__':
    logging.basicConfig()
    node = grpcNode()
    # list = mid.FileList(files=[mid.File(name='test.py', data=b'print("Hello World!")')])
    # node.upload_tournment(list)
