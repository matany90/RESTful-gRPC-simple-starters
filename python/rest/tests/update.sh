#!/bin/bash

TOKEN="some-basic-string-token..."

# update foo
curl \
-X PUT \
http://localhost:50005/api/v1/foo \
-d '{"data": {"message": "Hi from matan"}, "id": "someId"}' \
-H "Content-type: application/json" \
-H "authorization: Basic $TOKEN"