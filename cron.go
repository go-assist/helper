package helper

import "github.com/robfig/cron/v3"

// NewWithSecond win linux兼容问题.
func (tc *TsCorn) NewWithSecond() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}
