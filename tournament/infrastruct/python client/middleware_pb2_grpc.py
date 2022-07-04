# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import middleware_pb2 as middleware__pb2


class MiddlewareStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UploadTournament = channel.unary_unary(
                '/pb.Middleware/UploadTournament',
                request_serializer=middleware__pb2.TournamentReq.SerializeToString,
                response_deserializer=middleware__pb2.TournamentResp.FromString,
                )
        self.RunTournament = channel.unary_stream(
                '/pb.Middleware/RunTournament',
                request_serializer=middleware__pb2.RunReq.SerializeToString,
                response_deserializer=middleware__pb2.RunResp.FromString,
                )
        self.GetStats = channel.unary_unary(
                '/pb.Middleware/GetStats',
                request_serializer=middleware__pb2.StatsReq.SerializeToString,
                response_deserializer=middleware__pb2.StatsResp.FromString,
                )


class MiddlewareServicer(object):
    """Missing associated documentation comment in .proto file."""

    def UploadTournament(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RunTournament(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetStats(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MiddlewareServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UploadTournament': grpc.unary_unary_rpc_method_handler(
                    servicer.UploadTournament,
                    request_deserializer=middleware__pb2.TournamentReq.FromString,
                    response_serializer=middleware__pb2.TournamentResp.SerializeToString,
            ),
            'RunTournament': grpc.unary_stream_rpc_method_handler(
                    servicer.RunTournament,
                    request_deserializer=middleware__pb2.RunReq.FromString,
                    response_serializer=middleware__pb2.RunResp.SerializeToString,
            ),
            'GetStats': grpc.unary_unary_rpc_method_handler(
                    servicer.GetStats,
                    request_deserializer=middleware__pb2.StatsReq.FromString,
                    response_serializer=middleware__pb2.StatsResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'pb.Middleware', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Middleware(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def UploadTournament(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/pb.Middleware/UploadTournament',
            middleware__pb2.TournamentReq.SerializeToString,
            middleware__pb2.TournamentResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RunTournament(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/pb.Middleware/RunTournament',
            middleware__pb2.RunReq.SerializeToString,
            middleware__pb2.RunResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetStats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/pb.Middleware/GetStats',
            middleware__pb2.StatsReq.SerializeToString,
            middleware__pb2.StatsResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
