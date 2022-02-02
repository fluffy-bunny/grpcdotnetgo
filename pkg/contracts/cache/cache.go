package cache

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ICache,IMemoryCache"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ICache,IMemoryCache

import (
	"time"

	ttlcache "github.com/ReneKroon/ttlcache/v2"
)

type (
	// ICache interface
	ICache interface {
		ttlcache.SimpleCache
		GetOrInsert(k string, adder func() (interface{}, time.Duration, error)) interface{}
	}
	// IMemoryCache interface
	IMemoryCache interface {
		ICache
	}
)

// some generic expire time define.
const (
	// Forever Always exist
	Forever = 0

	// Seconds1 1 second
	Seconds1 = time.Second
	// Seconds2 2 second
	Seconds2 = 2 * time.Second
	// Seconds3 3 second
	Seconds3 = 3 * time.Second
	// Seconds5 5 second
	Seconds5 = 5 * time.Second
	// Seconds6 6 second
	Seconds6 = 6 * time.Second
	// Seconds7 7 second
	Seconds7 = 7 * time.Second
	// Seconds8 8 second
	Seconds8 = 8 * time.Second
	// Seconds9 9 second
	Seconds9 = 9 * time.Second
	// Seconds10 10 second
	Seconds10 = 10 * time.Second
	// Seconds15 15 second
	Seconds15 = 15 * time.Second
	// Seconds20 20 second
	Seconds20 = 20 * time.Second
	// Seconds30 30 second
	Seconds30 = 30 * time.Second

	// OneMinutes 1 minutes
	OneMinutes = 60 * time.Second
	// TwoMinutes 2 minutes
	TwoMinutes = 120 * time.Second
	// ThreeMinutes 3 minutes
	ThreeMinutes = 180 * time.Second
	// FiveMinutes 5 minutes
	FiveMinutes = 300 * time.Second
	// TenMinutes 10 minutes
	TenMinutes = 600 * time.Second
	// FifteenMinutes 15 minutes
	FifteenMinutes = 900 * time.Second
	// HalfHour half an hour
	HalfHour = 1800 * time.Second
	// OneHour 1 hour
	OneHour = 3600 * time.Second
	// TwoHour 2 hours
	TwoHour = 7200 * time.Second
	// ThreeHour 3 hours
	ThreeHour = 10800 * time.Second
	// HalfDay 12 hours(half of the day)
	HalfDay = 43200 * time.Second
	// OneDay 24 hours(1 day)
	OneDay = 86400 * time.Second
	// TwoDay 2 day
	TwoDay = 172800 * time.Second
	// ThreeDay 3 day
	ThreeDay = 259200 * time.Second
	// OneWeek 7 day(one week)
	OneWeek = 604800 * time.Second
)
