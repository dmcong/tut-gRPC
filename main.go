package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()


	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.GET("/sum", func(c *gin.Context) {
		startTime := time.Now()

		numStr := c.DefaultQuery("number", "0")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.String(400, "Invalid number")
			return
		}
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		c.JSON(200, gin.H{"time":  time.Since(startTime), "data": strconv.Itoa(sum)})
	})

	router.GET("/sumv2", func(c *gin.Context) {
		startTime := time.Now()

		numStr := c.DefaultQuery("number", "0")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid number"})
			return
		}

		sum := parallelSum(1,num)

		c.JSON(200, gin.H{"time":  time.Since(startTime), "data": strconv.Itoa(sum)})
	})

	router.POST("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.PUT("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.DELETE("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.Run(":8080")
}



func parallelSum(start, end int) int {
	delta := (end - start + 1) / 20
	if delta == 0 {
		sum := 0
		for i := start; i <= end; i++ {
			sum += i
		}
		return sum
	}

	resultChan := make(chan int)
	for i := start; i <= end; i += delta {
		s := i
		e := i + delta - 1
		if e > end {
			e = end
		}

		go func(s, e int) {
			sum := 0
			for i := s; i <= e; i++ {
				sum += i
			}
			resultChan <- sum
		}(s, e)
	}

	sum := 0
	for i := start; i <= end; i += delta {
		sum += <-resultChan
	}

	return sum
}