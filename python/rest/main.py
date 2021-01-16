import os
from flask import Flask
from flask_restful import Api

from src.api.v1.foo.foo_router import foo_router
from src.middlewares.middleware import Middleware

def create_app():
    '''
    Init flask app instance.
    '''
    # init app
    app = Flask(__name__)

    # define host and port
    host = os.environ.get("HOST") or "0.0.0.0"
    port = os.environ.get("PORT") or "50005"

    # init general hanlders middlewares
    app.wsgi_app = Middleware(app.wsgi_app)

    # register foo router
    app.register_blueprint(foo_router, url_prefix="/api/v1/foo")

    # returns app instance
    return app, host, port

if __name__ == "__main__":
    # define app, host and port
    app, host, port = create_app()

    # run app instance
    app.run(host, port, debug=True)
