package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Stats struct {
	Status	map[string]int	`json:"status"`
}

var statWater, statWind, colorWater, colorWind string
var randWater, randWind int

func main() {
    engine := html.New("views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })
	app.Get("/", func(c *fiber.Ctx) error {
       return serve(c)
   	})
	app.Static("/", "./public")

    log.Fatal(app.Listen(":3000"))
}

func serve(c *fiber.Ctx) error {
	fileContent, err := os.Open("json/data.json")

	if err != nil {
		panic(err)
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteResult), &result)

	water := result["status"].(map[string]interface{})["water"].(float64)
	wind := result["status"].(map[string]interface{})["wind"].(float64)

	randWater := rand.Intn(100 - int(water)) + 1
	randWind := rand.Intn(100 - int(wind)) + 1

	if randWater <= 5 {
		statWater = "Aman"
		colorWater = "#41BEAE"
	} else if randWater >= 6 && randWater <= 8 {
		statWater = "Siaga"
		colorWater = "#FFE45E"
	} else if randWater >= 8 {
		statWater = "Bahaya"
		colorWater = "#F2555E"
	}
	
	if randWind <= 6 {
		statWind = "Aman"
		colorWind = "#41BEAE"
	} else if randWind >= 7 && randWind <= 15 {
		statWind = "Siaga"
		colorWind = "#FFE45E"
	} else if randWind >= 15 {
		statWind = "Bahaya"
		colorWind = "#F2555E"
	}

	f := Stats{
		Status: map[string]int{
            "water": randWater,
            "wind": randWind,
        },
	}
	j, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("json/data.json", j, 0644)

	return c.Render("index", fiber.Map{
		"Title": "Assignment 3 - Hendri Muda",
		"Water": randWater,
		"Wind": randWind,
		"statWater": statWater,
		"statWind": statWind,
		"colorWater": colorWater,
		"colorWind": colorWind,
    })
}
