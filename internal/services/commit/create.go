package commit

import (
	"time"

	"github.com/fxivan/commitnamegen_ia/internal/domain"
)

func (s Service) Create(commit domain.Commit) (id interface{}, err error) {
	commit.MetaInfo = time.Now().UTC()

	insertResult, err := s.Repo.Insert(commit)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}
