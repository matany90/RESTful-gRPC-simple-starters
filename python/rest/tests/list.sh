#!/bin/bash

TOKEN="some-basic-string-token..."

# list of foos
curl \
-X GET \
http://localhost:50005/api/v1/foo \
-H "Content-type: application/json" \
-H "authorization: Basic $TOKEN"