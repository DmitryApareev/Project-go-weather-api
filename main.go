package main

import (
	"flag"
	"fmt"
	"study/weather/geo"
	"study/weather/weather"
)

func main() {
	fmt.Println("Новый проект")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	weatherData, _ := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
