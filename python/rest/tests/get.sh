#!/bin/bash

TOKEN="some-basic-string-token..."

# get foo
curl \
-X GET \
http://localhost:50005/api/v1/foo/someId \
-d '{"name": "matan"}' \
-H "Content-type: application/json" \
-H "authorization: Basic $TOKEN"