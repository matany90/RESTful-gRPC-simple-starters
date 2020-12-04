#!/bin/bash

# delete foo
curl \
-X DELETE \
http://localhost:5000/api/v1/foo \
-H "Content-type: application/json"