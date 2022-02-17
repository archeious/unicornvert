package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type charSets struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LowerOffset int32  `json:"lower-offset"`
	UpperOffset int32  `json:"upper-offset"`
}

type message struct {
	Message string `json:"message"`
}

// Character Sets.
var CharSets = []charSets{
	{ID: "1", Name: "Bubble Text", LowerOffset: 9327, UpperOffset: 9333},
}

func main() {
	router := gin.Default()
	router.POST("/convert", convertText)

	router.Run("0.0.0.0:8689")
}

// postAlbums adds an album from JSON received in the request body.
func convertText(c *gin.Context) {
	var original_message message
	var new_message message

	if err := c.BindJSON(&original_message); err != nil {
		return
	}

	for _, c := range original_message.Message {
		fmt.Println(string(c), ":", int(c))
		if c >= 'a' && c <= 'z' {
			new_message.Message += string(int(c) + int(CharSets[0].LowerOffset))
		} else if c >= 'A' && c <= 'Z' {
			new_message.Message += string(int(c) + int(CharSets[0].UpperOffset))
		} else {
			new_message.Message += string(c)
		}
	}
	c.IndentedJSON(http.StatusOK, new_message.Message)
}
