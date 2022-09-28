// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/useCase/interface.go

// Package useCase is a generated GoMock package.
package useCase

import (
	company "companies/internal/domain/company"
	employee "companies/internal/domain/employee"
	project "companies/internal/domain/project"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockStorage) CreateCompany(ctx context.Context, company *company.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockStorageMockRecorder) CreateCompany(ctx, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockStorage)(nil).CreateCompany), ctx, company)
}

// CreateEmployee mocks base method.
func (m *MockStorage) CreateEmployee(ctx context.Context, employee *employee.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployee", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEmployee indicates an expected call of CreateEmployee.
func (mr *MockStorageMockRecorder) CreateEmployee(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployee", reflect.TypeOf((*MockStorage)(nil).CreateEmployee), ctx, employee)
}

// CreateProject mocks base method.
func (m *MockStorage) CreateProject(ctx context.Context, project *project.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockStorageMockRecorder) CreateProject(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockStorage)(nil).CreateProject), ctx, project)
}

// DeleteCompany mocks base method.
func (m *MockStorage) DeleteCompany(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockStorageMockRecorder) DeleteCompany(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockStorage)(nil).DeleteCompany), ctx, ID)
}

// DeleteEmployee mocks base method.
func (m *MockStorage) DeleteEmployee(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployee indicates an expected call of DeleteEmployee.
func (mr *MockStorageMockRecorder) DeleteEmployee(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockStorage)(nil).DeleteEmployee), ctx, ID)
}

// DeleteProject mocks base method.
func (m *MockStorage) DeleteProject(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockStorageMockRecorder) DeleteProject(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockStorage)(nil).DeleteProject), ctx, ID)
}

// ReadAllEmployeesOfProject mocks base method.
func (m *MockStorage) ReadAllEmployeesOfProject(ctx context.Context, projectID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllEmployeesOfProject", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadAllEmployeesOfProject indicates an expected call of ReadAllEmployeesOfProject.
func (mr *MockStorageMockRecorder) ReadAllEmployeesOfProject(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllEmployeesOfProject", reflect.TypeOf((*MockStorage)(nil).ReadAllEmployeesOfProject), ctx, projectID)
}

// ReadCompanyByID mocks base method.
func (m *MockStorage) ReadCompanyByID(ctx context.Context, ID uuid.UUID) (*company.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCompanyByID", ctx, ID)
	ret0, _ := ret[0].(*company.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCompanyByID indicates an expected call of ReadCompanyByID.
func (mr *MockStorageMockRecorder) ReadCompanyByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCompanyByID", reflect.TypeOf((*MockStorage)(nil).ReadCompanyByID), ctx, ID)
}

// ReadEmployeeByID mocks base method.
func (m *MockStorage) ReadEmployeeByID(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadEmployeeByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadEmployeeByID indicates an expected call of ReadEmployeeByID.
func (mr *MockStorageMockRecorder) ReadEmployeeByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadEmployeeByID", reflect.TypeOf((*MockStorage)(nil).ReadEmployeeByID), ctx, ID)
}

// ReadFiltredEmployeesOfProject mocks base method.
func (m *MockStorage) ReadFiltredEmployeesOfProject(ctx context.Context, projectID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFiltredEmployeesOfProject", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadFiltredEmployeesOfProject indicates an expected call of ReadFiltredEmployeesOfProject.
func (mr *MockStorageMockRecorder) ReadFiltredEmployeesOfProject(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFiltredEmployeesOfProject", reflect.TypeOf((*MockStorage)(nil).ReadFiltredEmployeesOfProject), ctx, projectID)
}

// ReadProjectByID mocks base method.
func (m *MockStorage) ReadProjectByID(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProjectByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadProjectByID indicates an expected call of ReadProjectByID.
func (mr *MockStorageMockRecorder) ReadProjectByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProjectByID", reflect.TypeOf((*MockStorage)(nil).ReadProjectByID), ctx, ID)
}

// ReadProjectsOfCompany mocks base method.
func (m *MockStorage) ReadProjectsOfCompany(ctx context.Context, companyID uuid.UUID) ([]*project.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProjectsOfCompany", ctx, companyID)
	ret0, _ := ret[0].([]*project.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProjectsOfCompany indicates an expected call of ReadProjectsOfCompany.
func (mr *MockStorageMockRecorder) ReadProjectsOfCompany(ctx, companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProjectsOfCompany", reflect.TypeOf((*MockStorage)(nil).ReadProjectsOfCompany), ctx, companyID)
}

// UpdateCompany mocks base method.
func (m *MockStorage) UpdateCompany(ctx context.Context, company *company.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", ctx, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockStorageMockRecorder) UpdateCompany(ctx, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockStorage)(nil).UpdateCompany), ctx, company)
}

// UpdateEmployee mocks base method.
func (m *MockStorage) UpdateEmployee(ctx context.Context, employee *employee.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockStorageMockRecorder) UpdateEmployee(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockStorage)(nil).UpdateEmployee), ctx, employee)
}

// UpdateProject mocks base method.
func (m *MockStorage) UpdateProject(ctx context.Context, project *project.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockStorageMockRecorder) UpdateProject(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockStorage)(nil).UpdateProject), ctx, project)
}

// MockCompany is a mock of Company interface.
type MockCompany struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyMockRecorder
}

// MockCompanyMockRecorder is the mock recorder for MockCompany.
type MockCompanyMockRecorder struct {
	mock *MockCompany
}

// NewMockCompany creates a new mock instance.
func NewMockCompany(ctrl *gomock.Controller) *MockCompany {
	mock := &MockCompany{ctrl: ctrl}
	mock.recorder = &MockCompanyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompany) EXPECT() *MockCompanyMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockCompany) CreateCompany(ctx context.Context, company *company.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockCompanyMockRecorder) CreateCompany(ctx, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockCompany)(nil).CreateCompany), ctx, company)
}

// DeleteCompany mocks base method.
func (m *MockCompany) DeleteCompany(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockCompanyMockRecorder) DeleteCompany(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockCompany)(nil).DeleteCompany), ctx, ID)
}

// ReadCompanyByID mocks base method.
func (m *MockCompany) ReadCompanyByID(ctx context.Context, ID uuid.UUID) (*company.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCompanyByID", ctx, ID)
	ret0, _ := ret[0].(*company.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCompanyByID indicates an expected call of ReadCompanyByID.
func (mr *MockCompanyMockRecorder) ReadCompanyByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCompanyByID", reflect.TypeOf((*MockCompany)(nil).ReadCompanyByID), ctx, ID)
}

// UpdateCompany mocks base method.
func (m *MockCompany) UpdateCompany(ctx context.Context, company *company.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", ctx, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockCompanyMockRecorder) UpdateCompany(ctx, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockCompany)(nil).UpdateCompany), ctx, company)
}

// MockProject is a mock of Project interface.
type MockProject struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMockRecorder
}

// MockProjectMockRecorder is the mock recorder for MockProject.
type MockProjectMockRecorder struct {
	mock *MockProject
}

// NewMockProject creates a new mock instance.
func NewMockProject(ctrl *gomock.Controller) *MockProject {
	mock := &MockProject{ctrl: ctrl}
	mock.recorder = &MockProjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProject) EXPECT() *MockProjectMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockProject) CreateProject(ctx context.Context, project *project.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectMockRecorder) CreateProject(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProject)(nil).CreateProject), ctx, project)
}

// DeleteProject mocks base method.
func (m *MockProject) DeleteProject(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectMockRecorder) DeleteProject(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProject)(nil).DeleteProject), ctx, ID)
}

// ReadProjectByID mocks base method.
func (m *MockProject) ReadProjectByID(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProjectByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadProjectByID indicates an expected call of ReadProjectByID.
func (mr *MockProjectMockRecorder) ReadProjectByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProjectByID", reflect.TypeOf((*MockProject)(nil).ReadProjectByID), ctx, ID)
}

// ReadProjectsOfCompany mocks base method.
func (m *MockProject) ReadProjectsOfCompany(ctx context.Context, companyID uuid.UUID) ([]*project.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProjectsOfCompany", ctx, companyID)
	ret0, _ := ret[0].([]*project.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProjectsOfCompany indicates an expected call of ReadProjectsOfCompany.
func (mr *MockProjectMockRecorder) ReadProjectsOfCompany(ctx, companyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProjectsOfCompany", reflect.TypeOf((*MockProject)(nil).ReadProjectsOfCompany), ctx, companyID)
}

// UpdateProject mocks base method.
func (m *MockProject) UpdateProject(ctx context.Context, project *project.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", ctx, project)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectMockRecorder) UpdateProject(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProject)(nil).UpdateProject), ctx, project)
}

// MockEmployee is a mock of Employee interface.
type MockEmployee struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeMockRecorder
}

// MockEmployeeMockRecorder is the mock recorder for MockEmployee.
type MockEmployeeMockRecorder struct {
	mock *MockEmployee
}

// NewMockEmployee creates a new mock instance.
func NewMockEmployee(ctrl *gomock.Controller) *MockEmployee {
	mock := &MockEmployee{ctrl: ctrl}
	mock.recorder = &MockEmployeeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployee) EXPECT() *MockEmployeeMockRecorder {
	return m.recorder
}

// CreateEmployee mocks base method.
func (m *MockEmployee) CreateEmployee(ctx context.Context, employee *employee.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployee", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEmployee indicates an expected call of CreateEmployee.
func (mr *MockEmployeeMockRecorder) CreateEmployee(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployee", reflect.TypeOf((*MockEmployee)(nil).CreateEmployee), ctx, employee)
}

// DeleteEmployee mocks base method.
func (m *MockEmployee) DeleteEmployee(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployee indicates an expected call of DeleteEmployee.
func (mr *MockEmployeeMockRecorder) DeleteEmployee(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockEmployee)(nil).DeleteEmployee), ctx, ID)
}

// ReadAllEmployeesOfProject mocks base method.
func (m *MockEmployee) ReadAllEmployeesOfProject(ctx context.Context, projectID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllEmployeesOfProject", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadAllEmployeesOfProject indicates an expected call of ReadAllEmployeesOfProject.
func (mr *MockEmployeeMockRecorder) ReadAllEmployeesOfProject(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllEmployeesOfProject", reflect.TypeOf((*MockEmployee)(nil).ReadAllEmployeesOfProject), ctx, projectID)
}

// ReadEmployeeByID mocks base method.
func (m *MockEmployee) ReadEmployeeByID(ctx context.Context, ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadEmployeeByID", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadEmployeeByID indicates an expected call of ReadEmployeeByID.
func (mr *MockEmployeeMockRecorder) ReadEmployeeByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadEmployeeByID", reflect.TypeOf((*MockEmployee)(nil).ReadEmployeeByID), ctx, ID)
}

// ReadFiltredEmployeesOfProject mocks base method.
func (m *MockEmployee) ReadFiltredEmployeesOfProject(ctx context.Context, projectID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFiltredEmployeesOfProject", ctx, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadFiltredEmployeesOfProject indicates an expected call of ReadFiltredEmployeesOfProject.
func (mr *MockEmployeeMockRecorder) ReadFiltredEmployeesOfProject(ctx, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFiltredEmployeesOfProject", reflect.TypeOf((*MockEmployee)(nil).ReadFiltredEmployeesOfProject), ctx, projectID)
}

// UpdateEmployee mocks base method.
func (m *MockEmployee) UpdateEmployee(ctx context.Context, employee *employee.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockEmployeeMockRecorder) UpdateEmployee(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockEmployee)(nil).UpdateEmployee), ctx, employee)
}
