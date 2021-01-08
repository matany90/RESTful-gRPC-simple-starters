#!/bin/bash

# get foo
curl \
-X GET \
http://localhost:50005/api/v1/foo/someId \
-H "Content-type: application/json"