#!/bin/bash

# create foo
curl \
-X POST \
http://localhost:5000/api/v1/foo \
-d '{"name": "foo"}' \
-H "Content-type: application/json"

# get foo
curl \
-X GET \
http://localhost:5000/api/v1/foo \
-d '{"name": "foo"}' \
-H "Content-type: application/json"

# update foo
curl \
-X UPDATE \
http://localhost:5000/api/v1/foo \
-d '{"name": "foo"}' \
-H "Content-type: application/json"

# delete foo
curl \
-X DELETE \
http://localhost:5000/api/v1/foo \
-d '{"name": "foo"}' \
-H "Content-type: application/json"