#!/bin/bash

# create foo
curl \
-X POST \
http://localhost:5000/api/v1/foo \
-d '{"name": {"foo": "foo", "bar": "bar"}}' \
-H "Content-type: application/json"
