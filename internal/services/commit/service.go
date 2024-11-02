
package commit

import (
	"github.com/fxivan/commitnamegen_ia/internal/ports"
)

type Service struct {
	Repo ports.CommitRepository
}

