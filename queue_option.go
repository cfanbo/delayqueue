package delayqueue

import "time"

type Option func(*Options)

type Options struct {
	frequency time.Duration
	slotsCount int
}

func NewQueueOptions(opts ...Option) *Options {
	defaultOptions := &Options{
		frequency: time.Second,
		slotsCount: 3600,
	}

	for _, apply := range opts {
		apply(defaultOptions)
	}

	return defaultOptions
}

// WithFrequency 移动slot的频率，这里批一秒移动一个slot
func WithFrequency(frequency time.Duration) Option {
	return func(opts *Options) {
		opts.frequency = frequency
	}
}

//func WithSlotsCount(n int) Option {
//	return func(opts *Options) {
//		opts.slotsCount = n
//	}
//}