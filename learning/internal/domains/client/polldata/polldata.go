package polldata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/lukemakhanu/learning/internal/domains/pollData"
)

// Compile time interface assertion.
var _ pollData.DataFetcher = (*PollDataClient)(nil)

type PollDataClient struct {
	pollDataEndPoint string
	timeouts         time.Duration
	client           *http.Client
}

// New initializes a new instance of Live Score Client.
func New(pollDataEndPoint string, timeouts time.Duration, client *http.Client) (*PollDataClient, error) {

	pollDataURL, err := url.Parse(pollDataEndPoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse poll data endpoint: %w", err)
	}

	if timeouts <= 0 {
		return nil, fmt.Errorf("Timeout not set")
	}

	c := &PollDataClient{
		pollDataEndPoint: pollDataURL.String(),

		timeouts: timeouts,
		client:   client,
	}
	if c.client == nil {
		c.client = defaultHTTPClient
	}

	return c, nil
}

func (s *PollDataClient) GetData(ctx context.Context) ([]pollData.Sports, error) {

	var pollDataUrl = fmt.Sprintf("%s", s.pollDataEndPoint)
	log.Printf("Calling... : %s", pollDataUrl)

	response, err := s.client.Get(pollDataUrl)
	if err != nil {
		return nil, fmt.Errorf("Failed to call poll data API")
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w | responseBody : %s ", err, string(responseBody))
	}

	if response.StatusCode == http.StatusOK {
		return parseApiData(responseBody)
	}

	return nil, fmt.Errorf("failed to get data : status: %d, error body: %s", response.StatusCode, responseBody)

}

func parseApiData(content []byte) ([]pollData.Sports, error) {

	var g []pollData.Sports

	var s pollData.AllData
	err := json.Unmarshal(content, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json : %w", err)
	}

	log.Printf("statusCode %s | statusDescription %d", s.StatusDescription, s.StatusCode)

	for _, i := range s.Data.Sports {

		log.Printf("SportID %s | SportName %s", i.SportID, i.SportName)

		d := pollData.Sports{
			MatchCount:  i.SportID,
			SBinomen:    i.SBinomen,
			SportID:     i.SportID,
			SportName:   i.SportName,
			SportTypeID: i.SportTypeID,
		}

		g = append(g, d)

	}

	if len(g) == 0 {
		return g, fmt.Errorf("sport is empty ")
	}

	return g, nil
}

var defaultHTTPClient = &http.Client{
	Timeout: time.Second * 15,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Second * 15,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	},
}
