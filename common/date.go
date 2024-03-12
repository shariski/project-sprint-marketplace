package common

import "time"

func GetDateNowUTCFormat() string {
	return time.Now().UTC().Format(time.RFC3339)
}
