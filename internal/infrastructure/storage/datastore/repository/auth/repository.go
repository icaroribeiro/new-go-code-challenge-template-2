package auth

import (
	"strings"

	domainmodel "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/domain/model"
	authdatastorerepository "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/infrastructure/storage/datastore/repository/auth"
	datastoremodel "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/infrastructure/storage/datastore/model"
	"github.com/icaroribeiro/new-go-code-challenge-template-2/pkg/customerror"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

var initDB *gorm.DB

// New is the factory function that encapsulates the implementation related to auth.
func New(db *gorm.DB) authdatastorerepository.IRepository {
	initDB = db
	return &Repository{
		DB: db,
	}
}

// Create is the function that creates an auth in the datastore.
func (r *Repository) Create(auth domainmodel.Auth) (domainmodel.Auth, error) {
	authDatastore := datastoremodel.Auth{}
	authDatastore.FromDomain(auth)

	result := r.DB.Create(&authDatastore)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "auths_user_id_key") {
			loginDatastore := datastoremodel.Login{}

			if result := r.DB.Find(&loginDatastore, "user_id=?", authDatastore.UserID); result.Error != nil {
				return domainmodel.Auth{}, result.Error
			}

			if result.RowsAffected == 0 && loginDatastore.IsEmpty() {
				return domainmodel.Auth{}, customerror.NotFound.Newf("the login record with user id %s was not found", authDatastore.UserID)
			}

			return domainmodel.Auth{}, customerror.Conflict.Newf("The user with id %s is already logged in", authDatastore.UserID)
		}

		return domainmodel.Auth{}, result.Error
	}

	return authDatastore.ToDomain(), nil
}

// GetByUserID is the function that gets an auth by user id from the datastore.
func (r *Repository) GetByUserID(userID string) (domainmodel.Auth, error) {
	authDatastore := datastoremodel.Auth{}

	if result := r.DB.Find(&authDatastore, "user_id=?", userID); result.Error != nil {
		return domainmodel.Auth{}, result.Error
	}

	return authDatastore.ToDomain(), nil
}

// Delete is the function that deletes an auth by id from the datastore.
func (r *Repository) Delete(id string) (domainmodel.Auth, error) {
	authDatastore := datastoremodel.Auth{}

	result := r.DB.Find(&authDatastore, "id=?", id)
	if result.Error != nil {
		return domainmodel.Auth{}, result.Error
	}

	if result.RowsAffected == 0 {
		return domainmodel.Auth{}, customerror.NotFound.Newf("the auth with id %s was not found", id)
	}

	if result = r.DB.Delete(&authDatastore); result.Error != nil {
		return domainmodel.Auth{}, result.Error
	}

	if result.RowsAffected == 0 {
		return domainmodel.Auth{}, customerror.NotFound.Newf("the auth with id %s was not deleted", id)
	}

	return authDatastore.ToDomain(), nil
}

// WithDBTrx is the function that enables the repository with datastore transaction.
func (r *Repository) WithDBTrx(dbTrx *gorm.DB) authdatastorerepository.IRepository {
	if dbTrx == nil {
		r.DB = initDB
		return r
	}

	r.DB = dbTrx
	return r
}