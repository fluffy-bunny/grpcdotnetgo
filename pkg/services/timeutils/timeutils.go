package timeutils

import (
	"reflect"
	"time"

	contracts_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/timeutils"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

type (
	serviceTimeUtils struct {
		Time contracts_timeutils.ITime `inject:""`
	}
)

// AddSingletonITimeUtils ...
func AddSingletonITimeUtils(builder *di.Builder) {
	log.Info().Msg("IoC: AddSingletonITimeUtils")

	contracts_timeutils.AddSingletonITimeUtils(builder, reflect.TypeOf(&serviceTimeUtils{}))
}

// NewTimeUtils ...
func NewTimeUtils(timeService contracts_timeutils.ITime) contracts_timeutils.ITimeUtils {
	return &serviceTimeUtils{
		Time: timeService,
	}
}

// StartOfMonthUTC returns the start of current month in UTC
func (s *serviceTimeUtils) StartOfMonthUTC(offsetMonth int) time.Time {
	now := s.Time.Now()
	currentYear := now.Year()
	nextYear := currentYear
	currentMonth := now.Month()
	tt := time.Date(nextYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
	tt = tt.AddDate(0, offsetMonth, 0)
	return tt
}
