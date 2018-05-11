#!/bin/bash

#Script to install the activity into the flogo ui before hand and build it so it will be just restart of the flogo web-ui.

docker exec -w /tmp/flogo-web/build/server/local/engines/flogo-web flogo.example.com flogo install -v 0.0.2 github.com/vinayak03/Akatsuki-flogo/activity/hyperledgerFabricChainCode
docker exec -w /tmp/flogo-web/build/server/local/engines/flogo-web flogo.example.com flogo ensure -add github.com/hyperledger/fabric-sdk-go@ba0e035b4a43d42233f0b40a3a35d0eb0e3a1c98
docker exec -w /tmp/flogo-web/build/server/local/engines/flogo-web flogo.example.com flogo build