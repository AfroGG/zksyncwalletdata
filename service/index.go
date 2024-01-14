package service

import (
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func GetTable(c *gin.Context) {
	c.HTML(200, "table.html", gin.H{})
}

func GetAnalyze(c *gin.Context) {
	c.HTML(200, "analyze.html", gin.H{})
}

func GetCoffee(c *gin.Context) {
	c.HTML(200, "coffee.html", gin.H{})
}

func GetRule(c *gin.Context) {
	c.HTML(200, "rule.html", gin.H{})
}
