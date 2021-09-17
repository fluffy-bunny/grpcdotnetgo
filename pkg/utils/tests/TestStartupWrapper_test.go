package tests

import (
	"testing"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/stretchr/testify/assert"
)

func TestSimpleStartup(t *testing.T) {
	childStartup := &TestStartupWrapper{}
	wrapper := NewTestStartupWrapper(childStartup, nil)

	assert.NotNil(t, wrapper.ChildStartup)
	assert.Nil(t, wrapper.ConfigureServicesOverride)

	wrapper = NewTestStartupWrapper(childStartup, func(builder *di.Builder) {})
	assert.NotNil(t, wrapper.ChildStartup)
	assert.NotNil(t, wrapper.ConfigureServicesOverride)
}
