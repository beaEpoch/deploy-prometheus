package main

import (
	"flag"
	"net/http"
	"model"
	"github.com/gin-gonic/gin"
	"notifier"
)

var (
	h            bool
	robot string
	address      string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&robot, "robot", "", "global robot webhook, you can overwrite by alert rule with annotations robot")
	flag.StringVar(&address, "address", "", "you can specify the web listen address")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if address == "" {
		address = ":8991"
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = notifier.Send(notification, robot)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to lark successful!"})

	})
	router.Run(address)
}
