package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedact(t *testing.T) {
	type Sensitive struct {
		Name     string `json:"name"`
		Password string `json:"password" redact:"true"`
	}
	obj := &Sensitive{
		Name:     "John",
		Password: "secret",
	}
	fmt.Println(PrettyJSON(obj))
	jsonV, _ := json.Marshal(obj)
	fmt.Println(string(jsonV))

	dst := &Sensitive{}
	PrettyPrintRedacted(obj, dst)
	require.NotEqual(t, obj.Password, dst.Password)
}
