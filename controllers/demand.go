package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDemandPulse(c *gin.Context) {
	resp, err := http.Get("http://localhost:5001/api/demand")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "Demand Pulse Engine is not available",
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read demand data",
		})
		return
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	c.JSON(http.StatusOK, result)
}

func GetAssessment(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	jsonData, _ := json.Marshal(body)
	resp, err := http.Post(
		"http://localhost:5001/api/assess",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "Assessment service is not available",
		})
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var result map[string]interface{}
	json.Unmarshal(responseBody, &result)
	c.JSON(http.StatusOK, result)
}
