package auth

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	domainmodel "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/domain/model"
	authpkg "github.com/icaroribeiro/new-go-code-challenge-template-2/pkg/auth"
	"github.com/icaroribeiro/new-go-code-challenge-template-2/pkg/customerror"
	authmiddlewarepkg "github.com/icaroribeiro/new-go-code-challenge-template-2/pkg/middleware/auth"
	"gorm.io/gorm"
)

var authDetailsCtxKey = &contextKey{"auth_details"}

type contextKey struct {
	name string
}

// NewContext is the function that returns a new Context that carries auth_details value.
func NewContext(ctx context.Context, auth domainmodel.Auth) context.Context {
	return context.WithValue(ctx, authDetailsCtxKey, auth)
}

// FromContext is the function that returns the auth_details value stored in context, if any.
func FromContext(ctx context.Context) (domainmodel.Auth, bool) {
	raw, ok := ctx.Value(authDetailsCtxKey).(domainmodel.Auth)
	return raw, ok
}

func buildAuth(db *gorm.DB, authN authpkg.IAuth, token *jwt.Token) (domainmodel.Auth, error) {
	auth, err := authN.FetchAuthFromToken(token)
	if err != nil {
		return domainmodel.Auth{}, err
	}

	// Before proceeding is necessary to check if the user who is performing operations is logged
	// based on the authentication details inserted within in the token.
	authAux := domainmodel.Auth{}

	result := db.Find(&authAux, "id=?", auth.ID)
	if result.Error != nil {
		return domainmodel.Auth{}, result.Error
	}

	if authAux.IsEmpty() {
		errorMessage := "you are not logged in, then perform a login to get a token before proceeding"
		return domainmodel.Auth{}, customerror.BadRequest.New(errorMessage)
	}

	if auth.UserID.String() != authAux.UserID.String() {
		errorMessage := "the token's auth_id and user_id are not associated"
		return domainmodel.Auth{}, customerror.BadRequest.New(errorMessage)
	}

	return auth, nil
}

// IsAuthenticated is the function that...
func IsAuthenticated(db *gorm.DB, authN authpkg.IAuth) func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		tokenString, ok := authmiddlewarepkg.FromContext(ctx)
		if !ok || tokenString == "" {
			return nil, customerror.New("failed to get the auth_details value from the request context")
		}

		token, err := authN.DecodeToken(tokenString)
		if err != nil {
			return nil, err
		}

		auth, err := buildAuth(db, authN, token)
		if err != nil {
			return nil, err
		}

		ctx = NewContext(ctx, auth)

		return next(ctx)
	}
}

// CanTokenAlreadyBeRenewed is the function that...
func CanTokenAlreadyBeRenewed(db *gorm.DB, authN authpkg.IAuth, timeBeforeTokenExpTimeInSec int) func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		tokenString, ok := authmiddlewarepkg.FromContext(ctx)
		if !ok || tokenString == "" {
			return nil, customerror.New("failed to get the auth_details value from the request context")
		}

		token, err := authN.ValidateTokenRenewal(tokenString, timeBeforeTokenExpTimeInSec)
		if err != nil {
			return nil, err
		}

		auth, err := buildAuth(db, authN, token)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, authDetailsCtxKey, auth)
		return next(ctx)
	}
}