// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

// copied and modified from datadog/dd-trace-go/internal/env.go
package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	// add zerolog
	"github.com/rs/zerolog/log"
)

// BoolEnv returns the parsed boolean value of an environment variable, or
// def otherwise.
func BoolEnv(key string, def bool) bool {
	vv, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	v, err := strconv.ParseBool(vv)
	if err != nil {
		log.Warn().Msgf("Non-boolean value for env var %s, defaulting to %t. Parse failed with error: %v", key, def, err)
		return def
	}
	return v
}

// IntEnv returns the parsed int value of an environment variable, or
// def otherwise.
func IntEnv(key string, def int) int {
	vv, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	v, err := strconv.Atoi(vv)
	if err != nil {
		log.Warn().Msgf("Non-integer value for env var %s, defaulting to %d. Parse failed with error: %v", key, def, err)
		return def
	}
	return v
}

// DurationEnv returns the parsed duration value of an environment variable, or
// def otherwise.
func DurationEnv(key string, def time.Duration) time.Duration {
	vv, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	v, err := time.ParseDuration(vv)
	if err != nil {
		log.Warn().Msgf("Non-duration value for env var %s, defaulting to %d. Parse failed with error: %v", key, def, err)
		return def
	}
	return v
}

// BindAddress returns "127.0.0.1:port" when DEV_BIND_LOCALHOST=true,
// otherwise ":port". Binding to localhost prevents Windows Firewall prompts.
func BindAddress(port int) string {
	if BoolEnv("DEV_BIND_LOCALHOST", false) {
		return fmt.Sprintf("127.0.0.1:%d", port)
	}
	return fmt.Sprintf(":%d", port)
}
