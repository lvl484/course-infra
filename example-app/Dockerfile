FROM golang:1.13 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.13 as builder
ENV CGO_ENABLED 0

RUN mkdir -p /opt/resource/

COPY --from=modules /go/pkg /go/pkg

WORKDIR /opt/resource/
COPY api             api
COPY cmd             cmd
COPY discovery       discovery
COPY metrics         metrics
COPY go.*            ./

WORKDIR /opt/resource/cmd/app
RUN go build -o /opt/services/example-app .

FROM alpine:3.7
COPY --from=builder /opt/services/example-app /opt/services/example-app
CMD /opt/services/example-app
