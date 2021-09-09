package oauth2

import (
	"testing"

	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	"github.com/stretchr/testify/suite"
)

var (
	configAndOr = middleware_oidc.ClaimsConfig{
		AND: []middleware_oidc.Claim{
			{
				Type:  "claimAnd1",
				Value: "claimAnd1_1",
			},
			{
				Type:  "claimAnd1",
				Value: "claimAnd1_2",
			},
		},
		OR: []middleware_oidc.Claim{
			{
				Type:  "claimOr1",
				Value: "claimOr1_1",
			},
			{
				Type:  "claimOr1",
				Value: "claimOr1_2",
			},
		},
	}

	configAndOnly = middleware_oidc.ClaimsConfig{
		AND: []middleware_oidc.Claim{
			{
				Type:  "claimAnd1",
				Value: "claimAnd1_1",
			},
			{
				Type:  "claimAnd1",
				Value: "claimAnd1_2",
			},
		},
	}
	configOrOnly = middleware_oidc.ClaimsConfig{
		OR: []middleware_oidc.Claim{
			{
				Type:  "claimOr1",
				Value: "claimOr1_1",
			},
			{
				Type:  "claimOr1",
				Value: "claimOr1_2",
			},
		},
	}
)

type testSuite struct {
	suite.Suite
	testCases []struct {
		Desc            string
		Config          *middleware_oidc.ClaimsConfig
		ClaimsPrincipal *ClaimsPrincipal
		expected        bool
	}
}

// before each test
func (suite *testSuite) SetupTest() {

	suite.testCases = []struct {
		Desc            string
		Config          *middleware_oidc.ClaimsConfig
		ClaimsPrincipal *ClaimsPrincipal
		expected        bool
	}{

		{
			"TestFullAndOrTrue",
			&configAndOr,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"claimAnd1_1", "claimAnd1_2"},
				"claimOr1":  []interface{}{"claimOr1_1", "claimOr1_2"},
				"random":    []interface{}{"a", "d"},
			}),
			true,
		},
		{
			"TestFullAndOrFalse",
			&configAndOr,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"junk", "claimAnd1_2"},
				"claimOr1":  []interface{}{"claimOr1_1", "claimOr1_2"},
			}),
			false,
		},
		{
			"TestFullAndOrFalse2",
			&configAndOr,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"junk", "claimAnd1_2"},
			}),
			false,
		},
		{
			"TestFullAndOnlyTrue",
			&configAndOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"claimAnd1_1", "claimAnd1_2"},
			}),
			true,
		},
		{
			"TestFullAndOnlyFalse",
			&configAndOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"junk", "claimAnd1_2"},
			}),
			false,
		},
		{
			"TestFullOrOnlyTrue",
			&configOrOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimOr1": []interface{}{"claimOr1_1", "claimOr1_2"},
			}),
			true,
		},
		{
			"TestFullOrOnlyFalse",
			&configOrOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimOr1": []interface{}{"junk", "junk2"},
			}),
			false,
		},
		{
			"TestFullOrOnlyFalse2",
			&configOrOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{}),
			false,
		},
	}
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *testSuite) TestValidation() {
	for _, tc := range suite.testCases {
		actual := validate(*tc.Config, tc.ClaimsPrincipal)
		suite.Equal(actual, tc.expected, tc.Desc)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestValidationTestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
