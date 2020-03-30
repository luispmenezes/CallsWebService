#!/bin/bash

# Start Server
./bin/CallServer conf/server.json > server.log 2>&1 &

# Start Client
./bin/CallClient conf/client.json > client.log 2>&1 &
