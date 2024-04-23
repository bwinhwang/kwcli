package common

import (
	"bytes"
	"encoding/json"
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

type KWResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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

func (r *KWResponse) Validate() error {
	if r.Status == 0 {
		return fmt.Errorf("missing 'status' field in response")
	}
	if r.Message == "" {
		return fmt.Errorf("missing 'message' field in response")
	}
	// ... Add more checks if needed ...
	return nil
}

func (client *KWClient) Execute(data map[string]interface{}) ([]string, error) {

	var kwresp KWResponse
	if client.user != "" {
		data["user"] = client.user
	}
	if client.token != "" && validTokenPattern.MatchString(client.token) {
		data["ltoken"] = client.token
	}

	//fmt.Println(data)
	s, ok := constructFields(data)
	if !ok {
		return nil, fmt.Errorf("constructFields fail due to unsupport types")
	}
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
	if string(body) == "" {
		//action succeeds without response
		return nil, nil
	}

	lines := strings.Split(string(body), "\n")
	first := lines[0]
	//fmt.Println(first)

	_ = json.Unmarshal([]byte(first), &kwresp)

	/*
		if err != nil {
			return nil, err
		}
	*/

	err = kwresp.Validate()

	if err != nil {
		return lines, nil
	}

	return nil, fmt.Errorf("status: %d, message: %s", kwresp.Status, kwresp.Message)
}

func constructFields(data map[string]interface{}) (string, bool) {

	fields := url.Values{}
	for name, value := range data {
		fields.Add(name, value.(string))
	}

	return url.Values(fields).Encode(), true
}
