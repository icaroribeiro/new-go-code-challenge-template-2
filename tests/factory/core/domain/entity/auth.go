package entity

import (
	"github.com/bluele/factory-go/factory"
	domainentity "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/domain/entity"
	uuid "github.com/satori/go.uuid"
)

// NewAuth is the function that returns an instance of the auth's domain model for performing tests.
func NewAuth(args map[string]interface{}) domainentity.Auth {
	authFactory := factory.NewFactory(
		domainentity.Auth{},
	).Attr("ID", func(fArgs factory.Args) (interface{}, error) {
		id := uuid.NewV4()

		if val, ok := args["id"]; ok {
			id = val.(uuid.UUID)
		}

		return id, nil
	}).Attr("UserID", func(fArgs factory.Args) (interface{}, error) {
		userID := uuid.NewV4()

		if val, ok := args["userID"]; ok {
			userID = val.(uuid.UUID)
		}

		return userID, nil
	})

	return authFactory.MustCreate().(domainentity.Auth)
}