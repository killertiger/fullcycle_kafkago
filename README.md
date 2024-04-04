# Full Cycle - Kafka GO

Based on the course: "Full Cycle 3.0 - Kafka"

## Control Center
http://localhost:9021/


## Setup

```
$docker exec -it gokafka_kafka bash

$ kafka-topics --create --bootstrap-server=localhost:9092 --topic=mytest --partitions=3
Created topic teste.

$ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=mytest
```

## Running the project

Producer
```
$ docker exec -it gokafka bash
$ go run cmd/producer/main.go
```

Consumer
```
$ docker exec -it gokafka bash
$ go run cmd/consumer/main.go 
```