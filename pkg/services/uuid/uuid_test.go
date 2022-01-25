package uuid

import (
	"testing"
	"time"

	contracts_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/timeutils"
	contracts_uuid "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/uuid"
	services_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/timeutils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKSUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	builder, err := di.NewBuilder(di.App, di.Request, "transient")
	assert.NoError(t, err)
	mockITime := services_timeutils.NewMockITimeYearMonthDate(ctrl, 2020, time.January)
	require.Equal(t, mockITime.Now().Year(), 2020)
	require.Equal(t, mockITime.Now().Month(), time.January)

	contracts_timeutils.AddSingletonITimeByObj(builder, mockITime)
	AddSingletonIKSUID(builder)
	container := builder.Build()
	assert.NoError(t, err)
	assert.NotNil(t, container)

	ksuid := contracts_uuid.GetIKSUIDFromContainer(container)
	require.NotNil(t, ksuid)
	require.True(t, ksuid.UUID() != "")
}
