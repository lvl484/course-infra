package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // nolint: gosec
	"strconv"
	"time"

	zipkinExporter "contrib.go.opencensus.io/exporter/zipkin"
	"github.com/kelseyhightower/envconfig"
	"github.com/lvl484/course-infra/example-app/api"
	discovery "github.com/lvl484/course-infra/example-app/discovery/consul"
	appMetrics "github.com/lvl484/course-infra/example-app/metrics"
	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTPReporter "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

var (
	version = "unknown"
)

func main() {
	log.Print("start example.app")

	appConfig := struct {
		ServiceName   string `envconfig:"SERVICE_NAME"                   required:"true"`
		BindIP        string `envconfig:"BIND_IP"                        required:"true"`
		BindPort      int    `envconfig:"BIND_PORT"                      required:"true"`
		BindDebugPort int    `envconfig:"BIND_DEBUG_PORT"`
		Consul        struct {
			Address string `envconfig:"CONSUL_ADDRESS"                   required:"true"`
		}
		Zipkin struct {
			Address      string  `envconfig:"ZIPKIN_ADDRESS"             required:"true"`
			ReportingAPI string  `envconfig:"ZIPKIN_REPORTING_API"       required:"true"`
			Fraction     float64 `envconfig:"ZIPKIN_FRACTION"            required:"true"`
		}
	}{}

	envconfig.MustProcess("", &appConfig)

	localEndpoint, err := openzipkin.NewEndpoint(
		appConfig.ServiceName,
		appConfig.Zipkin.Address,
	)
	if err != nil {
		log.Fatalf("can't initialize zipkin tracing: %v", err)
	}

	reporter := zipkinHTTPReporter.NewReporter(
		"http://" + appConfig.Zipkin.Address + appConfig.Zipkin.ReportingAPI,
	)
	exporter := zipkinExporter.NewExporter(reporter, localEndpoint)
	sampler := trace.ProbabilitySampler(appConfig.Zipkin.Fraction)

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: sampler})

	// Service discovery =======================
	consulClient, err := discovery.NewClient("http://" + appConfig.Consul.Address)
	if err != nil {
		log.Fatalf("can't create consul client: %s", err)
	}

	agentCfg := discovery.AgentConfig(
		appConfig.ServiceName,
		appConfig.BindPort,
		"/health",
	)

	err = discovery.ServiceRegister(consulClient, agentCfg)
	if err != nil {
		log.Fatalf("can't register service discovery: %s", err)
	}
	// =========================================

	// Prometheus exporter =====================
	prometheusExporter, err := appMetrics.NewExporter(appConfig.ServiceName)
	if err != nil {
		log.Fatalf("can't create prometheus exporter: %s", err)
	}
	// =========================================

	// Register opencensus view
	err = view.Register(ochttp.ServerRequestCountView)
	if err != nil {
		log.Fatalf("can't register tracing views: %s", err)
	}

	webAPI := http.Server{
		Addr:    appConfig.BindIP + ":" + strconv.Itoa(appConfig.BindPort),
		Handler: api.API(version, prometheusExporter),
		// timeouts to protect from broken networks
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	// debug and profile =======================
	if appConfig.BindDebugPort != 0 {
		go func() {
			log.Printf("debug listening address: %s:%d\n", appConfig.BindIP, appConfig.BindDebugPort)
			log.Printf(
				"debug listening closed: %v\n",
				http.ListenAndServe(
					appConfig.BindIP+":"+strconv.Itoa(appConfig.BindDebugPort),
					http.DefaultServeMux,
				),
			)
		}()
	}
	// =========================================

	log.Printf("http listening on: %s", webAPI.Addr)
	log.Printf("http listening closed: %v\n", webAPI.ListenAndServe())
}
