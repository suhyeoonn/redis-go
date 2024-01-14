package main

import (
	"fmt"
	"net/http"

	"redis-go/src/fetch"
	"redis-go/src/photos"
	"redis-go/src/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	redis.ConnectRedis()

	r := gin.Default()

	r.GET("/photos", func(c *gin.Context) {
		// redis에 저장된 값이 있으면 리턴
		if cache := redis.GetCache("photos"); cache != "" {
			fmt.Println("Cache Hit")

			c.JSON(http.StatusOK, gin.H{
				"photos": photos.JSONParse([]byte(cache)),
			})
			return
		}

		fmt.Println("Cache Miss")

		data := fetch.FetchData("https://jsonplaceholder.typicode.com/photos")

		if err := redis.SetCache("photos", string(data)); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"photos": photos.JSONParse(data),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
