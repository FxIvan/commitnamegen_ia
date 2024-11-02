
	package ports

	import "github.com/fxivan/commitnamegen_ia/internal/domain"

	type CommitService interface {
		Create(commit domain.Commit) (id interface{}, err error)
	}

	type CommitRepository interface {
		Insert(commit domain.Commit) (id interface{}, err error)
	}
