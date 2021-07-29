package shared

import "time"

//go:generate moq -out timeProvider_mock.go . TimeProvider

type TimeProvider interface {
	Now() time.Time
}
