package delayqueue

import "time"

type Option func(*Options)

// Options 队列选项
type Options struct {
	frequency time.Duration
	slotsNum  int
}

// Frequency 获取定时间隔
func (o *Options) Frequency() time.Duration {
	return o.frequency
}

// NewQueueOptions 创建队列选项
func NewQueueOptions(opts ...Option) *Options {
	defaultOptions := &Options{
		frequency: time.Second,
		slotsNum:  SlotsNum,
	}

	for _, apply := range opts {
		apply(defaultOptions)
	}

	return defaultOptions
}

// WithFrequency 移动slot的频率，默认情况下每秒移动一个slot
func WithFrequency(frequency time.Duration) Option {
	return func(opts *Options) {
		opts.frequency = frequency
	}
}
