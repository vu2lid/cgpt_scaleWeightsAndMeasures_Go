// Reads OpenAI API configration from .env creates the prompt based on input
// sends the request to OpenAI API endpoint and returns the result.
package weightsandmeasures

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// ConvertQuantity takes in a scaleFactor, quantity, fromUnit, and toUnit,
// and returns the converted quantity and an error (if any).
func ConvertQuantity(scaleFactor float64, quantity float64, fromUnit string, toUnit string) (string, error) {
	// Get OpenAI API call configuration parameters
	apiKey, model, temperature, max_tokens, n, shouldReturn, returnValue, returnValue1 := readConfiguration()
	if shouldReturn {
		return returnValue, returnValue1
	}

	// Create input prompt for OpenAI
	prompt := generatePrompt(generateQuestion(scaleFactor, quantity, fromUnit, toUnit))

	data := map[string]interface{}{
		"model":       model,
		"prompt":      prompt,
		"temperature": temperature,
		"max_tokens":  max_tokens,
		"n":           n,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error Marshalling JSON data:", err)
		return "", err
	}

	fmt.Println("Request sent was: ", string(jsonData))

	body := strings.NewReader(string(jsonData))

	// Set up the HTTP request with the API key in the headers
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", body)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the HTTP request and read the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response:", err)
		return "", err
	}

	// Pretty print the JSON response
	var jsonResponse interface{}
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON response:", err)
		return "", err
	}
	prettyJson, err := json.MarshalIndent(jsonResponse, "", "    ")
	if err != nil {
		fmt.Println("Error pretty printing JSON response:", err)
		return "", err
	}

	// Print the generated text
	fmt.Println(string(prettyJson))
	return string(prettyJson), err
}

func readConfiguration() (string, string, float64, int, int, bool, string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	model := os.Getenv("MODEL")
	temperature, err := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 64)
	if err != nil {
		fmt.Println("Error reading OpenAI configuration data temperature:", err)
		return "", "", 0, 0, 0, true, "", err
	}
	max_tokens, err := strconv.Atoi(os.Getenv("MAX_TOKENS"))
	if err != nil {
		fmt.Println("Error reading OpenAI configuration data max_tokens:", err)
		return "", "", 0, 0, 0, true, "", err
	}
	n, err := strconv.Atoi(os.Getenv("N"))
	if err != nil {
		fmt.Println("Error reading OpenAI configuration data n:", err)
		return "", "", 0, 0, 0, true, "", err
	}
	return apiKey, model, temperature, max_tokens, n, false, "", nil
}

// The CSV data table given below used for training comes from:
// https://en.wikipedia.org/wiki/Imperial_units
func generatePrompt(question string) string {
	return `Read the CSV data given below and answer the question: ` + question + `.

Unit,Imperial ounces,Imperial pints,Millilitres,Cubic inches,US ounces,US pints
fluid ounce (fl oz),1 ,1/20 ,28.4130625 ,1.7339 ,0.96076 ,0.060047
gill (gi),5 ,1/4 ,142.0653125 ,8.6694 ,4.8038 ,0.30024
pint (pt),20 ,1 ,568.26125 ,34.677 ,19.215 ,1.2009
quart (qt),40 ,2 ,1136.5225 ,69.355 ,38.43 ,2.4019
gallon (gal),160 ,8 ,4546.09 ,277.42 ,153.72 ,9.6076

`
}

func generateQuestion(scaleFactor float64, quantity float64, fromUnit string, toUnit string) string {
	return fmt.Sprintf("Multiply quantity %.2f with a scale factor of %.2f and convert from %s to %s", quantity, scaleFactor, fromUnit, toUnit)
}
