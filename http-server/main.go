package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var addr string
	flag.StringVar(&addr, "listen", "0.0.0.0:8000", "listen addr")
	flag.Parse()

	run(addr)
}

func run(addr string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/users", getUsers)
	router.Run(addr)
}

func getUsers(c *gin.Context) {
	users := &[]User{
		User{
			Name:      "zhangsan",
			Age:       18,
			Telephone: "13011111111",
			Email:     "13011111111@qq.com",
		},
		User{
			Name:      "lisi",
			Age:       28,
			Telephone: "13111111111",
			Email:     "13111111111@qq.com",
		},
	}
	c.JSON(http.StatusOK, users)
}

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
}
