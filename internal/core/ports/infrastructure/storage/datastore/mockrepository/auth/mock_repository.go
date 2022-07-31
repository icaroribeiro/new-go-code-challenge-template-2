// Code generated by mockery v2.10.0. DO NOT EDIT.

package auth

import (
	entity "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/domain/entity"
	auth "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/infrastructure/storage/datastore/repository/auth"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the IRepository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Repository) Create(_a0 entity.Auth) (entity.Auth, error) {
	ret := _m.Called(_a0)

	var r0 entity.Auth
	if rf, ok := ret.Get(0).(func(entity.Auth) entity.Auth); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Auth) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id string) (entity.Auth, error) {
	ret := _m.Called(id)

	var r0 entity.Auth
	if rf, ok := ret.Get(0).(func(string) entity.Auth); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: userID
func (_m *Repository) GetByUserID(userID string) (entity.Auth, error) {
	ret := _m.Called(userID)

	var r0 entity.Auth
	if rf, ok := ret.Get(0).(func(string) entity.Auth); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(entity.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithDBTrx provides a mock function with given fields: dbTrx
func (_m *Repository) WithDBTrx(dbTrx *gorm.DB) auth.IRepository {
	ret := _m.Called(dbTrx)

	var r0 auth.IRepository
	if rf, ok := ret.Get(0).(func(*gorm.DB) auth.IRepository); ok {
		r0 = rf(dbTrx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(auth.IRepository)
		}
	}

	return r0
}
