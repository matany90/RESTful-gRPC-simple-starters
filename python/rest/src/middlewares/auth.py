import json
from werkzeug.wrappers import Response

from src.utils.configs import EXAMPLE_REST_SERVICE_BASIC_TOKEN

def check_if_authenticated(request, environ, start_response):
    '''
    Checking if the current user if authenticated
    for making current request.
    '''
    # split auth
    splitted_auth = request.headers.get("authorization", "").split(" ")

    # check authorization
    if len(splitted_auth) == 2:

        # check basic authrozation
        if (splitted_auth[0] == "Basic"):
            return check_basic_token(splitted_auth[1])

        # check bearer authorization
        elif (splitted_auth[0] == "Bearer"):
            return check_bearer_token(splitted_auth[1])

        # generate error
        else:
            return generate_auth_error("unauthorized request. expect authorization header includes Basic or Bearer token.", 401, environ, start_response)
    else:
        return generate_auth_error("unauthorized request. authorization header does not have valid structure.", 401, environ, start_response)

def check_basic_token(token):
    '''
    Validate Basic token.
    '''
    # TODO: Implement Basic token authorization.
    # If the token is valid, return None
    return token != EXAMPLE_REST_SERVICE_BASIC_TOKEN

def check_bearer_token(token):
    '''
    Validate Bearer token.
    '''
    # TODO: Implement Bearer token authorization.
    # If the token is valid, return None
    return None

def generate_auth_error(err_message, status_code, environ, start_response):
    '''
    Handle auth error.
    '''
    # define err json as string
    json_response = json.dumps({
        "message": err_message,
        "code": status_code
    })

    # init response instance
    response = Response(json_response, content_type="application/json; charset=utf-8", status=status_code)

    # return callable response
    return response(environ, start_response)
