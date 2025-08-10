package std

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type testClient struct {
	url string
}

func (c *testClient) MakeGetRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, c.url, nil)
}

func Get[Data any](c *testClient) (Data, error) {
	var (
		request = Async(c.MakeGetRequest)
		resp    = Chain(request, FetchHTTP)
		respF   = Finally(resp, WriteLog)
		bytes   = Chain(respF, FetchBody)
		parsed  = Chain(bytes, ParseJSON[Data])
		parsedC = Catch(parsed, LogError[Data])
	)

	return Await(parsedC)
}

func FetchHTTP(req *http.Request) (*http.Response, error) {
	// can set up headers, auth, etc.
	return http.DefaultClient.Do(req)
}

func FetchBody(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func ParseJSON[Data any](bytes []byte) (Data, error) {
	var d Data

	err := json.Unmarshal(bytes, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func LogError[Ok any](err error) (Ok, error) {
	var empty Ok
	if err != nil {
		slog.Warn("cannot do GET request: %w", err)
		return empty, err
	}
	return empty, nil
}

func WriteLog() error {
	logStr := fmt.Sprintf("GET request ended at %s", time.Now())
	return os.WriteFile("getter.log", []byte(logStr), os.ModeAppend)
}
