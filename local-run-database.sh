#!/bin/bash

docker run -d --name postgres -p 5432:5432 -v $(pwd)/db:/docker-entrypoint-initdb.d -e POSTGRES_DB=callWs -e POSTGRES_USER=callWsAdmin -e POSTGRES_PASSWORD=batatas123 postgres
