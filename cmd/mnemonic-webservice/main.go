package main

import (
	"strconv"

	"github.com/alexlovelltroy/mnemonic"
	"github.com/docker/docker/pkg/namesgenerator"
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

func getRandomName(c *gin.Context) {

	c.JSON(200, gin.H{
		"name": namesgenerator.GetRandomName(0),
	})
}

func main() {
	r := gin.Default()
	r.GET("/random/:count", getRandomWord)
	r.GET("/randomname", getRandomName)
	r.Run(":8080")
}
