package service

import "github.com/faizalom/go-api/internal/repository"

type ServiceA struct {
	repoA *repository.RepoA
}

func NewServiceA(repoA *repository.RepoA) *ServiceA {
	return &ServiceA{repoA: repoA}
}

func (s *ServiceA) DoWorkA() (string, error) {
	return s.repoA.GetDataA()
}
