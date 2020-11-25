# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import empty_pb2 as empty__pb2
import user_pb2 as user__pb2


class UserServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetUserByEmail = channel.unary_unary(
                '/api.UserService/GetUserByEmail',
                request_serializer=user__pb2.GetUserByEmailRequest.SerializeToString,
                response_deserializer=user__pb2.User.FromString,
                )
        self.CreateUser = channel.unary_unary(
                '/api.UserService/CreateUser',
                request_serializer=user__pb2.NewUser.SerializeToString,
                response_deserializer=user__pb2.User.FromString,
                )
        self.UpdateUserEmail = channel.unary_unary(
                '/api.UserService/UpdateUserEmail',
                request_serializer=user__pb2.UpdateUserEmailRequest.SerializeToString,
                response_deserializer=user__pb2.User.FromString,
                )
        self.RemoveUser = channel.unary_unary(
                '/api.UserService/RemoveUser',
                request_serializer=user__pb2.RemoveUserRequest.SerializeToString,
                response_deserializer=empty__pb2.Empty.FromString,
                )


class UserServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetUserByEmail(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateUserEmail(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveUser(self, request, context):
        """TODO i think there is a packaged golang empty data type
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetUserByEmail': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUserByEmail,
                    request_deserializer=user__pb2.GetUserByEmailRequest.FromString,
                    response_serializer=user__pb2.User.SerializeToString,
            ),
            'CreateUser': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateUser,
                    request_deserializer=user__pb2.NewUser.FromString,
                    response_serializer=user__pb2.User.SerializeToString,
            ),
            'UpdateUserEmail': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateUserEmail,
                    request_deserializer=user__pb2.UpdateUserEmailRequest.FromString,
                    response_serializer=user__pb2.User.SerializeToString,
            ),
            'RemoveUser': grpc.unary_unary_rpc_method_handler(
                    servicer.RemoveUser,
                    request_deserializer=user__pb2.RemoveUserRequest.FromString,
                    response_serializer=empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.UserService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class UserService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetUserByEmail(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.UserService/GetUserByEmail',
            user__pb2.GetUserByEmailRequest.SerializeToString,
            user__pb2.User.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.UserService/CreateUser',
            user__pb2.NewUser.SerializeToString,
            user__pb2.User.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateUserEmail(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.UserService/UpdateUserEmail',
            user__pb2.UpdateUserEmailRequest.SerializeToString,
            user__pb2.User.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.UserService/RemoveUser',
            user__pb2.RemoveUserRequest.SerializeToString,
            empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
