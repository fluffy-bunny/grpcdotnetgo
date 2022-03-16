package oauth2

import (
	"testing"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	"github.com/rs/zerolog/log"
	suiteTestify "github.com/stretchr/testify/suite"
)

var (
	configAndOr = middleware_oidc.ClaimsConfig{
		AND: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd1",
					Value: "claimAnd1_1",
				},
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd1",
					Value: "claimAnd1_2",
				},
			},
		},
		OR: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr1",
					Value: "claimOr1_1",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr1",
					Value: "claimOr1_2",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
		},
	}
	configAndOrAndTypeOrType = middleware_oidc.ClaimsConfig{
		AND: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd1",
					Value: "claimAnd1_1",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd1",
					Value: "claimAnd1_2",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},

			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd2",
					Value: "random",
				},
				Directive: middleware_oidc.ClaimType,
			},
		},
		OR: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr1",
					Value: "claimOr1_1",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr1",
					Value: "claimOr1_2",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr2",
					Value: "random",
				},
				Directive: middleware_oidc.ClaimType,
			},
		},
	}

	configAndOnly = middleware_oidc.ClaimsConfig{
		AND: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd1",
					Value: "claimAnd1_1",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd1",
					Value: "claimAnd1_2",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
		},
	}
	configAndTypeOnly = middleware_oidc.ClaimsConfig{
		AND: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd1",
					Value: "random",
				},
				Directive: middleware_oidc.ClaimType,
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimAnd2",
					Value: "random",
				},
				Directive: middleware_oidc.ClaimType,
			},
		},
	}
	configOrOnly = middleware_oidc.ClaimsConfig{
		OR: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr1",
					Value: "claimOr1_1",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			}, {
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr1",
					Value: "claimOr1_2",
				},
				Directive: middleware_oidc.ClaimTypeAndValue,
			},
		},
	}
	configOrTypeOnly = middleware_oidc.ClaimsConfig{
		OR: []middleware_oidc.ClaimFact{
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr1",
					Value: "random",
				},
				Directive: middleware_oidc.ClaimType,
			},
			{
				Claim: claimsprincipalContracts.Claim{
					Type:  "claimOr2",
					Value: "random",
				},
				Directive: middleware_oidc.ClaimType,
			},
		},
	}
)

type testSuite struct {
	suiteTestify.Suite
	testCases []struct {
		Desc            string
		Config          *middleware_oidc.ClaimsConfig
		ClaimsPrincipal claimsprincipalContracts.IClaimsPrincipal
		expected        bool
	}
}

// before each test
func (suite *testSuite) SetupTest() {
	suite.testCases = []struct {
		Desc            string
		Config          *middleware_oidc.ClaimsConfig
		ClaimsPrincipal claimsprincipalContracts.IClaimsPrincipal
		expected        bool
	}{
		{
			"TestFullAndOrAndTypeOrTypeTrue",
			&configAndOrAndTypeOrType,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"claimAnd1_1", "claimAnd1_2"},
				"claimAnd2": []interface{}{"test"},
				"claimOr1":  []interface{}{"claimOr1_1", "claimOr1_2"},
				"claimOr2":  []interface{}{"test"},
				"random":    []interface{}{"a", "d"},
			}),
			true,
		},
		{
			"TestFullAndOrAndTypeOrTypeFalse",
			&configAndOrAndTypeOrType,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"claimAnd1_1", "claimAnd1_2"},
				"blah":      []interface{}{"test"},
				"claimOr1":  []interface{}{"claimOr1_1", "claimOr1_2"},
				"claimOr2":  []interface{}{"test"},
				"random":    []interface{}{"a", "d"},
			}),
			false,
		},
		{
			"TestAndTypeOnlyTrue",
			&configAndTypeOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"test"},
				"claimAnd2": []interface{}{"test"},
				"random":    []interface{}{"a", "d"},
			}),
			true,
		},
		{
			"TestAndTypeOnlyFalse",
			&configAndTypeOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimAnd1": []interface{}{"test"},
				"bla":       []interface{}{"test"},
				"random":    []interface{}{"a", "d"},
			}),
			false,
		},
		{
			"TestOrTypeOnlyTrue",
			&configOrTypeOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"claimOr1": []interface{}{"test"},
				"claimOr2": []interface{}{"test"},
				"random":   []interface{}{"a", "d"},
			}),
			true,
		},
		{
			"TestOrTypeOnlyFalse",
			&configOrTypeOnly,
			ClaimsPrincipalFromClaimsMap(map[string]interface{}{
				"random": []interface{}{"a", "d"},
			}),
			false,
		},
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

// TestValidation ...
// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *testSuite) TestValidation() {
	for _, tc := range suite.testCases {
		actual := validate(log.Debug(), *tc.Config, tc.ClaimsPrincipal)
		suite.Equal(actual, tc.expected, tc.Desc)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestValidationTestSuite(t *testing.T) {
	suiteTestify.Run(t, new(testSuite))
}
