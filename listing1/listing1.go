package main

import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
    "time"
    "fmt"
)

//Try  /time?zone=UTC    /time?zone=Asia/Shanghai   /time?zone=Europe/Amsterdam   /time?zone=MST

func main() {
    router := gin.Default()
    router.GET("/status", StatusHandler)
    router.GET("/time", TimeHandler)
    router.Run(":8888")
}

// StatusHandler returns the status of the Time Service
func StatusHandler(c *gin.Context) {
    c.String(http.StatusOK, "TimeServer Status = OK")
}

// TimeHandler returns the current time in a specific time zone, UTC being the default
func TimeHandler(c *gin.Context) {
    requestedTimeZone := c.DefaultQuery("zone", "UTC")
    if timeZone, err := time.LoadLocation(requestedTimeZone); err == nil {
        c.JSON(http.StatusOK, gin.H{"time": time.Now().In(timeZone).Format("2006-01-02T15:04:05.99MST")})
    } else {
        c.String(http.StatusBadRequest, fmt.Sprintf("time zone %v not found", requestedTimeZone))
    }
}