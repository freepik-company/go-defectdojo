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

func (this Client) EnableDebug() {
	this.debug = true
}

func (this *Client) Get(path string) string {
	request := this.getRequest(path)
	return this.doRequest(request)
}

func (this *Client) doRequest(request *http.Request) string {

	if this.debug {
		reqDump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(reqDump))
	}

	response, err := this.httpClient.Do(request)
	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if this.debug {
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

func (this Client) getRequest(path string) *http.Request {
	return this.createRequest(path, http.MethodGet, "")
}

func (this Client) createRequest(path string, method string, body string) *http.Request {
	request, err := http.NewRequest(method, this.getUrl(path), bytes.NewBuffer([]byte(body)))

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Token %s", this.apiToken))
	request.Header.Add("accept", "application/json")

	return request
}

func (this Client) getUrl(path string) string {
	return strings.TrimRight(this.serverUrl, "/") + "/" + strings.TrimLeft(path, "/")
}

func (this *Client) Post(path string, body string) string {
	request := this.postRequest(path, body)
	return this.doRequest(request)
}

func (this *Client) postRequest(path string, body string) *http.Request {
	request := this.createRequest(path, http.MethodPost, body)
	request.Header.Add("Content-Type", "*/*; charset=UTF-8")
	return request
}

func (this *Client) Put(path string, body string) string {
	request := this.putRequest(path, body)
	return this.doRequest(request)
}

func (this *Client) putRequest(path string, body string) *http.Request {
	request := this.createRequest(path, http.MethodPut, body)
	request.Header.Add("Content-Type", "*/*; charset=UTF-8")

	return request
}

func (this Client) Delete(path string) string {
	request := this.createRequest(path, http.MethodDelete, "")
	return this.doRequest(request)
}
