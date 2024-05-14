package pollData

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/anamliz/learning/internal/domains/client/polldata"
	"github.com/anamliz/learning/internal/domains/pollData"
	pollDataMysql "github.com/anamliz/learning/internal/domains/pollData/pollDataMyql"
)

// PollDataServicesConfiguration is an alias for a function that will take in a pointer to an PollDataService and modify it
type PollDataServicesConfiguration func(os *PollDataService) error

// PollDataService is a implementation of the PollDataService
type PollDataService struct {
	pollMysql pollData.PollDataRepository
}

type Data struct {
	RawData string
}

// NewPollDataService : instantiate every connection we need to run season service
func NewPollDataService(cfgs ...PollDataServicesConfiguration) (*PollDataService, error) {
	// Create the seasonService
	os := &PollDataService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithMysqlPollDataRepository : instantiates mysql to connect to matches interface
func WithMysqlPollDataRepository(connectionString string) PollDataServicesConfiguration {
	return func(os *PollDataService) error {
		// Create Matches repo
		d, err := pollDataMysql.New(connectionString)
		if err != nil {
			return err
		}
		os.pollMysql = d
		return nil
	}
}

// ProcessLiveScores : processes live scores
func (s *PollDataService) PollData(ctx context.Context, pollDataEndPoint string, timeouts time.Duration, client *http.Client) error {

	// Poll Data from external API

	d, err := polldata.New(pollDataEndPoint, timeouts, client)
	if err != nil {
		return err
	}

	data, err := d.GetData(ctx)
	if err != nil {
		return fmt.Errorf("Unable to fetch data |  %v", err)
	}

	for _, d := range data {

		log.Printf("*** SportID: %s | SportName: %s ", d.SportID, d.SportName)

		// Save into database
		data, err := pollData.NewPollData(d.MatchCount, d.SBinomen, d.SportID, d.SportName, d.SportTypeID)
		if err != nil {
			log.Printf("Err : %s", err)
		} else {

			_, err = s.pollMysql.Save(ctx, *data)
			if err != nil {
				log.Printf("Err : %s", err)
			}
		}

	}

	return nil

}
