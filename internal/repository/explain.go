package repository

import "github.com/spacetronot-research-team/erago/internal/repository/templ"

//go:generate mockgen -source=explain.go -destination=mockrepository/explain.go -package=mockrepository

type Explain interface {
	// GetCodeArchExplanationTemplate return string code arch explanation.
	GetCodeArchExplanationTemplate() string
}

type explainRepository struct {
}

func NewExplainRepository() Explain {
	return &explainRepository{}
}

// GetCodeArchExplanationTemplate implements Explain.
// GetCodeArchExplanationTemplate return string code arch explanation.
func (*explainRepository) GetCodeArchExplanationTemplate() string {
	return templ.Explain
}
