package main

import (
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	a := api.NewAPI(gbot)
	a.Port = "8080"
	a.AddHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q \n", html.EscapeString(r.URL.Path))
	})
	a.Debug()
	a.Start()

	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/ttyACM0")
	led := gpio.NewLedDriver(firmataAdaptor, "led", "2")

	work := func() {
		//gobot.On(func() {
		led.On()
		//})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)
	robot.AddCommand("Offled", func(params map[string]interface{}) interface{} {
		led.Off()
		return fmt.Sprintf("says hello!")
	})
	robot.AddCommand("Onled", func(params map[string]interface{}) interface{} {
		led.On()
		return fmt.Sprintf("says hello!")
	})
	robot.AddCommand("Toggleled", func(params map[string]interface{}) interface{} {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
		return fmt.Sprintf("says hello!")
	})
	gbot.AddRobot(robot)
	gbot.Start()
}
