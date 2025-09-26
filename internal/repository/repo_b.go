package repository

import "database/sql"

type RepoB struct {
	DB *sql.DB
}

func NewRepoB(db *sql.DB) *RepoB {
	return &RepoB{DB: db}
}

func (r *RepoB) GetDataB() (string, error) {
	return "Data from Repo B", nil
}
