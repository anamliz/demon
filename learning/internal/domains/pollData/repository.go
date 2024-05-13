package pollData

import "context"

type PollDataRepository interface {
	Save(context.Context, Sports) (int, error)
	Get(ctx context.Context) ([]Sports, error)
}
