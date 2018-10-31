# GoCQLSockets
Communicate with Cassandra using GoCQL and websockets

# Install Cassandra
http://cassandra.apache.org/download/

## Run Cassandra
In your cassandra folder run file 'bin/cassandra.exe'

## Create SSL Certificates
https://medium.com/the-new-control-plane/generating-self-signed-certificates-on-windows-7812a600c2d8

# Run sockets
 
## Setup config files
Edit examples/config/conf.json & server/config/conf.json to your setup

## Run examples
go run main.go --mode example1_GET
go run main.go --mode example1_INC
go run main.go --mode example1_DEC

## Run server
go run main.go --mode server

