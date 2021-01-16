#!/bin/bash

TOKEN="some-basic-string-token..."

# create foo
curl \
-X POST \
http://localhost:50005/api/v1/foo \
-d '{"data": {"message": "Hi from matan"}}' \
-H "Content-type: application/json" \
-H "authorization: Basic $TOKEN"
