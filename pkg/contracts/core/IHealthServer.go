package core

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IHealthServer"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE IHealthServer

import (
	health "google.golang.org/grpc/health/grpc_health_v1"
)

// IHealthServer contract
type IHealthServer interface {
	health.HealthServer
}
