package diccionario

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AddRequest represents the request body for the /add endpoint.
type AddRequest struct {
	// Word is the word to add to the word list.
	Word string `json:"word"`
}

// Add a new word to the word list.
func (s *Server) Add(c *gin.Context) {
	var req AddRequest
	if err := c.BindJSON(&req); err != nil {
		// should return in JSON with appropriate err info
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// implement your logic here

	word := strings.TrimSpace(req.Word)

	// validate the input first
	if !isValid(word) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Word can only contain alphabetic characters"})
		return
	}

	// check if it already exists
	wordlist, err := s.w.GetWords() 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to read wordList"})
	}

	target := strings.ToLower(word)
	for _, w := range wordlist {
		if strings.ToLower(strings.TrimSpace(w)) == target {
			c.Status(http.StatusConflict)
			return
		}
	}

	if err := s.w.AddWord(word); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to add new word"})
		return
	}

	c.Status(http.StatusNoContent)

}

func isValid(word string) bool {
	if word == "" {
		return false
	}

	for _, r := range word {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}
