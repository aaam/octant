// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/heptio/developer-dash/internal/cluster (interfaces: ClientInterface)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	cluster "github.com/heptio/developer-dash/internal/cluster"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	discovery "k8s.io/client-go/discovery"
	dynamic "k8s.io/client-go/dynamic"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockClientInterface is a mock of ClientInterface interface
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// DiscoveryClient mocks base method
func (m *MockClientInterface) DiscoveryClient() (discovery.DiscoveryInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoveryClient")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoveryClient indicates an expected call of DiscoveryClient
func (mr *MockClientInterfaceMockRecorder) DiscoveryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryClient", reflect.TypeOf((*MockClientInterface)(nil).DiscoveryClient))
}

// DynamicClient mocks base method
func (m *MockClientInterface) DynamicClient() (dynamic.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DynamicClient")
	ret0, _ := ret[0].(dynamic.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DynamicClient indicates an expected call of DynamicClient
func (mr *MockClientInterfaceMockRecorder) DynamicClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DynamicClient", reflect.TypeOf((*MockClientInterface)(nil).DynamicClient))
}

// InfoClient mocks base method
func (m *MockClientInterface) InfoClient() (cluster.InfoInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfoClient")
	ret0, _ := ret[0].(cluster.InfoInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfoClient indicates an expected call of InfoClient
func (mr *MockClientInterfaceMockRecorder) InfoClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfoClient", reflect.TypeOf((*MockClientInterface)(nil).InfoClient))
}

// KubernetesClient mocks base method
func (m *MockClientInterface) KubernetesClient() (kubernetes.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubernetesClient")
	ret0, _ := ret[0].(kubernetes.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubernetesClient indicates an expected call of KubernetesClient
func (mr *MockClientInterfaceMockRecorder) KubernetesClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubernetesClient", reflect.TypeOf((*MockClientInterface)(nil).KubernetesClient))
}

// NamespaceClient mocks base method
func (m *MockClientInterface) NamespaceClient() (cluster.NamespaceInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamespaceClient")
	ret0, _ := ret[0].(cluster.NamespaceInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamespaceClient indicates an expected call of NamespaceClient
func (mr *MockClientInterfaceMockRecorder) NamespaceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamespaceClient", reflect.TypeOf((*MockClientInterface)(nil).NamespaceClient))
}

// RESTClient mocks base method
func (m *MockClientInterface) RESTClient() (rest.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockClientInterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockClientInterface)(nil).RESTClient))
}

// RESTConfig mocks base method
func (m *MockClientInterface) RESTConfig() *rest.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTConfig")
	ret0, _ := ret[0].(*rest.Config)
	return ret0
}

// RESTConfig indicates an expected call of RESTConfig
func (mr *MockClientInterfaceMockRecorder) RESTConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTConfig", reflect.TypeOf((*MockClientInterface)(nil).RESTConfig))
}

// Resource mocks base method
func (m *MockClientInterface) Resource(arg0 schema.GroupKind) (schema.GroupVersionResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resource", arg0)
	ret0, _ := ret[0].(schema.GroupVersionResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resource indicates an expected call of Resource
func (mr *MockClientInterfaceMockRecorder) Resource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resource", reflect.TypeOf((*MockClientInterface)(nil).Resource), arg0)
}

// ResourceExists mocks base method
func (m *MockClientInterface) ResourceExists(arg0 schema.GroupVersionResource) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ResourceExists indicates an expected call of ResourceExists
func (mr *MockClientInterfaceMockRecorder) ResourceExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceExists", reflect.TypeOf((*MockClientInterface)(nil).ResourceExists), arg0)
}
