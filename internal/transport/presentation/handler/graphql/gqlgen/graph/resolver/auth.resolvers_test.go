package resolver_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	fake "github.com/brianvoe/gofakeit/v5"
	domainmodel "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/domain/model"
	authmockservice "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/application/mockservice/auth"
	healthcheckmockservice "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/application/mockservice/healthcheck"
	usermockservice "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/application/mockservice/user"
	"github.com/icaroribeiro/new-go-code-challenge-template-2/internal/transport/presentation/handler/graphql/gqlgen/graph/generated"
	resolverpkg "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/transport/presentation/handler/graphql/gqlgen/graph/resolver"
	"github.com/icaroribeiro/new-go-code-challenge-template-2/pkg/customerror"
	domainmodelfactory "github.com/icaroribeiro/new-go-code-challenge-template-2/tests/factory/core/domain/model"
	securitypkgfactory "github.com/icaroribeiro/new-go-code-challenge-template-2/tests/factory/pkg/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

func TestAuthResolversUnit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) TestSignUp() {
	driver := "postgres"
	db, _ := NewMockDB(driver)

	dbTrx := &gorm.DB{}

	credentials := securitypkgfactory.NewCredentials(nil)

	opts := []client.Option{}

	tokenString := ""

	returnArgs := ReturnArgs{}

	ts.Cases = Cases{
		{
			Context: "ItShouldSucceedInSigningUp",
			SetUp: func(t *testing.T) {
				dbTrx = db

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))
				ctx := context.Background()
				opts = append(opts, AddDBTrxToCtx(ctx, dbTrx))

				tokenString = fake.Word()

				returnArgs = ReturnArgs{
					{tokenString, nil},
				}
			},
			WantError: false,
		},
		{
			Context: "ItShouldFailIfItIsNotPossibleToGetTheDatabaseTransactionFromTheRequestContext",
			SetUp: func(t *testing.T) {
				dbTrx = nil

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))

				returnArgs = ReturnArgs{
					{"", nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfTheDatabaseTransactionFromTheRequestContextIsNull",
			SetUp: func(t *testing.T) {
				dbTrx = nil

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))
				ctx := context.Background()
				opts = append(opts, AddDBTrxToCtx(ctx, dbTrx))

				returnArgs = ReturnArgs{
					{"", nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenRegisteringTheCredentials",
			SetUp: func(t *testing.T) {
				dbTrx = db

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))
				ctx := context.Background()
				opts = append(opts, AddDBTrxToCtx(ctx, dbTrx))

				returnArgs = ReturnArgs{
					{"", customerror.New("failed")},
				}

			},
			WantError: true,
		},
	}

	for _, tc := range ts.Cases {
		ts.T().Run(tc.Context, func(t *testing.T) {
			tc.SetUp(t)

			healthCheckService := new(healthcheckmockservice.Service)
			authService := new(authmockservice.Service)
			authService.On("WithDBTrx", dbTrx).Return(authService)
			authService.On("Register", credentials).Return(returnArgs[0]...)
			userService := new(usermockservice.Service)

			resolver := resolverpkg.New(healthCheckService, authService, userService)

			c := generated.Config{Resolvers: resolver}

			c.Directives.UseDBTrxMiddleware = MockSchemaDirective()

			srv := handler.NewDefaultServer(
				generated.NewExecutableSchema(
					c,
				),
			)

			mutation := signUpMutation
			resp := SignUpMutationResponse{}

			cl := client.New(srv)
			err := cl.Post(mutation, &resp, opts...)

			if !tc.WantError {
				assert.Nil(t, err, fmt.Sprintf("Unexpected error: %v.", err))
				assert.NotEmpty(t, resp.SignUp.Token)
				assert.Equal(t, tokenString, resp.SignUp.Token)
			} else {
				assert.NotNil(t, err, "Predicted error lost.")
			}
		})
	}
}

func (ts *TestSuite) TestSignIn() {
	driver := "postgres"
	db, _ := NewMockDB(driver)

	dbTrx := &gorm.DB{}

	credentials := securitypkgfactory.NewCredentials(nil)

	opts := []client.Option{}

	tokenString := ""

	returnArgs := ReturnArgs{}

	ts.Cases = Cases{
		{
			Context: "ItShouldSucceedInSigningIn",
			SetUp: func(t *testing.T) {
				dbTrx = db

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))
				ctx := context.Background()
				opts = append(opts, AddDBTrxToCtx(ctx, dbTrx))

				tokenString = fake.Word()

				returnArgs = ReturnArgs{
					{tokenString, nil},
				}
			},
			WantError: false,
		},
		{
			Context: "ItShouldFailIfItIsNotPossibleToGetTheDatabaseTransactionFromTheRequestContext",
			SetUp: func(t *testing.T) {
				dbTrx = nil

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))

				returnArgs = ReturnArgs{
					{"", nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfTheDatabaseTransactionFromTheRequestContextIsNull",
			SetUp: func(t *testing.T) {
				dbTrx = nil

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))
				ctx := context.Background()
				opts = append(opts, AddDBTrxToCtx(ctx, dbTrx))

				returnArgs = ReturnArgs{
					{"", nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenLoggingIn",
			SetUp: func(t *testing.T) {
				dbTrx = db

				opts = []client.Option{}
				opts = append(opts, client.Var("input", credentials))
				ctx := context.Background()
				opts = append(opts, AddDBTrxToCtx(ctx, dbTrx))

				returnArgs = ReturnArgs{
					{"", customerror.New("failed")},
				}
			},
			WantError: true,
		},
	}

	for _, tc := range ts.Cases {
		ts.T().Run(tc.Context, func(t *testing.T) {
			tc.SetUp(t)

			healthCheckService := new(healthcheckmockservice.Service)
			authService := new(authmockservice.Service)
			authService.On("WithDBTrx", dbTrx).Return(authService)
			authService.On("LogIn", credentials).Return(returnArgs[0]...)
			userService := new(usermockservice.Service)

			resolver := resolverpkg.New(healthCheckService, authService, userService)

			c := generated.Config{Resolvers: resolver}

			c.Directives.UseDBTrxMiddleware = MockSchemaDirective()

			srv := handler.NewDefaultServer(
				generated.NewExecutableSchema(
					c,
				),
			)

			mutation := signInMutation
			resp := SignInMutationResponse{}

			cl := client.New(srv)
			err := cl.Post(mutation, &resp, opts...)

			if !tc.WantError {
				assert.Nil(t, err, fmt.Sprintf("Unexpected error: %v.", err))
				assert.NotEmpty(t, resp.SignIn.Token)
				assert.Equal(t, tokenString, resp.SignIn.Token)
			} else {
				assert.NotNil(t, err, "Predicted error lost.")
			}
		})
	}
}

func (ts *TestSuite) TestRefreshToken() {
	dbTrx := &gorm.DB{}
	dbTrx = nil

	tokenString := fake.Word()

	auth := domainmodelfactory.NewAuth(nil)

	opt := func(bd *client.Request) {}

	returnArgs := ReturnArgs{}

	ts.Cases = Cases{
		{
			Context: "ItShouldSucceedInRefreshingTheToken",
			SetUp: func(t *testing.T) {
				ctx := context.Background()
				opt = AddAuthDetailsToCtx(ctx, auth)

				returnArgs = ReturnArgs{
					{tokenString, nil},
				}
			},
			WantError: false,
		},
		{
			Context: "ItShouldFailIfItIsNotPossibleToGetTheAuthFromTheRequestContext",
			SetUp: func(t *testing.T) {
				opt = func(bd *client.Request) {}

				returnArgs = ReturnArgs{
					{"", nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfTheAuthFromTheRequestContextIsEmpty",
			SetUp: func(t *testing.T) {
				ctx := context.Background()
				opt = AddAuthDetailsToCtx(ctx, domainmodel.Auth{})

				returnArgs = ReturnArgs{
					{"", nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenRefreshingTheToken",
			SetUp: func(t *testing.T) {
				ctx := context.Background()
				opt = AddAuthDetailsToCtx(ctx, auth)

				returnArgs = ReturnArgs{
					{"", customerror.New("failed")},
				}
			},
			WantError: true,
		},
	}

	for _, tc := range ts.Cases {
		ts.T().Run(tc.Context, func(t *testing.T) {
			tc.SetUp(t)

			healthCheckService := new(healthcheckmockservice.Service)
			authService := new(authmockservice.Service)
			authService.On("WithDBTrx", dbTrx).Return(authService)
			authService.On("RenewToken", auth).Return(returnArgs[0]...)
			userService := new(usermockservice.Service)

			resolver := resolverpkg.New(healthCheckService, authService, userService)

			c := generated.Config{Resolvers: resolver}

			c.Directives.UseAuthRenewalMiddleware = MockSchemaDirective()

			srv := handler.NewDefaultServer(
				generated.NewExecutableSchema(
					c,
				),
			)

			mutation := refreshTokenMutation
			resp := RefreshTokenMutationResponse{}

			cl := client.New(srv)
			err := cl.Post(mutation, &resp, opt)

			if !tc.WantError {
				assert.Nil(t, err, fmt.Sprintf("Unexpected error: %v.", err))
				assert.NotEmpty(t, resp.RefreshToken.Token)
				assert.Equal(t, tokenString, resp.RefreshToken.Token)
			} else {
				assert.NotNil(t, err, "Predicted error lost.")
			}
		})
	}
}

func (ts *TestSuite) TestChangePassword() {
	passwords := securitypkgfactory.NewPasswords(nil)

	opt := client.Var("input", passwords)

	dbTrx := &gorm.DB{}
	dbTrx = nil

	message := ""

	auth := domainmodelfactory.NewAuth(nil)

	authDetails := domainmodel.Auth{}

	ctx := context.Background()

	returnArgs := ReturnArgs{}

	ts.Cases = Cases{
		{
			Context: "ItShouldSucceedInResettingThePassword",
			SetUp: func(t *testing.T) {
				authDetails = auth

				message = "the password has been updated successfully"

				returnArgs = ReturnArgs{
					{nil},
				}
			},
			WantError: false,
		},
		{
			Context: "ItShouldFailIfItIsNotPossibleToGetTheAuthFromTheRequestContext",
			SetUp: func(t *testing.T) {
				authDetails = domainmodel.Auth{}

				returnArgs = ReturnArgs{
					{nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenResettingThePassword",
			SetUp: func(t *testing.T) {
				authDetails = auth

				returnArgs = ReturnArgs{
					{customerror.New("failed")},
				}
			},
			WantError: true,
		},
	}

	for _, tc := range ts.Cases {
		ts.T().Run(tc.Context, func(t *testing.T) {
			tc.SetUp(t)

			healthCheckService := new(healthcheckmockservice.Service)
			authService := new(authmockservice.Service)
			authService.On("WithDBTrx", dbTrx).Return(authService)
			authService.On("ModifyPassword", auth.UserID.String(), passwords).Return(returnArgs[0]...)
			userService := new(usermockservice.Service)

			resolver := resolverpkg.New(healthCheckService, authService, userService)

			c := generated.Config{Resolvers: resolver}

			c.Directives.UseAuthMiddleware = MockSchemaDirective()

			srv := handler.NewDefaultServer(
				generated.NewExecutableSchema(
					c,
				),
			)

			mutation := changePasswordMutation
			resp := ChangePasswordMutationResponse{}

			cl := client.New(srv)
			err := cl.Post(mutation, &resp, opt, AddAuthDetailsToCtx(ctx, authDetails))

			if !tc.WantError {
				assert.Nil(t, err, fmt.Sprintf("Unexpected error: %v.", err))
				assert.NotEmpty(t, resp.ChangePassword.Message)
				assert.Equal(t, message, resp.ChangePassword.Message)
			} else {
				assert.NotNil(t, err, "Predicted error lost.")
			}
		})
	}
}

func (ts *TestSuite) TestSignOut() {
	dbTrx := &gorm.DB{}
	dbTrx = nil

	auth := domainmodelfactory.NewAuth(nil)

	opt := func(bd *client.Request) {}

	message := ""

	returnArgs := ReturnArgs{}

	ts.Cases = Cases{
		{
			Context: "ItShouldSucceedInSigningOut",
			SetUp: func(t *testing.T) {
				ctx := context.Background()
				opt = AddAuthDetailsToCtx(ctx, auth)

				message = "you have logged out successfully"

				returnArgs = ReturnArgs{
					{nil},
				}
			},
			WantError: false,
		},
		{
			Context: "ItShouldFailIfItIsNotPossibleToGetTheAuthFromTheRequestContext",
			SetUp: func(t *testing.T) {
				opt = func(bd *client.Request) {}

				returnArgs = ReturnArgs{
					{nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfTheAuthFromTheRequestContextIsEmpty",
			SetUp: func(t *testing.T) {
				ctx := context.Background()
				opt = AddAuthDetailsToCtx(ctx, domainmodel.Auth{})

				returnArgs = ReturnArgs{
					{nil},
				}
			},
			WantError: true,
		},
		{
			Context: "ItShouldFailIfAnErrorOccursWhenSigningOut",
			SetUp: func(t *testing.T) {
				ctx := context.Background()
				opt = AddAuthDetailsToCtx(ctx, auth)

				returnArgs = ReturnArgs{
					{customerror.New("failed")},
				}
			},
			WantError: true,
		},
	}

	for _, tc := range ts.Cases {
		ts.T().Run(tc.Context, func(t *testing.T) {
			tc.SetUp(t)

			healthCheckService := new(healthcheckmockservice.Service)
			authService := new(authmockservice.Service)
			authService.On("WithDBTrx", dbTrx).Return(authService)
			authService.On("LogOut", auth.ID.String()).Return(returnArgs[0]...)
			userService := new(usermockservice.Service)

			resolver := resolverpkg.New(healthCheckService, authService, userService)

			c := generated.Config{Resolvers: resolver}

			c.Directives.UseAuthMiddleware = MockSchemaDirective()

			srv := handler.NewDefaultServer(
				generated.NewExecutableSchema(
					c,
				),
			)

			mutation := signOutMutation
			resp := SignOutMutationResponse{}

			cl := client.New(srv)
			err := cl.Post(mutation, &resp, opt)

			if !tc.WantError {
				assert.Nil(t, err, fmt.Sprintf("Unexpected error: %v.", err))
				assert.NotEmpty(t, resp.SignOut.Message)
				assert.Equal(t, message, resp.SignOut.Message)
			} else {
				assert.NotNil(t, err, "Predicted error lost.")
			}
		})
	}
}
