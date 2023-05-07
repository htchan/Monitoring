package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const TEMPERATURE_FILENAME = "/sys/class/thermal/thermal_zone0/temp"
const TEMPERATURE_UPDATE_INTERVAL = time.Second * 10
const SERVER_PORT = ":8000"

func readTemperature() (float64, error) {
	data, fileErr := os.ReadFile(TEMPERATURE_FILENAME)
	if fileErr != nil {
		return 0, fileErr
	}

	result, convErr := strconv.ParseFloat(strings.TrimSpace(string(data)), 64)
	if convErr != nil {
		return 9, convErr
	}

	return result, nil
}

func runServer(temperature *float64) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "# HELP device_temperature Temperature measured by lm sensors\n")
		fmt.Fprintf(w, "# TYPE device_temperature gauge\n")
		fmt.Fprintf(w, "device_temperature %f\n", *temperature)
	})

	log.Printf("Start server at port %s", SERVER_PORT)
	if err := http.ListenAndServe(SERVER_PORT, nil); err != nil {
		return err
	}

	return nil
}

func main() {
	temperature := 0.0

	go func() {
		for true {
			temp, readErr := readTemperature()
			if readErr != nil {
				log.Printf("read temperature error: %v", readErr)
			} else {
				temperature = temp
			}

			time.Sleep(TEMPERATURE_UPDATE_INTERVAL)
		}
	}()

	runServer(&temperature)
}

/* sample response
'''
# HELP device_temperature Temperature measured by lm sensors
# TYPE device_temperature gauge
device_temperature 5.0
'''
*/
