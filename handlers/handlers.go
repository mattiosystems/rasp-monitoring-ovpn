package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"monitoring/database"
	"monitoring/resources"
	"net/http"
)

func GET_Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"sys_hostname":       resources.GetHostname(),
		"sys_os":             resources.GetOS(),
		"sys_kernel_version": resources.GetKernel(),
		"sys_cpu_vendor_id":  resources.GetCPUVendorID(),
		"sys_cpu_model":      resources.GetCPUModel(),
		"sys_cpu_mhz":        resources.GetCPUMhz(),
	})
}

func GET_VPN(c *gin.Context, db *gorm.DB) {
	conn := database.GetAllConnections(db)
	c.HTML(http.StatusOK, "vpn.html", gin.H{
		"vpn": conn,
	})
}
