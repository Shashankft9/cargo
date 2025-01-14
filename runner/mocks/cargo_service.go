// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	libcnb "github.com/buildpacks/libcnb"
	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// CargoService is an autogenerated mock type for the CargoService type
type CargoService struct {
	mock.Mock
}

// AsBOMEntry provides a mock function with given fields:
func (_m *CargoService) AsBOMEntry() (libcnb.BOMEntry, error) {
	ret := _m.Called()

	var r0 libcnb.BOMEntry
	if rf, ok := ret.Get(0).(func() libcnb.BOMEntry); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(libcnb.BOMEntry)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CargoVersion provides a mock function with given fields:
func (_m *CargoService) CargoVersion() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanCargoHomeCache provides a mock function with given fields:
func (_m *CargoService) CleanCargoHomeCache() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Install provides a mock function with given fields: srcDir, destLayer
func (_m *CargoService) Install(srcDir string, destLayer libcnb.Layer) error {
	ret := _m.Called(srcDir, destLayer)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, libcnb.Layer) error); ok {
		r0 = rf(srcDir, destLayer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstallMember provides a mock function with given fields: memberPath, srcDir, destLayer
func (_m *CargoService) InstallMember(memberPath string, srcDir string, destLayer libcnb.Layer) error {
	ret := _m.Called(memberPath, srcDir, destLayer)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, libcnb.Layer) error); ok {
		r0 = rf(memberPath, srcDir, destLayer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProjectTargets provides a mock function with given fields: srcDir
func (_m *CargoService) ProjectTargets(srcDir string) ([]string, error) {
	ret := _m.Called(srcDir)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(srcDir)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(srcDir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RustVersion provides a mock function with given fields:
func (_m *CargoService) RustVersion() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WorkspaceMembers provides a mock function with given fields: srcDir, destLayer
func (_m *CargoService) WorkspaceMembers(srcDir string, destLayer libcnb.Layer) ([]url.URL, error) {
	ret := _m.Called(srcDir, destLayer)

	var r0 []url.URL
	if rf, ok := ret.Get(0).(func(string, libcnb.Layer) []url.URL); ok {
		r0 = rf(srcDir, destLayer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]url.URL)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, libcnb.Layer) error); ok {
		r1 = rf(srcDir, destLayer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
