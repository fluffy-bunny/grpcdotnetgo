package timeutils

import "time"

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/func-types.go      -out=gen-func-$GOFILE gen "FuncType=TimeNow"
//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE      gen "InterfaceType=ITimeUtils,ITime"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ITimeUtils,ITime

type (

	// ITimeUtils ...
	ITimeUtils interface {
		// StartOfMonthUTC where offsetMonth is 0-based (0 = Current Month)
		StartOfMonthUTC(offsetMonth int) time.Time
		// format is "2006-01-02T15:04:05Z07:00"
		//Format(layout string, t time.Time) string
	}
	// ITime ...
	ITime interface {
		Now() time.Time
	}
	// TimeNow ...
	TimeNow func() time.Time
)
