// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "rena-server-v2/internal/account/model"
)

// AccountDB is an autogenerated mock type for the AccountDB type
type AccountDB struct {
	mock.Mock
}

// FindByAddress provides a mock function with given fields: ctx, address
func (_m *AccountDB) FindByAddress(ctx context.Context, address string) (*model.Account, error) {
	ret := _m.Called(ctx, address)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Account); ok {
		r0 = rf(ctx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, account
func (_m *AccountDB) Save(ctx context.Context, account *model.Account) error {
	ret := _m.Called(ctx, account)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, address, account
func (_m *AccountDB) Update(ctx context.Context, address string, account *model.Account) error {
	ret := _m.Called(ctx, address, account)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Account) error); ok {
		r0 = rf(ctx, address, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
