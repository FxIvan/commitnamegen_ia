
	package commit

	import (
		"net/http"

		"github.com/gin-gonic/gin"
		"github.com/fxivan/commitnamegen_ia/internal/domain"
	)

	func (h Handler) CreateCommit(c *gin.Context) {
		var commit domain.Commit
		if err := c.BindJSON(&commit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertedId, err := h.CommitService.Create(commit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"commit_id": insertedId})
	}
