package Client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Client struct {
	serverUrl  string
	apiToken   string
	httpClient *http.Client
	debug      bool
}

func NewClient(serverUrl string, apiToken string) Client {
	return Client{
		serverUrl:  serverUrl,
		apiToken:   apiToken,
		httpClient: &http.Client{},
		debug:      false,
	}
}

func (c Client) EnableDebug() {
	c.debug = true
}

func (c *Client) Get(path string) string {
	request := c.getRequest(path)
	return c.doRequest(request)
}

func (c *Client) doRequest(request *http.Request) string {

	if c.debug {
		reqDump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(reqDump))
	}

	response, err := c.httpClient.Do(request)
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if c.debug {
		respDump, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(respDump))
	}

	body, err := io.ReadAll(response.Body)

	if response.StatusCode > 299 {
		log.Fatal(fmt.Sprintf("Error: Response code %d is not a valid status. Response says: %s", response.StatusCode, body))
	}

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func (c Client) getRequest(path string) *http.Request {
	return c.createRequest(path, http.MethodGet, "")
}

func (c Client) createRequest(path string, method string, body string) *http.Request {
	request, err := http.NewRequest(method, c.getUrl(path), bytes.NewBuffer([]byte(body)))

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Token %s", c.apiToken))
	request.Header.Add("accept", "application/json")

	return request
}

func (c Client) getUrl(path string) string {
	return strings.TrimRight(c.serverUrl, "/") + "/" + strings.TrimLeft(path, "/")
}

func (c *Client) Post(path string, body string) string {
	request := c.postRequest(path, body)
	return c.doRequest(request)
}

func (c *Client) postRequest(path string, body string) *http.Request {
	request := c.createRequest(path, http.MethodPost, body)
	request.Header.Add("Content-Type", "*/*; charset=UTF-8")
	return request
}

func (c *Client) Put(path string, body string) string {
	request := c.putRequest(path, body)
	return c.doRequest(request)
}

func (c *Client) putRequest(path string, body string) *http.Request {
	request := c.createRequest(path, http.MethodPut, body)
	request.Header.Add("Content-Type", "*/*; charset=UTF-8")

	return request
}
