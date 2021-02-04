
import grpc
import os
from concurrent.futures import ThreadPoolExecutor

from proto.matan_service_pb2_grpc import add_MatanServiceServicer_to_server
from src.handlers.matan_service import MatanServiceServicer

def create_app():
    '''
    Init gRPC server instance.
    '''
    # init gRPC server
    server = grpc.server(ThreadPoolExecutor(max_workers=10))

    # define host and port
    host = os.environ.get("HOST") or "0.0.0.0"
    port = os.environ.get("PORT") or "50006"

    # attach handlers to server instance
    add_MatanServiceServicer_to_server(
        MatanServiceServicer(),
        server
    )

    # register port and host to server instance
    server.add_insecure_port(f"{host}:{port}")

    # return server instance, host and port
    return server, host, port

if __name__ == "__main__":
    # init app
    app, host, port = create_app()

    # run gRPC app
    app.start()
    print(f"gRPC python server running on port {port}...")

    # exit on termination
    app.wait_for_termination()