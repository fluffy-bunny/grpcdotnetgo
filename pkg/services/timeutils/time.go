package timeutils

import (
	"reflect"
	"time"

	contracts_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/timeutils"
	mocks_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/mocks/timeutils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/golang/mock/gomock"
)

type (
	serviceTime struct {
	}
)

// AddSingletonITime ...
func AddSingletonITime(builder *di.Builder) {
	contracts_timeutils.AddSingletonITime(builder, reflect.TypeOf(&serviceTime{}))
}
func newTime() contracts_timeutils.ITime {
	return &serviceTime{}
}
func (s *serviceTime) Now() time.Time {
	return time.Now()
}

var (
	// Months ...
	Months = []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
)

// NewMockITimeYearMonthDate ...
func NewMockITimeYearMonthDate(ctrl *gomock.Controller, year int, month time.Month) contracts_timeutils.ITime {
	return NewMockITimeDate(ctrl, year, month, 1, 0, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayDate ...
func NewMockITimeYearMonthDayDate(ctrl *gomock.Controller, year int, month time.Month, day int) contracts_timeutils.ITime {
	return NewMockITimeDate(ctrl, year, month, day, 0, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayHourDate ...
func NewMockITimeYearMonthDayHourDate(ctrl *gomock.Controller, year int, month time.Month, day int, hour int) contracts_timeutils.ITime {
	return NewMockITimeDate(ctrl, year, month, day, hour, 0, 0, 0, time.UTC)
}

// NewMockITimeYearMonthDayHourMinDate ...
func NewMockITimeYearMonthDayHourMinDate(ctrl *gomock.Controller, year int, month time.Month, day int, hour int, min int) contracts_timeutils.ITime {
	return NewMockITimeDate(ctrl, year, month, day, hour, min, 0, 0, time.UTC)
}

// NewMockITimeDate ...
func NewMockITimeDate(ctrl *gomock.Controller, year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) contracts_timeutils.ITime {
	mockITime := mocks_timeutils.NewMockITime(ctrl)
	mockTimeNow := time.Date(year, month, day, hour, min, sec, nsec, loc)
	mockITime.EXPECT().Now().Return(mockTimeNow).AnyTimes()
	return mockITime
}
