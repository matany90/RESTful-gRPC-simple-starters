#!/bin/bash

# create foo
curl \
-X PUT \
http://localhost:50005/api/v1/foo \
-d '{"data": {"message": "Hi from matan"}}' \
-H "Content-type: application/json"
