package wellknown

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
)

// MetaDataFilter list wellknown grpc metadata
var MetaDataFilter = []string{
	":authority",
	"authorization",
	"content-type",
	"user-agent",
	"grpc-accept-encoding",
	"accept-encoding",
	XCorrelationIDName,
	LogCorrelationIDName,
	XSpanName,
	LogSpanName,
	XParentName,
	LogParentName,
	XRequestID,
}

// MetaDataFilterSet ...
var MetaDataFilterSet = hashset.NewStringSet(MetaDataFilter...)
