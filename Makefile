# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.24
ALPINE          := alpine:3.21
KIND            := kindest/node:v1.32.0
POSTGRES        := postgres:17.3
GRAFANA         := grafana/grafana:11.5.0
PROMETHEUS      := prom/prometheus:v3.1.0
TEMPO           := grafana/tempo:2.7.0
LOKI            := grafana/loki:3.4.0
PROMTAIL        := grafana/promtail:3.4.0

KIND_CLUSTER    := akhirah-cluster
NAMESPACE       := akhirah-system
AKHIRAH_APP     := akhirah
AUTH_APP        := auth
BASE_IMAGE_NAME := dawahlabs
VERSION         := 0.0.1
AKHIRAH_IMAGE   := $(BASE_IMAGE_NAME)/$(AKHIRAH_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)


# ==============================================================================
# Building containers

build: akhirah

akhirah:
	docker build \
		-f zarf/docker/dockerfile.akhirah \
		-t $(AKHIRAH_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

run:
	@go run api/services/akhirah/main.go | go run api/tooling/logfmt/main.go