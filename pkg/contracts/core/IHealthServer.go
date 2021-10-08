package core

import (
	health "google.golang.org/grpc/health/grpc_health_v1"
)

// IHealthServer contract
type IHealthServer interface {
	health.HealthServer
}
