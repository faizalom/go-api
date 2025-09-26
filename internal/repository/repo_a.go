package repository

import "database/sql"

type RepoA struct {
	DB *sql.DB
}

func NewRepoA(db *sql.DB) *RepoA {
	return &RepoA{DB: db}
}

func (r *RepoA) GetDataA() (string, error) {
	return "Data from Repo A", nil
}
