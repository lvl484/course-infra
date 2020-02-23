package api

import (
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"net/http"
)

func example(w http.ResponseWriter, r *http.Request) {
	ctx, span := trace.StartSpan(
		r.Context(),
		"api.example",
		trace.WithSampler(trace.AlwaysSample()),
	)

	defer span.End()

	ctx, _ = tag.New(ctx,
		tag.Upsert(ochttp.Host, r.Host),
		tag.Upsert(ochttp.Path, r.URL.Path),
		tag.Upsert(ochttp.Method, r.Method),
	)

	stats.Record(ctx, ochttp.ServerRequestCount.M(1))
}
