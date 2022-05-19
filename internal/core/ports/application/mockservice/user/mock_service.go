// Code generated by mockery v2.10.0. DO NOT EDIT.

package user

import (
	model "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/domain/model"
	user "github.com/icaroribeiro/new-go-code-challenge-template-2/internal/core/ports/application/service/user"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// Service is an autogenerated mock type for the IService type
type Service struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *Service) GetAll() (model.Users, error) {
	ret := _m.Called()

	var r0 model.Users
	if rf, ok := ret.Get(0).(func() model.Users); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithDBTrx provides a mock function with given fields: dbTrx
func (_m *Service) WithDBTrx(dbTrx *gorm.DB) user.IService {
	ret := _m.Called(dbTrx)

	var r0 user.IService
	if rf, ok := ret.Get(0).(func(*gorm.DB) user.IService); ok {
		r0 = rf(dbTrx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(user.IService)
		}
	}

	return r0
}