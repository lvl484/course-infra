# course-infra

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

## Dashboards
|Name      |Address                                       |
|----------|----------------------------------------------|
|Consul    |[http://127.0.0.1:8500](http://127.0.0.1:8500)|
|Grafana   |[http://127.0.0.1:3000](http://127.0.0.1:3000)|
|Zipkin    |[http://127.0.0.1:9411](http://127.0.0.1:9411)|
|Prometheus|[http://127.0.0.1:9090](http://127.0.0.1:9090)|
