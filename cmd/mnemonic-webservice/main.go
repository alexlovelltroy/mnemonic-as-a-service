package main

import (
	"strconv"

	"github.com/alexlovelltroy/mnemonic"
	"github.com/gin-gonic/gin"
)

func getRandomWord(c *gin.Context) {
	wordcount := c.Params.ByName("count")
	intWordCount, err := strconv.Atoi(wordcount)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid word count.  Must be an integer."})
	}
	c.JSON(200, gin.H{
		"words": mnemonic.GetRandomWords(intWordCount),
	})
}

func main() {
	r := gin.Default()
	r.GET("/random/:count", getRandomWord)
	r.Run(":8080")
}
