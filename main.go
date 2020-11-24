package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	url := "https://api.github.com/users/"

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	app.Get("/:user", func(c *fiber.Ctx) error {

		resp, err := http.Get(url + c.Params("user"))
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		return c.Send(body)
	})

	app.Listen(":4321")
}
