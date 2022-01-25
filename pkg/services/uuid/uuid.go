package uuid

import (
	"reflect"

	contracts_timeutils "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/timeutils"
	contracts_uuid "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/uuid"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/segmentio/ksuid"
)

type (
	serviceKSUID struct {
		Time contracts_timeutils.ITime `inject:""`
	}
)

func (s *serviceKSUID) UUID() string {
	d, _ := ksuid.NewRandomWithTime(s.Time.Now())
	return d.String()
}

// AddSingletonIKSUID ...
func AddSingletonIKSUID(builder *di.Builder) {
	contracts_uuid.AddSingletonIKSUID(builder, reflect.TypeOf(&serviceKSUID{}))
}
