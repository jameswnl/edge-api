// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/devices.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	inventory "github.com/redhatinsights/edge-api/pkg/clients/inventory"
	models "github.com/redhatinsights/edge-api/pkg/models"
	reflect "reflect"
)

// MockDeviceServiceInterface is a mock of DeviceServiceInterface interface
type MockDeviceServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceInterfaceMockRecorder
}

// MockDeviceServiceInterfaceMockRecorder is the mock recorder for MockDeviceServiceInterface
type MockDeviceServiceInterfaceMockRecorder struct {
	mock *MockDeviceServiceInterface
}

// NewMockDeviceServiceInterface creates a new mock instance
func NewMockDeviceServiceInterface(ctrl *gomock.Controller) *MockDeviceServiceInterface {
	mock := &MockDeviceServiceInterface{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceServiceInterface) EXPECT() *MockDeviceServiceInterfaceMockRecorder {
	return m.recorder
}

// GetDeviceByID mocks base method
func (m *MockDeviceServiceInterface) GetDeviceByID(deviceID uint) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByID", deviceID)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByID indicates an expected call of GetDeviceByID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceByID(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceByID), deviceID)
}

// GetDeviceByUUID mocks base method
func (m *MockDeviceServiceInterface) GetDeviceByUUID(deviceUUID string) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUID", deviceUUID)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUID indicates an expected call of GetDeviceByUUID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceByUUID), deviceUUID)
}

// GetUpdateAvailableForDeviceByUUID mocks base method
func (m *MockDeviceServiceInterface) GetUpdateAvailableForDeviceByUUID(deviceUUID string) ([]models.ImageUpdateAvailable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateAvailableForDeviceByUUID", deviceUUID)
	ret0, _ := ret[0].([]models.ImageUpdateAvailable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdateAvailableForDeviceByUUID indicates an expected call of GetUpdateAvailableForDeviceByUUID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetUpdateAvailableForDeviceByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateAvailableForDeviceByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetUpdateAvailableForDeviceByUUID), deviceUUID)
}

// GetDeviceImageInfo mocks base method
func (m *MockDeviceServiceInterface) GetDeviceImageInfo(deviceUUID string) (*models.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceImageInfo", deviceUUID)
	ret0, _ := ret[0].(*models.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceImageInfo indicates an expected call of GetDeviceImageInfo
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceImageInfo(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceImageInfo", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceImageInfo), deviceUUID)
}

// GetDeviceDetails mocks base method
func (m *MockDeviceServiceInterface) GetDeviceDetails(deviceUUID string) (*models.DeviceDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceDetails", deviceUUID)
	ret0, _ := ret[0].(*models.DeviceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceDetails indicates an expected call of GetDeviceDetails
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceDetails(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceDetails", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceDetails), deviceUUID)
}

// GetDevices mocks base method
func (m *MockDeviceServiceInterface) GetDevices(params *inventory.Params) (*models.DeviceDetailsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevices", params)
	ret0, _ := ret[0].(*models.DeviceDetailsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevices indicates an expected call of GetDevices
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDevices(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevices", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDevices), params)
}
