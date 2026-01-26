package application

import (
	testDomain "main_module/internal/test/domain"
)

type TestService struct {
	Repo testDomain.Repository
}

func NewTestService(repo testDomain.Repository) *TestService {
	return &TestService{
		Repo: repo,
	}
}