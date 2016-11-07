package blackjack

import "context"

type Table interface {
	Open(ctx context.Context)
	Rules() *Rules
}
