from proto.matan_service_pb2_grpc import MatanServiceServicer
from proto.matan_service_pb2 import HelloReply

class MatanServiceServicer(MatanServiceServicer):
    '''
    MatanServiceServicer inherits from the compiled
    proto file, and implements the handlers
    defined in it.
    '''
    def SayHello(self, request, context):
        '''
        SayHello Handler implementation.
        '''
        return HelloReply(message=f"Hello {request.name}")

    def SayHi(self, request, context):
        '''
        SayHi Handler implementation.
        '''
        return HelloReply(message=f"Hi {request.name}")