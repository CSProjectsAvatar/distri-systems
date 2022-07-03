import logging
import middleware_pb2 as mid
import middleware_pb2_grpc as mid_grpc
import grpc
from concurrent import futures

class Middleware(mid_grpc.MiddlewareServicer):
    def UploadTournament(self, request, context):
        print('Received : ', request)
        return mid.TournamentResp(msg='OK')


def serve():
    print('Starting server...')
    # logging.info('Starting server...')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    mid_grpc.add_MiddlewareServicer_to_server(Middleware(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print('Server started.')
    # logging.info('Server started.')
    server.wait_for_termination()


if __name__ == '__main__':
    # logging.basicConfig()
    serve()
