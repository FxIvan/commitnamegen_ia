package commit

import (
	"time"

	"github.com/fxivan/commitnamegen_ia/internal/domain"
	vertexai "github.com/fxivan/commitnamegen_ia/internal/repositories/vertexia"
)

func (s Service) Create(commit domain.Commit) (id interface{}, err error) {
	commit.MetaInfo = time.Now().UTC()
	vertexai.MakeRequests()
	insertResult, err := s.Repo.Insert(commit)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}
