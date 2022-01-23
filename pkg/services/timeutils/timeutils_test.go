package timeutils

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestMockTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockITime := NewMockITimeYearMonthDate(ctrl, 2020, time.January)
	timeUtils := NewTimeUtils(mockITime)

	require.Equal(t, mockITime.Now().Year(), 2020)
	require.Equal(t, mockITime.Now().Month(), time.January)

	mockITime = NewMockITimeYearMonthDayDate(ctrl, 2020, time.January, 2)
	require.Equal(t, mockITime.Now().Year(), 2020)
	require.Equal(t, mockITime.Now().Month(), time.January)
	require.Equal(t, mockITime.Now().Day(), 2)

	mockITime = NewMockITimeYearMonthDayHourDate(ctrl, 2020, time.January, 2, 2)
	require.Equal(t, mockITime.Now().Year(), 2020)
	require.Equal(t, mockITime.Now().Month(), time.January)
	require.Equal(t, mockITime.Now().Day(), 2)
	require.Equal(t, mockITime.Now().Hour(), 2)

	mockITime = NewMockITimeYearMonthDayHourMinDate(ctrl, 2020, time.January, 2, 2, 2)
	require.Equal(t, mockITime.Now().Year(), 2020)
	require.Equal(t, mockITime.Now().Month(), time.January)
	require.Equal(t, mockITime.Now().Day(), 2)
	require.Equal(t, mockITime.Now().Hour(), 2)
	require.Equal(t, mockITime.Now().Minute(), 2)

	fmt.Println(mockITime.Now())
	fmt.Println(mockITime.Now().UTC())
	fmt.Println(timeUtils.StartOfPreviousMonthUTC())
	fmt.Println(timeUtils.StartOfCurrentMonthUTC())
	fmt.Println(timeUtils.StartOfNextMonthUTC())

	tt := newTime()
	require.NotNil(t, tt)
	require.Equal(t, tt.Now().Year(), time.Now().Year())

}
func TestStartOfNextMonthUTC(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, month := range Months {
		mockITime := NewMockITimeYearMonthDate(ctrl, 2020, month)
		timeUtils := NewTimeUtils(mockITime)
		now := mockITime.Now().UTC()
		fmt.Printf("now date is = %s\n", now)
		ttNextMonth := timeUtils.StartOfNextMonthUTC()
		fmt.Printf("next months date is = %s\n", ttNextMonth)
		require.NotEqual(t, ttNextMonth.Month(), now.Month())
		if ttNextMonth.Month() == time.January {
			require.Equal(t, ttNextMonth.Year(), now.Year()+1)
		} else {
			require.Equal(t, ttNextMonth.Year(), now.Year())
		}
		fmt.Println(ttNextMonth.Unix())
	}

}
func TestStartOfCurrentMonthUTC(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, month := range Months {
		mockITime := NewMockITimeYearMonthDate(ctrl, 2020, month)
		timeUtils := NewTimeUtils(mockITime)
		now := mockITime.Now().UTC()
		fmt.Printf("now date is = %s\n", now)
		ttCurrentMonth := timeUtils.StartOfCurrentMonthUTC()
		fmt.Printf("current months date is = %s\n", ttCurrentMonth)
		require.Equal(t, ttCurrentMonth.Month(), now.Month())
		require.Equal(t, ttCurrentMonth.Year(), now.Year())
		fmt.Println(ttCurrentMonth.Unix())
	}
}
func TestStartOfPreviousMonthUTC(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, month := range Months {
		mockITime := NewMockITimeYearMonthDate(ctrl, 2020, month)
		timeUtils := NewTimeUtils(mockITime)
		now := mockITime.Now().UTC()
		fmt.Printf("now date is = %s\n", now)
		ttPrevMonth := timeUtils.StartOfPreviousMonthUTC()
		fmt.Printf("prev months date is = %s\n", ttPrevMonth)

		require.NotEqual(t, ttPrevMonth.Month(), now.Month())
		if ttPrevMonth.Month() == time.December {
			require.Equal(t, ttPrevMonth.Year(), now.Year()-1)
		} else {
			require.Equal(t, ttPrevMonth.Year(), now.Year())
		}
		fmt.Println(ttPrevMonth.Unix())
	}
}
