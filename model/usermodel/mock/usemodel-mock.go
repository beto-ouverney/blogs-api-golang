// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/beto-ouverney/blogs-api-golang/model/usermodel (interfaces: IUserModel)

// Package mock_usermodel is a generated GoMock package.
package mock

import (
        context "context"
        reflect "reflect"

        entities "github.com/beto-ouverney/blogs-api-golang/entities"
        errors "github.com/beto-ouverney/blogs-api-golang/errors"
        gomock "github.com/golang/mock/gomock"
)

// MockIUserModel is a mock of IUserModel interface.
type MockIUserModel struct {
        ctrl     *gomock.Controller
        recorder *MockIUserModelMockRecorder
}

// MockIUserModelMockRecorder is the mock recorder for MockIUserModel.
type MockIUserModelMockRecorder struct {
        mock *MockIUserModel
}

// NewMockIUserModel creates a new mock instance.
func NewMockIUserModel(ctrl *gomock.Controller) *MockIUserModel {
        mock := &MockIUserModel{ctrl: ctrl}
        mock.recorder = &MockIUserModelMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserModel) EXPECT() *MockIUserModelMockRecorder {
        return m.recorder
}

// AddUser mocks base method.
func (m *MockIUserModel) AddUser(arg0 context.Context, arg1 *entities.User) (*entities.User, *errors.CustomError) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
        ret0, _ := ret[0].(*entities.User)
        ret1, _ := ret[1].(*errors.CustomError)
        return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockIUserModelMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockIUserModel)(nil).AddUser), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockIUserModel) GetAllUsers(arg0 context.Context) ([]entities.UserWithoutPassword, *errors.CustomError) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAllUsers", arg0)
        ret0, _ := ret[0].([]entities.UserWithoutPassword)
        ret1, _ := ret[1].(*errors.CustomError)
        return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserModelMockRecorder) GetAllUsers(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserModel)(nil).GetAllUsers), arg0)
}

// GetByEmail mocks base method.
func (m *MockIUserModel) GetByEmail(arg0 context.Context, arg1 string) (*entities.User, *errors.CustomError) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetByEmail", arg0, arg1)
        ret0, _ := ret[0].(*entities.User)
        ret1, _ := ret[1].(*errors.CustomError)
        return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockIUserModelMockRecorder) GetByEmail(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockIUserModel)(nil).GetByEmail), arg0, arg1)
}