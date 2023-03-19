// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/icaroribeiro/go-code-challenge-template-2/pkg/security"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
	UseAuthMiddleware        func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	UseAuthRenewalMiddleware func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	UseDBTrxMiddleware       func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
}

type ComplexityRoot struct {
	AuthPayload struct {
		Token func(childComplexity int) int
	}

	HealthCheck struct {
		Status func(childComplexity int) int
	}

	InfoPayload struct {
		Message func(childComplexity int) int
	}

	Mutation struct {
		ChangePassword func(childComplexity int, input security.Passwords) int
		RefreshToken   func(childComplexity int) int
		SignIn         func(childComplexity int, input security.Credentials) int
		SignOut        func(childComplexity int) int
		SignUp         func(childComplexity int, input security.Credentials) int
	}

	Query struct {
		GetAllUsers    func(childComplexity int) int
		GetHealthCheck func(childComplexity int) int
	}

	User struct {
		ID       func(childComplexity int) int
		Username func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "AuthPayload.token":
		if e.complexity.AuthPayload.Token == nil {
			break
		}

		return e.complexity.AuthPayload.Token(childComplexity), true

	case "HealthCheck.status":
		if e.complexity.HealthCheck.Status == nil {
			break
		}

		return e.complexity.HealthCheck.Status(childComplexity), true

	case "InfoPayload.message":
		if e.complexity.InfoPayload.Message == nil {
			break
		}

		return e.complexity.InfoPayload.Message(childComplexity), true

	case "Mutation.changePassword":
		if e.complexity.Mutation.ChangePassword == nil {
			break
		}

		args, err := ec.field_Mutation_changePassword_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ChangePassword(childComplexity, args["input"].(security.Passwords)), true

	case "Mutation.refreshToken":
		if e.complexity.Mutation.RefreshToken == nil {
			break
		}

		return e.complexity.Mutation.RefreshToken(childComplexity), true

	case "Mutation.signIn":
		if e.complexity.Mutation.SignIn == nil {
			break
		}

		args, err := ec.field_Mutation_signIn_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SignIn(childComplexity, args["input"].(security.Credentials)), true

	case "Mutation.signOut":
		if e.complexity.Mutation.SignOut == nil {
			break
		}

		return e.complexity.Mutation.SignOut(childComplexity), true

	case "Mutation.signUp":
		if e.complexity.Mutation.SignUp == nil {
			break
		}

		args, err := ec.field_Mutation_signUp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SignUp(childComplexity, args["input"].(security.Credentials)), true

	case "Query.getAllUsers":
		if e.complexity.Query.GetAllUsers == nil {
			break
		}

		return e.complexity.Query.GetAllUsers(childComplexity), true

	case "Query.getHealthCheck":
		if e.complexity.Query.GetHealthCheck == nil {
			break
		}

		return e.complexity.Query.GetHealthCheck(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.username":
		if e.complexity.User.Username == nil {
			break
		}

		return e.complexity.User.Username(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCredentials,
		ec.unmarshalInputPasswords,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schema/auth.graphql", Input: `extend type Mutation {
    signUp(input: Credentials!): AuthPayload! @useDBTrxMiddleware
    signIn(input: Credentials!): AuthPayload! @useDBTrxMiddleware
    refreshToken: AuthPayload! @useAuthRenewalMiddleware
    changePassword(input: Passwords!): InfoPayload! @useAuthMiddleware
    signOut: InfoPayload! @useAuthMiddleware
}`, BuiltIn: false},
	{Name: "../schema/authpayload.graphql", Input: `type AuthPayload {
  token: String!
}`, BuiltIn: false},
	{Name: "../schema/credentials.graphql", Input: `input Credentials {
    username: String!
    password: String!
}`, BuiltIn: false},
	{Name: "../schema/directives.graphql", Input: `directive @useDBTrxMiddleware on FIELD_DEFINITION

directive @useAuthMiddleware on FIELD_DEFINITION
directive @useAuthRenewalMiddleware on FIELD_DEFINITION`, BuiltIn: false},
	{Name: "../schema/healthcheck.graphql", Input: `type HealthCheck {
    status: String!
}

extend type Query {
    getHealthCheck: HealthCheck!
}`, BuiltIn: false},
	{Name: "../schema/infopayload.graphql", Input: `type InfoPayload {
  message: String!
}`, BuiltIn: false},
	{Name: "../schema/passwords.graphql", Input: `input Passwords {
	currentPassword: String!
	newPassword: String!
}
`, BuiltIn: false},
	{Name: "../schema/scalars.graphql", Input: `scalar UUID`, BuiltIn: false},
	{Name: "../schema/user.graphql", Input: `type User {
    id: UUID!
    username: String!
}

extend type Query {
    getAllUsers: [User!]! @useAuthMiddleware
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
