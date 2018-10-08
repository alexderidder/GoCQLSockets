# GoCQLSockets
Communicate with Cassandra using GoCQL and websockets

## Run client
go run main.go --mode client

## Run server
go run main.go --mode server

## Cassandra

Keyspace name = cycling

### Table query

CREATE TABLE crownstone.get_power_factor_and_usage_by_stone_id_and_time (
    stone_id uuid,
    range int,
    pow_fac float,
    usage float, 
    PRIMARY KEY ((stone_id, variable_time), range)
) WITH CLUSTERING ORDER BY (range DESC)

