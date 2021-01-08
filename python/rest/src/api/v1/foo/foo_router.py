from flask import Blueprint, request, jsonify

# init foo blueprint router
foo_router = Blueprint("foo", __name__)

@foo_router.route("", methods=["GET"])
def list():
    '''
    Get list of foo's.
    '''
    # Do an API call...

    # returns json response
    return jsonify(
        foos=["foo1", "foo2", "foo3"]
    )

@foo_router.route("/<id>", methods=["GET"])
def get(id):
    '''
    Get a foo by id.
    '''
    # Do an API call...

    # returns json response
    return jsonify(
        foo=f"Get foo with id of {id}"
    )

@foo_router.route("", methods=["PUT"])
def update():
    '''
    Update a foo.
    '''
    # pull data
    data = request.json.get("data", {})
    id = request.json.get("id", "")

    # Do an API call...

    # returns json response
    return jsonify(
        foo=f"Update foo with id of {id} and data of {data}"
    )

@foo_router.route("", methods=["POST"])
def create():
    '''
    Create a foo.
    '''
    # pull data
    data = request.json.get("data", {})

    # Do an API call...

    # returns json response
    return jsonify(
        foo=f"Create foo with data of {data}"
    )