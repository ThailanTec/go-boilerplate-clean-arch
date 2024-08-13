package mocks

import (
	"reflect"

	"github.com/ThailanTec/challenger/pousada/domain"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

type UserRepositoryMockDB struct {
	ctrl     *gomock.Controller
	recorder *UserRepositoryMockDBRecorder
}

type UserRepositoryMockDBRecorder struct {
	mock *UserRepositoryMockDB
}

func NewUserRepositoryMockDB(ctrl *gomock.Controller) *UserRepositoryMockDB {
	mock := &UserRepositoryMockDB{ctrl: ctrl}
	mock.recorder = &UserRepositoryMockDBRecorder{mock}
	return mock
}

func (m *UserRepositoryMockDB) EXPECT() *UserRepositoryMockDBRecorder {
	return m.recorder
}

func (m *UserRepositoryMockDB) GetUsers() ([]*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *UserRepositoryMockDBRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*UserRepositoryMockDB)(nil).GetUsers))
}

func (m *UserRepositoryMockDB) GetUserByData(document string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByData", document)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *UserRepositoryMockDBRecorder) GetUserByData(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByData", reflect.TypeOf((*UserRepositoryMockDB)(nil).GetUserByData), document)
}

func (m *UserRepositoryMockDB) DeleteUser(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *UserRepositoryMockDBRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*UserRepositoryMockDB)(nil).DeleteUser), id)
}

func (m *UserRepositoryMockDB) UpdateUser(id uuid.UUID, user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", id, user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *UserRepositoryMockDBRecorder) UpdateUser(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*UserRepositoryMockDB)(nil).UpdateUser), id, user)
}

func NewUserRepositoryMock(ctrl *gomock.Controller) *UserRepositoryMockDB {
	mock := &UserRepositoryMockDB{ctrl: ctrl}
	mock.recorder = &UserRepositoryMockDBRecorder{mock}
	return mock
}

func (m *UserRepositoryMockDB) CreateUser(_ *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", gomock.Any())
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *UserRepositoryMockDBRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCall(mr.mock, "CreateUser", user)
}

func (m *UserRepositoryMockDB) GetUserByID(id uuid.UUID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
