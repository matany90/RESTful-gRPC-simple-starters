#!/bin/bash

# get foo
curl \
-X GET \
http://localhost:5000/api/v1/foo \
-H "Content-type: application/json"