package utils

import (
	"time"
)

const STORE_FILE = "./store.db"

const DATE_LAYOUT = time.RFC3339

type IsDone string

const (
	YES     IsDone = "YES"
	NO      IsDone = "NO"
	UNKNOWN IsDone = "UNKNOWN"
)

func (c IsDone) String() string {
	switch c {
	case YES:
		return "Yes"
	case NO:
		return "No"
	default:
		return "Unknown"
	}
}
