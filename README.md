## Start all services:
```bash
docker-compose up -d --build
```

## Kafka quick start
```
docker-compose exec kafka bash
kafka-console-producer.sh --broker-list      127.0.0.1:9092 --topic test
kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic test
```

## Prometheus
```query
rate(example_opencensus_io_http_server_request_count{job="lvl484"}[1m])
```
