package helper

import (
	"log"
	"testing"
)

func TestNewWithSecond(t *testing.T) {
	c := TCorn.NewWithSecond()
	spec := "*/5 * * * * ?"
	i := 0
	_, err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	if err != nil {
		t.Errorf("cron errors: %v\n", err)
	}
	c.Start()
	c.Stop()
}