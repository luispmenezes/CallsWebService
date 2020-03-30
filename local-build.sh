#!/bin/bash

#Build Server
cd src/CallServer; go build ; cd - > /dev/null
mv src/CallServer/CallServer bin/

#Build Client
cd src/CallClient; go build ; cd - > /dev/null
mv src/CallClient/CallClient bin/
