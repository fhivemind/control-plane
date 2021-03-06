// Code generated by mockery v2.6.0. DO NOT EDIT.

package automock

import (
	cls "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/cls"

	mock "github.com/stretchr/testify/mock"

	servicemanager "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/servicemanager"
)

// ClsDeprovisioner is an autogenerated mock type for the ClsDeprovisioner type
type ClsDeprovisioner struct {
	mock.Mock
}

// Deprovision provides a mock function with given fields: smClient, request
func (_m *ClsDeprovisioner) Deprovision(smClient servicemanager.Client, request *cls.DeprovisionRequest) error {
	ret := _m.Called(smClient, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(servicemanager.Client, *cls.DeprovisionRequest) error); ok {
		r0 = rf(smClient, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
