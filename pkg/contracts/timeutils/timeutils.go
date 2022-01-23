package timeutils

import "time"

//go:generate genny   -pkg $GOPACKAGE        -in=../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=ITimeUtils,ITime"

//go:generate mockgen -package=$GOPACKAGE -destination=../../mocks/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/$GOPACKAGE ITimeUtils,ITime

type (

	// ITimeUtils ...
	ITimeUtils interface {
		// StartOfMonthUTC where offsetMonth is 0-based (0 = Current Month)
		StartOfMonthUTC(offsetMonth int) time.Time
	}
	// ITime ...
	ITime interface {
		Now() time.Time
	}
)
