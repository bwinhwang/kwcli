package common

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type KWClient struct {
	baseURL    string
	httpClient *http.Client
	user       string
	token      string
}

var validTokenPattern = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// Constructor
func NewKWClient(baseURL, user, token string) *KWClient {
	return &KWClient{
		baseURL:    baseURL,
		httpClient: &http.Client{}, // Create a default http client
		user:       user,
		token:      token,
	}
}

// Example of a common function
/*
func (client *KWClient) makeAPIRequest(method, endpoint string, data url.Values) (*http.Response, error) {
	// 1. Construct the full URL
	// 2. Create 'http.Request' with method, headers (add apiKey if necessary) etc.
	// 3. Use client.httpClient.Do to execute the request
	// 4. Handle and return response (or error if any)
}
*/

func (client *KWClient) Echo(data map[string]interface{}) ([]string, error) {

	if client.user != "" {
		data["user"] = client.user
	}
	if client.token != "" && validTokenPattern.MatchString(client.token) {
		data["ltoken"] = client.token
	}

	s := constructFields(data)
	//fmt.Println(s)
	req, err := http.NewRequest(http.MethodPost, client.baseURL, bytes.NewReader([]byte(s)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	//fmt.Println(string(body))

	return strings.Split(string(body), "\n"), nil
}

func constructFields(data map[string]interface{}) string {

	fields := url.Values{}
	for name, value := range data {
		//fmt.Printf("Parameter: %s, Value: %s\n", name, value)
		fields.Add(name, value.(string))
	}

	return url.Values(fields).Encode()
}
