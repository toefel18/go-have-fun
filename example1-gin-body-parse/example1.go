package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/tojson", toJsonHandler)
	router.Run(":8888")
}

type message struct {
	Text   string
	Author string
}

//toJson returns the message as JSON
func (msg message) toJson() string {
	jsonMsg, _ := json.Marshal(msg)
	return string(jsonMsg)
}

type Jsonizable interface {
	toJson() string
}

// toJsonHandler parses the request body and responds with the same message in JSON
// curl -X POST http://localhost:8888/tojson -H "Content-Type: application/xml"  -d "<Message><Text>Hi Gopher</Text><Author>Christope</Author></Message>"
// curl -X POST http://localhost:8888/tojson -H "Content-Type: application/json" -d '{"Text":"Hi Gopher", "Author":"Christophe"}'
// curl -X POST http://localhost:8888/tojson -d "Text=Hi Gopher&Author=Christophe"
func toJsonHandler(c *gin.Context) {
	var msg message
	if err := c.Bind(&msg); err != nil {
		c.String(http.StatusBadRequest, "Error parsing body: "+err.Error())
	} else {
		c.Status(http.StatusOK)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteString(msg.toJson())
	}
}
