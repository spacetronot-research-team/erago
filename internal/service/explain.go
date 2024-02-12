package service

import (
	"github.com/spacetronot-research-team/erago/internal/repository"
)

//go:generate mockgen -source=explain.go -destination=mockservice/explain.go -package=mockservice

type Explain interface {
	// GetCodeArchExplanation return string code arch explanation.
	GetCodeArchExplanation() string
}

type explainService struct {
	explainRepository repository.Explain
}

func NewExplainService(explainRepository repository.Explain) Explain {
	return &explainService{
		explainRepository: explainRepository,
	}
}

// GetCodeArchExplanation implements Explain.
// GetCodeArchExplanation return string code arch explanation.
func (es *explainService) GetCodeArchExplanation() string {
	codeArchExplanation := es.explainRepository.GetCodeArchExplanationTemplate()
	return codeArchExplanation
}
