package api

import (
	"github.com/gin-gonic/gin"
	"monitoring/resources"
	"net/http"
)

func GetResources(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"cpu_usage_percentage":    resources.GetCPUUsedPercentage(),
		"ram_usage":               resources.GetRAMUsed(),
		"ram_usage_percent":       resources.GetRAMUsedPercentage(),
		"ram_total":               resources.GetRAMTotal(),
		"storage_usage":           resources.GetStorageUsed(),
		"storage_usage_percent":   resources.GetStorageUsedPercentage(),
		"storage_total":           resources.GetStorageTotal(),
		"cpu_temperature":         resources.GetCPUTemperature(),
		"cpu_temperature_percent": resources.GetCPUTemperaturePercentage(),
	})
}
