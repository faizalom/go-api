package service

import "github.com/faizalom/go-api/internal/repository"

type ServiceB struct {
	repoB *repository.RepoB
}

func NewServiceB(repoB *repository.RepoB) *ServiceB {
	return &ServiceB{repoB: repoB}
}

func (s *ServiceB) DoWorkB() (string, error) {
	return s.repoB.GetDataB()
}
