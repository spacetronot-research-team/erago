package service

import "github.com/spacetronot-research-team/erago/pkg/version"

//go:generate mockgen -source=version.go -destination=mockservice/version.go -package=mockservice

type Version interface {
	// GetCurrentVersion return string current version.
	GetCurrentVersion() string
}

type versionService struct {
}

func NewVersionService() Version {
	return &versionService{}
}

// GetCurrentVersion implements Version. Return string current version.
func (*versionService) GetCurrentVersion() string {
	return version.Current
}
