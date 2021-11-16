package blockchain

import "time"

type MedianTimeSource interface {
	AdjustedTime() time.Time

	AddTimeSample(id string, timeVal time.Time)

	Offset() time.Duration
}
