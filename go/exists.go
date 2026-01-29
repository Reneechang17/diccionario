package diccionario

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ExistsResponse is the response sent back for the exists endpoint.
type ExistsResponse struct {
	// Exists is true if the word exists; otherwise, false.
	Exists bool `json:"exists"`
}

// WordExists returns true if the word exists in the word list.
// It performs case insensitive matching to the words in the wordlist.
func (s *Server) WordExists(c *gin.Context) {
	word := c.Param("word")

	wordlist, err := s.w.GetWords()
	if err != nil {
		// read file err should be 5xx response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read wordlist"})
		return
	}

	// resp := ExistsResponse{Exists: false}

	// for _, w := range wordlist {
	// 	// 
	// 	if strings.HasPrefix(w, word) {
	// 		resp.Exists = true
	// 	}
	// }

	target := strings.ToLower(word)
	exists := false

	for _, w := range wordlist {
		if strings.ToLower(strings.TrimSpace(w)) == target {
			exists = true
			break
		}
	}

	c.JSON(http.StatusOK, ExistsResponse {Exists: exists})
}
