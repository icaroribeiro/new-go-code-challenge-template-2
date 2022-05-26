package resolver_test

import (
	"testing"

	authmockservice "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/application/mockservice/auth"
	healthcheckmockservice "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/application/mockservice/healthcheck"
	usermockservice "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/application/mockservice/user"
	resolverpkg "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/transport/presentation/handler/graphql/gqlgen/graph/resolver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestResolverUnit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) TestNew() {
	ts.Cases = Cases{
		{
			Context:   "ItShouldSucceedInSettingUpAResolver",
			WantError: false,
		},
	}

	for _, tc := range ts.Cases {
		ts.T().Run(tc.Context, func(t *testing.T) {
			healthCheckService := new(healthcheckmockservice.Service)
			authService := new(authmockservice.Service)
			userService := new(usermockservice.Service)

			resolver := &resolverpkg.Resolver{HealthCheckService: healthCheckService,
				AuthService: authService,
				UserService: userService,
			}

			returnedResolver := resolverpkg.New(healthCheckService, authService, userService)

			if !tc.WantError {
				assert.Equal(t, resolver, returnedResolver)
			}
		})
	}
}
