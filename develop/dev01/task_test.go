package main

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/beevik/ntp"
)

/*
	Взято из официальной документации
	https://github.com/beevik/ntp/blob/main/ntp_test.go
*/

var ErrKissOfDeath            = errors.New("kiss of death received")

func isNil(t *testing.T, host string, err error) bool {
	switch {
	case err == nil:
		return true
	case err == ErrKissOfDeath:
		// log instead of error, so test isn't failed
		t.Logf("[%s] Query kiss of death (ignored)", host)
		return false
	case strings.Contains(err.Error(), "timeout"):
		// log instead of error, so test isn't failed
		t.Logf("[%s] Query timeout (ignored): %s", host, err)
		return false
	default:
		// error, so test fails
		t.Errorf("[%s] Query failed: %s", host, err)
		return false
	}
}
var timeFormat = "Mon Jan _2 2006  15:04:05.00000000 (MST)"
func TestOnlineTime(t *testing.T) {
	tm, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	now := time.Now()
	if isNil(t, "0.beevik-ntp.pool.ntp.org", err) {
		t.Logf(" System Time: %s\n", now.Format(timeFormat))
		t.Logf("  ~True Time: %s\n", tm.Format(timeFormat))
		t.Logf("~ClockOffset: %v\n", tm.Sub(now))
	}
}