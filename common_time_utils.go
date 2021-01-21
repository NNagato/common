package common

import "time"

// TimeToTimestamp ...
func TimeToTimestamp(t time.Time) uint64 {
	return uint64(t.Unix())
}

// TimeToTimestampMs ...
func TimeToTimestampMs(t time.Time) uint64 {
	return uint64(t.Unix()) * 1000
}

// TimestampToTime ...
func TimestampToTime(ts uint64) time.Time {
	return time.Unix(int64(ts), 0)
}

// TimestampMsToTime ...
func TimestampMsToTime(ts uint64) time.Time {
	ms := ts / 1000
	return time.Unix(int64(ms), 0)
}

// CurrentTimestamp ...
func CurrentTimestamp() uint64 {
	return TimeToTimestamp(time.Now())
}

// CurrentTimestampMs ...
func CurrentTimestampMs() uint64 {
	return TimeToTimestampMs(time.Now())
}
