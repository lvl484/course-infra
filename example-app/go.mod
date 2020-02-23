module github.com/lvl484/course-infra/example-app

go 1.13

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/gorilla/mux v1.6.2
	github.com/hashicorp/consul/api v1.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/prometheus/client_golang v1.4.1
	go.opencensus.io v0.22.3
)
