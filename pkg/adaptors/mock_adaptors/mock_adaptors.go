// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/kubelogin/pkg/adaptors (interfaces: Kubeconfig,TokenCacheRepository,CredentialPluginInteraction)

// Package mock_adaptors is a generated GoMock package.
package mock_adaptors

import (
	gomock "github.com/golang/mock/gomock"
	credentialplugin "github.com/int128/kubelogin/pkg/models/credentialplugin"
	kubeconfig "github.com/int128/kubelogin/pkg/models/kubeconfig"
	reflect "reflect"
)

// MockKubeconfig is a mock of Kubeconfig interface
type MockKubeconfig struct {
	ctrl     *gomock.Controller
	recorder *MockKubeconfigMockRecorder
}

// MockKubeconfigMockRecorder is the mock recorder for MockKubeconfig
type MockKubeconfigMockRecorder struct {
	mock *MockKubeconfig
}

// NewMockKubeconfig creates a new mock instance
func NewMockKubeconfig(ctrl *gomock.Controller) *MockKubeconfig {
	mock := &MockKubeconfig{ctrl: ctrl}
	mock.recorder = &MockKubeconfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeconfig) EXPECT() *MockKubeconfigMockRecorder {
	return m.recorder
}

// GetCurrentAuthProvider mocks base method
func (m *MockKubeconfig) GetCurrentAuthProvider(arg0 string, arg1 kubeconfig.ContextName, arg2 kubeconfig.UserName) (*kubeconfig.AuthProvider, error) {
	ret := m.ctrl.Call(m, "GetCurrentAuthProvider", arg0, arg1, arg2)
	ret0, _ := ret[0].(*kubeconfig.AuthProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentAuthProvider indicates an expected call of GetCurrentAuthProvider
func (mr *MockKubeconfigMockRecorder) GetCurrentAuthProvider(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentAuthProvider", reflect.TypeOf((*MockKubeconfig)(nil).GetCurrentAuthProvider), arg0, arg1, arg2)
}

// UpdateAuthProvider mocks base method
func (m *MockKubeconfig) UpdateAuthProvider(arg0 *kubeconfig.AuthProvider) error {
	ret := m.ctrl.Call(m, "UpdateAuthProvider", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthProvider indicates an expected call of UpdateAuthProvider
func (mr *MockKubeconfigMockRecorder) UpdateAuthProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthProvider", reflect.TypeOf((*MockKubeconfig)(nil).UpdateAuthProvider), arg0)
}

// MockTokenCacheRepository is a mock of TokenCacheRepository interface
type MockTokenCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCacheRepositoryMockRecorder
}

// MockTokenCacheRepositoryMockRecorder is the mock recorder for MockTokenCacheRepository
type MockTokenCacheRepositoryMockRecorder struct {
	mock *MockTokenCacheRepository
}

// NewMockTokenCacheRepository creates a new mock instance
func NewMockTokenCacheRepository(ctrl *gomock.Controller) *MockTokenCacheRepository {
	mock := &MockTokenCacheRepository{ctrl: ctrl}
	mock.recorder = &MockTokenCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokenCacheRepository) EXPECT() *MockTokenCacheRepositoryMockRecorder {
	return m.recorder
}

// FindByKey mocks base method
func (m *MockTokenCacheRepository) FindByKey(arg0 string, arg1 credentialplugin.TokenCacheKey) (*credentialplugin.TokenCache, error) {
	ret := m.ctrl.Call(m, "FindByKey", arg0, arg1)
	ret0, _ := ret[0].(*credentialplugin.TokenCache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKey indicates an expected call of FindByKey
func (mr *MockTokenCacheRepositoryMockRecorder) FindByKey(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKey", reflect.TypeOf((*MockTokenCacheRepository)(nil).FindByKey), arg0, arg1)
}

// Save mocks base method
func (m *MockTokenCacheRepository) Save(arg0 string, arg1 credentialplugin.TokenCacheKey, arg2 credentialplugin.TokenCache) error {
	ret := m.ctrl.Call(m, "Save", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockTokenCacheRepositoryMockRecorder) Save(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTokenCacheRepository)(nil).Save), arg0, arg1, arg2)
}

// MockCredentialPluginInteraction is a mock of CredentialPluginInteraction interface
type MockCredentialPluginInteraction struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialPluginInteractionMockRecorder
}

// MockCredentialPluginInteractionMockRecorder is the mock recorder for MockCredentialPluginInteraction
type MockCredentialPluginInteractionMockRecorder struct {
	mock *MockCredentialPluginInteraction
}

// NewMockCredentialPluginInteraction creates a new mock instance
func NewMockCredentialPluginInteraction(ctrl *gomock.Controller) *MockCredentialPluginInteraction {
	mock := &MockCredentialPluginInteraction{ctrl: ctrl}
	mock.recorder = &MockCredentialPluginInteractionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCredentialPluginInteraction) EXPECT() *MockCredentialPluginInteractionMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockCredentialPluginInteraction) Write(arg0 credentialplugin.Output) error {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockCredentialPluginInteractionMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCredentialPluginInteraction)(nil).Write), arg0)
}
