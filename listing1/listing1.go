package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/time", currentTimeHandler)
	router.Run(":8888")
}

// currentTimeHandler returns the current time in a specific time zone, UTC being the default
// Try  /time?zone=UTC    /time?zone=Asia/Shanghai   /time?zone=Europe/Amsterdam   /time?zone=MST
func currentTimeHandler(c *gin.Context) {
	zoneQuery := c.DefaultQuery("zone", "UTC")
	if zone, err := time.LoadLocation(zoneQuery); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("zone %v not found", zoneQuery))
	} else {
		c.JSON(http.StatusOK, gin.H{"time": formatTimeInZone(time.Now(), zone)})
	}
}

func formatTimeInZone(t time.Time, zone *time.Location) string {
	return t.In(zone).Format("2006-01-02T15:04:05.99MST")
}
