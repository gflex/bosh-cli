// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudfoundry/bosh-cli/deployment (interfaces: Deployment,Factory,Deployer,Manager,ManagerFactory)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	agentclient "github.com/cloudfoundry/bosh-agent/agentclient"
	blobstore "github.com/cloudfoundry/bosh-cli/blobstore"
	cloud "github.com/cloudfoundry/bosh-cli/cloud"
	deployment "github.com/cloudfoundry/bosh-cli/deployment"
	disk "github.com/cloudfoundry/bosh-cli/deployment/disk"
	instance "github.com/cloudfoundry/bosh-cli/deployment/instance"
	manifest "github.com/cloudfoundry/bosh-cli/deployment/manifest"
	vm "github.com/cloudfoundry/bosh-cli/deployment/vm"
	stemcell "github.com/cloudfoundry/bosh-cli/stemcell"
	ui "github.com/cloudfoundry/bosh-cli/ui"
	gomock "github.com/golang/mock/gomock"
)

// MockDeployment is a mock of Deployment interface.
type MockDeployment struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentMockRecorder
}

// MockDeploymentMockRecorder is the mock recorder for MockDeployment.
type MockDeploymentMockRecorder struct {
	mock *MockDeployment
}

// NewMockDeployment creates a new mock instance.
func NewMockDeployment(ctrl *gomock.Controller) *MockDeployment {
	mock := &MockDeployment{ctrl: ctrl}
	mock.recorder = &MockDeploymentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployment) EXPECT() *MockDeploymentMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeployment) Delete(arg0 bool, arg1 ui.Stage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeploymentMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeployment)(nil).Delete), arg0, arg1)
}

// Start mocks base method.
func (m *MockDeployment) Start(arg0 ui.Stage, arg1 manifest.Update) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockDeploymentMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDeployment)(nil).Start), arg0, arg1)
}

// Stop mocks base method.
func (m *MockDeployment) Stop(arg0 bool, arg1 ui.Stage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockDeploymentMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDeployment)(nil).Stop), arg0, arg1)
}

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewDeployment mocks base method.
func (m *MockFactory) NewDeployment(arg0 []instance.Instance, arg1 []disk.Disk, arg2 []stemcell.CloudStemcell) deployment.Deployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeployment", arg0, arg1, arg2)
	ret0, _ := ret[0].(deployment.Deployment)
	return ret0
}

// NewDeployment indicates an expected call of NewDeployment.
func (mr *MockFactoryMockRecorder) NewDeployment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeployment", reflect.TypeOf((*MockFactory)(nil).NewDeployment), arg0, arg1, arg2)
}

// MockDeployer is a mock of Deployer interface.
type MockDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerMockRecorder
}

// MockDeployerMockRecorder is the mock recorder for MockDeployer.
type MockDeployerMockRecorder struct {
	mock *MockDeployer
}

// NewMockDeployer creates a new mock instance.
func NewMockDeployer(ctrl *gomock.Controller) *MockDeployer {
	mock := &MockDeployer{ctrl: ctrl}
	mock.recorder = &MockDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployer) EXPECT() *MockDeployerMockRecorder {
	return m.recorder
}

// Deploy mocks base method.
func (m *MockDeployer) Deploy(arg0 cloud.Cloud, arg1 manifest.Manifest, arg2 stemcell.CloudStemcell, arg3 vm.Manager, arg4 blobstore.Blobstore, arg5 bool, arg6 ui.Stage) (deployment.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(deployment.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deploy indicates an expected call of Deploy.
func (mr *MockDeployerMockRecorder) Deploy(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeployer)(nil).Deploy), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Cleanup mocks base method.
func (m *MockManager) Cleanup(arg0 ui.Stage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup.
func (mr *MockManagerMockRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockManager)(nil).Cleanup), arg0)
}

// FindCurrent mocks base method.
func (m *MockManager) FindCurrent() (deployment.Deployment, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCurrent")
	ret0, _ := ret[0].(deployment.Deployment)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindCurrent indicates an expected call of FindCurrent.
func (mr *MockManagerMockRecorder) FindCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCurrent", reflect.TypeOf((*MockManager)(nil).FindCurrent))
}

// MockManagerFactory is a mock of ManagerFactory interface.
type MockManagerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockManagerFactoryMockRecorder
}

// MockManagerFactoryMockRecorder is the mock recorder for MockManagerFactory.
type MockManagerFactoryMockRecorder struct {
	mock *MockManagerFactory
}

// NewMockManagerFactory creates a new mock instance.
func NewMockManagerFactory(ctrl *gomock.Controller) *MockManagerFactory {
	mock := &MockManagerFactory{ctrl: ctrl}
	mock.recorder = &MockManagerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerFactory) EXPECT() *MockManagerFactoryMockRecorder {
	return m.recorder
}

// NewManager mocks base method.
func (m *MockManagerFactory) NewManager(arg0 cloud.Cloud, arg1 agentclient.AgentClient, arg2 blobstore.Blobstore) deployment.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewManager", arg0, arg1, arg2)
	ret0, _ := ret[0].(deployment.Manager)
	return ret0
}

// NewManager indicates an expected call of NewManager.
func (mr *MockManagerFactoryMockRecorder) NewManager(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewManager", reflect.TypeOf((*MockManagerFactory)(nil).NewManager), arg0, arg1, arg2)
}
