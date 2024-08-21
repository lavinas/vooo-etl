package usecase

import (
	"github.com/lavinas/vooo-etl/internal/port"
)

const (
	InLimit = int64(5000)
)

type References struct {
	Id            int64
	Name          string
	Base          string
	Object        string
	Field         string
	Last          int64
	Limit         int64
	FieldReferrer string
	FieldReferred string
}

// Base is a struct that represents the base use case
type Base struct {
	RepoSource port.Repository
	RepoTarget port.Repository
}

// NewBase creates a new use case
func NewBase(repoSource port.Repository, repoTarget port.Repository) *Base {
	return &Base{RepoSource: repoSource, RepoTarget: repoTarget}
}
