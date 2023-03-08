// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"
	database "rena-server-v2/internal/registry/database"

	mock "github.com/stretchr/testify/mock"

	model "rena-server-v2/internal/registry/model"
)

// RegistryDB is an autogenerated mock type for the RegistryDB type
type RegistryDB struct {
	mock.Mock
}

// DeleteRegistryBySlug provides a mock function with given fields: ctx, authorId, slug
func (_m *RegistryDB) DeleteRegistryBySlug(ctx context.Context, authorId uint, slug string) error {
	ret := _m.Called(ctx, authorId, slug)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) error); ok {
		r0 = rf(ctx, authorId, slug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindRegistryBySlug provides a mock function with given fields: ctx, slug
func (_m *RegistryDB) FindRegistryBySlug(ctx context.Context, slug string) (*model.Registry, error) {
	ret := _m.Called(ctx, slug)

	var r0 *model.Registry
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Registry); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRegistries provides a mock function with given fields: ctx, criteria
func (_m *RegistryDB) FindRegistries(ctx context.Context, criteria database.IterateItemCriteria) ([]*model.Item, int64, error) {
	ret := _m.Called(ctx, criteria)

	var r0 []*model.Registry
	if rf, ok := ret.Get(0).(func(context.Context, database.IterateRegistryCriteria) []*model.Item); ok {
		r0 = rf(ctx, criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Registry)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, database.IterateRegistryCriteria) int64); ok {
		r1 = rf(ctx, criteria)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, database.IterateRegistryCriteria) error); ok {
		r2 = rf(ctx, criteria)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RunInTx provides a mock function with given fields: ctx, f
func (_m *RegistryDB) RunInTx(ctx context.Context, f func(context.Context) error) error {
	ret := _m.Called(ctx, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRegistry provides a mock function with given fields: ctx, Registry
func (_m *RegistryDB) SaveRegistry(ctx context.Context, registry *model.Registry) error {
	ret := _m.Called(ctx, registry)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Registry) error); ok {
		r0 = rf(ctx, registry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
