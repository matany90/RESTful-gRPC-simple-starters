from werkzeug.wrappers import Request

from src.middlewares.auth import check_if_authenticated

class Middleware():
    '''
    Middleware class.
    '''
    def __init__(self, app):
        # init app instance
        self.app = app

    def __call__(self, environ, start_response):
        # init request instance
        request = Request(environ)

        # check if authenticated middlware
        auth_err = check_if_authenticated(request, environ, start_response)
        if auth_err:
            return auth_err

        # pass all middlwares
        return self.app(environ, start_response)