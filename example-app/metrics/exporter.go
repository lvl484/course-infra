package metrics

import (
	"fmt"

	"contrib.go.opencensus.io/exporter/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// NewExporter returns an exporter that exports stats to Prometheus.
func NewExporter(namespace string) (*prometheus.Exporter, error) {
	r := stdprometheus.NewRegistry()

	err := r.Register(stdprometheus.NewProcessCollector(stdprometheus.ProcessCollectorOpts{}))
	if err != nil {
		return nil, fmt.Errorf("can't register prometheus process collector: %w", err)
	}

	err = r.Register(stdprometheus.NewGoCollector())
	if err != nil {
		return nil, fmt.Errorf("can't register prometheus go collector: %w", err)
	}

	return prometheus.NewExporter(
		prometheus.Options{
			Namespace: namespace,
			Registry:  r,
		},
	)
}
