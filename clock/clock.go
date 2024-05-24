package clock

import "time"

const FormatDateTime = "2006-01-02 15:04:05" //modified from time.RFC3339

type Clock interface {
	Now() time.Time
	NowAsString() string
}

type RealClock struct {
}

func (c RealClock) Now() time.Time {
	return time.Now()
}

func (c RealClock) NowAsString() string {
	return c.Now().Format(FormatDateTime)
}

type StaticClock struct {
}

func (c StaticClock) Now() time.Time {
	staticDate, _ := time.Parse(FormatDateTime, FormatDateTime)
	return staticDate
}

func (c StaticClock) NowAsString() string {
	return c.Now().Format(FormatDateTime)
}
