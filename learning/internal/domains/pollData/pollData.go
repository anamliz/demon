package pollData

import (
	"fmt"
	"time"
)

func NewPollData(matchCount, sBinomen, sportID, sportName, sportTypeID string) (*Sports, error) {

	if matchCount == "" {
		return &Sports{}, fmt.Errorf("Match Count not set")
	}

	if sBinomen == "" {
		return &Sports{}, fmt.Errorf("sBinomen not set")
	}

	if sportID == "" {
		return &Sports{}, fmt.Errorf("sportID not set")
	}

	if sportName == "" {
		return &Sports{}, fmt.Errorf("sportName not set")
	}

	if sportTypeID == "" {
		return &Sports{}, fmt.Errorf("sportTypeID not set")
	}

	created := time.Now().Format("2006-01-02 15:04:05")

	// Final Object
	return &Sports{
		MatchCount:  matchCount,
		SBinomen:    sBinomen,
		SportID:     sportID,
		SportName:   sportName,
		SportTypeID: sportTypeID,
		Created:     created,
		Modified:    created,
	}, nil

}
