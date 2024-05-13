package pollData

import "context"

type DataFetcher interface {
	GetData(ctx context.Context) ([]Sports, error)
}
