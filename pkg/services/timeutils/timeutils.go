package timeutils

import (
	"reflect"
	"time"

	contracts_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/timeutils"
	di "github.com/fluffy-bunny/sarulabsdi"
)

type (
	serviceTimeUtils struct {
		Time contracts_timeutils.ITime `inject:""`
	}
)

// AddSingletonITimeUtils ...
func AddSingletonITimeUtils(builder *di.Builder) {
	contracts_timeutils.AddSingletonITimeUtils(builder, reflect.TypeOf(&serviceTimeUtils{}))
}

// NewTimeUtils ...
func NewTimeUtils(timeService contracts_timeutils.ITime) contracts_timeutils.ITimeUtils {
	return &serviceTimeUtils{
		Time: timeService,
	}
}

func (s *serviceTimeUtils) StartOfNextMonthUTC() time.Time {
	now := s.Time.Now().UTC()
	currentYear := now.Year()
	nextYear := currentYear
	currentMonth := now.Month()
	nextMonth := currentMonth + 1

	if nextMonth > time.December {
		// bump to next year
		nextMonth = time.January
		nextYear++
	}
	tt := time.Date(nextYear, nextMonth, 1, 0, 0, 0, 0, time.UTC)
	return tt
}

// StartOfCurrentMonthUTC returns the start of current month in UTC
func (s *serviceTimeUtils) StartOfCurrentMonthUTC() time.Time {
	now := s.Time.Now().UTC()
	currentYear := now.Year()
	nextYear := currentYear
	currentMonth := now.Month()

	tt := time.Date(nextYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
	return tt
}

// StartOfPreviousMonthUTC returns the start of current month in UTC
func (s *serviceTimeUtils) StartOfPreviousMonthUTC() time.Time {
	now := s.Time.Now().UTC()
	currentYear := now.Year()
	prevYear := currentYear
	currentMonth := now.Month()
	prevMonth := currentMonth - 1
	if currentMonth == time.January {
		// fall back to last year
		prevMonth = time.December
		prevYear--
	}

	tt := time.Date(prevYear, prevMonth, 1, 0, 0, 0, 0, time.UTC)
	return tt
}
