// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/virtual-garden/pkg/provider (interfaces: InfrastructureProvider,BackupProvider)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockInfrastructureProvider is a mock of InfrastructureProvider interface.
type MockInfrastructureProvider struct {
	ctrl     *gomock.Controller
	recorder *MockInfrastructureProviderMockRecorder
}

// MockInfrastructureProviderMockRecorder is the mock recorder for MockInfrastructureProvider.
type MockInfrastructureProviderMockRecorder struct {
	mock *MockInfrastructureProvider
}

// NewMockInfrastructureProvider creates a new mock instance.
func NewMockInfrastructureProvider(ctrl *gomock.Controller) *MockInfrastructureProvider {
	mock := &MockInfrastructureProvider{ctrl: ctrl}
	mock.recorder = &MockInfrastructureProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfrastructureProvider) EXPECT() *MockInfrastructureProviderMockRecorder {
	return m.recorder
}

// ComputeStorageClassConfiguration mocks base method.
func (m *MockInfrastructureProvider) ComputeStorageClassConfiguration() (string, map[string]string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeStorageClassConfiguration")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(map[string]string)
	return ret0, ret1
}

// ComputeStorageClassConfiguration indicates an expected call of ComputeStorageClassConfiguration.
func (mr *MockInfrastructureProviderMockRecorder) ComputeStorageClassConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeStorageClassConfiguration", reflect.TypeOf((*MockInfrastructureProvider)(nil).ComputeStorageClassConfiguration))
}

// MockBackupProvider is a mock of BackupProvider interface.
type MockBackupProvider struct {
	ctrl     *gomock.Controller
	recorder *MockBackupProviderMockRecorder
}

// MockBackupProviderMockRecorder is the mock recorder for MockBackupProvider.
type MockBackupProviderMockRecorder struct {
	mock *MockBackupProvider
}

// NewMockBackupProvider creates a new mock instance.
func NewMockBackupProvider(ctrl *gomock.Controller) *MockBackupProvider {
	mock := &MockBackupProvider{ctrl: ctrl}
	mock.recorder = &MockBackupProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupProvider) EXPECT() *MockBackupProviderMockRecorder {
	return m.recorder
}

// BucketExists mocks base method.
func (m *MockBackupProvider) BucketExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BucketExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BucketExists indicates an expected call of BucketExists.
func (mr *MockBackupProviderMockRecorder) BucketExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BucketExists", reflect.TypeOf((*MockBackupProvider)(nil).BucketExists), arg0, arg1)
}

// ComputeETCDBackupConfiguration mocks base method.
func (m *MockBackupProvider) ComputeETCDBackupConfiguration(arg0 string) (string, map[string][]byte, []v1.EnvVar) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeETCDBackupConfiguration", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(map[string][]byte)
	ret2, _ := ret[2].([]v1.EnvVar)
	return ret0, ret1, ret2
}

// ComputeETCDBackupConfiguration indicates an expected call of ComputeETCDBackupConfiguration.
func (mr *MockBackupProviderMockRecorder) ComputeETCDBackupConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeETCDBackupConfiguration", reflect.TypeOf((*MockBackupProvider)(nil).ComputeETCDBackupConfiguration), arg0)
}

// CreateBucket mocks base method.
func (m *MockBackupProvider) CreateBucket(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockBackupProviderMockRecorder) CreateBucket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockBackupProvider)(nil).CreateBucket), arg0, arg1, arg2)
}

// DeleteBucket mocks base method.
func (m *MockBackupProvider) DeleteBucket(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockBackupProviderMockRecorder) DeleteBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockBackupProvider)(nil).DeleteBucket), arg0, arg1)
}