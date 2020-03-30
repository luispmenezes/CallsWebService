#!/bin/bash

#Build Server Image
docker build -t callws/server:1.0 ./src/CallServer

#Build Client Image
docker build -t callws/client:1.0 ./src/CallClient
