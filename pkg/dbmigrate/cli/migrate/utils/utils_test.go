package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSuite struct {
	Original string
	Expected string
}

var (
	testSuites = []testSuite{
		{
			Original: "mongodb://micro:*******@mapped-dev-shard-00-00-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-01-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-02-pri.awzwl.mongodb.net:27017?ssl=true&authSource=admin&replicaSet=atlas-uqxcpd-shard-0",
			Expected: "mongodb://micro:*******@mapped-dev-shard-00-00-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-01-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-02-pri.awzwl.mongodb.net:27017/test?ssl=true&authSource=admin&replicaSet=atlas-uqxcpd-shard-0",
		},
		{
			Original: "mongodb://micro:*******@mapped-dev-shard-00-00-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-01-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-02-pri.awzwl.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=atlas-uqxcpd-shard-0",
			Expected: "mongodb://micro:*******@mapped-dev-shard-00-00-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-01-pri.awzwl.mongodb.net:27017,mapped-dev-shard-00-02-pri.awzwl.mongodb.net:27017/test?ssl=true&authSource=admin&replicaSet=atlas-uqxcpd-shard-0",
		},
		{
			Original: "mongodb://localhost:27017",
			Expected: "mongodb://localhost:27017/test",
		}}
)

func TestFixupMongoConnectionString(t *testing.T) {
	for _, v := range testSuites {
		actual := fixupMongoConnectionString(v.Original, "test")
		// assert equality
		assert.Equal(t, v.Expected, actual, "they should be equal")
	}
}
