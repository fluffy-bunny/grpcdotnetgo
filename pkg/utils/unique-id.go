package utils

import (
	"github.com/rs/xid"
)

func GenerateUniqueID() string {
	return xid.New().String()
}
