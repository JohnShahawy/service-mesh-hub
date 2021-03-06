// Code generated by MockGen. DO NOT EDIT.
// Source: ./dependencies.go

// Package translation is a generated GoMock package.
package translation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	input "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/input"
	mesh "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/mesh"
	traffictarget "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/traffictarget"
	workload "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/workload"
)

// MockdependencyFactory is a mock of dependencyFactory interface
type MockdependencyFactory struct {
	ctrl     *gomock.Controller
	recorder *MockdependencyFactoryMockRecorder
}

// MockdependencyFactoryMockRecorder is the mock recorder for MockdependencyFactory
type MockdependencyFactoryMockRecorder struct {
	mock *MockdependencyFactory
}

// NewMockdependencyFactory creates a new mock instance
func NewMockdependencyFactory(ctrl *gomock.Controller) *MockdependencyFactory {
	mock := &MockdependencyFactory{ctrl: ctrl}
	mock.recorder = &MockdependencyFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdependencyFactory) EXPECT() *MockdependencyFactoryMockRecorder {
	return m.recorder
}

// makeMeshTranslator mocks base method
func (m *MockdependencyFactory) makeMeshTranslator(ctx context.Context, in input.Snapshot) mesh.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "makeMeshTranslator", ctx, in)
	ret0, _ := ret[0].(mesh.Translator)
	return ret0
}

// makeMeshTranslator indicates an expected call of makeMeshTranslator
func (mr *MockdependencyFactoryMockRecorder) makeMeshTranslator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "makeMeshTranslator", reflect.TypeOf((*MockdependencyFactory)(nil).makeMeshTranslator), ctx, in)
}

// makeWorkloadTranslator mocks base method
func (m *MockdependencyFactory) makeWorkloadTranslator(ctx context.Context, in input.Snapshot) workload.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "makeWorkloadTranslator", ctx, in)
	ret0, _ := ret[0].(workload.Translator)
	return ret0
}

// makeWorkloadTranslator indicates an expected call of makeWorkloadTranslator
func (mr *MockdependencyFactoryMockRecorder) makeWorkloadTranslator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "makeWorkloadTranslator", reflect.TypeOf((*MockdependencyFactory)(nil).makeWorkloadTranslator), ctx, in)
}

// makeTrafficTargetTranslator mocks base method
func (m *MockdependencyFactory) makeTrafficTargetTranslator(ctx context.Context) traffictarget.Translator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "makeTrafficTargetTranslator", ctx)
	ret0, _ := ret[0].(traffictarget.Translator)
	return ret0
}

// makeTrafficTargetTranslator indicates an expected call of makeTrafficTargetTranslator
func (mr *MockdependencyFactoryMockRecorder) makeTrafficTargetTranslator(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "makeTrafficTargetTranslator", reflect.TypeOf((*MockdependencyFactory)(nil).makeTrafficTargetTranslator), ctx)
}
