package useCase

import (
	"companies/internal/domain/company"
	"companies/internal/domain/employee"
	"companies/internal/domain/project"
	"context"

	"github.com/google/uuid"
)

type Storage interface {
	Company
	Project
	Employee
}

type Company interface {
	CreateCompany(ctx context.Context, company *company.Company) (err error)
	UpdateCompany(ctx context.Context, company *company.Company) (err error)
	DeleteCompany(ctx context.Context, ID uuid.UUID) (err error)
	ReadCompanyByID(ctx context.Context, ID uuid.UUID) (company *company.Company, err error)
}

type Project interface {
	CreateProject(ctx context.Context, project *project.Project) (err error)
	UpdateProject(ctx context.Context, project *project.Project) (err error)
	DeleteProject(ctx context.Context, ID uuid.UUID) (err error)
	ReadProjectByID(ctx context.Context, ID uuid.UUID) (err error)
	ReadProjectsOfCompany(ctx context.Context, companyID uuid.UUID) ([]*project.Project, error)
}

type Employee interface {
	CreateEmployee(ctx context.Context, employee *employee.Employee) (err error)
	UpdateEmployee(ctx context.Context, employee *employee.Employee) (err error)
	DeleteEmployee(ctx context.Context, ID uuid.UUID) (err error)
	ReadEmployeeByID(ctx context.Context, ID uuid.UUID) (err error)
	ReadAllEmployeesOfProject(ctx context.Context, projectID uuid.UUID) (err error)
	ReadFiltredEmployeesOfProject(ctx context.Context, projectID uuid.UUID) (err error)
}
