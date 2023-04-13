package main

import (
	"cgpt_scaleWeightsAndMeasures_Go/pkg/weightsandmeasures"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	http.HandleFunc("/scaleWeightsAndMeasures", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		// Get query parameters and convert to appropriate types
		quantity, err := strconv.ParseFloat(params.Get("quantity"), 64)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		fromUnit := params.Get("fromUnit")
		toUnit := params.Get("toUnit")

		scaleFactor, err := strconv.ParseFloat(params.Get("scaleFactor"), 64)
		if err != nil {
			http.Error(w, "Invalid scale factor", http.StatusBadRequest)
			return
		}

		// Perform the conversion using the provided values
		convertedValue, err := weightsandmeasures.ConvertQuantity(scaleFactor, quantity, fromUnit, toUnit)
		if err != nil {
			http.Error(w, "Error converting quantities", http.StatusBadRequest)
			return
		}

		// Send the response back to the client
		fmt.Fprintf(w, "Converted quantity : %s", convertedValue)
	})

	http.ListenAndServe(":"+serverPort, nil)
}
