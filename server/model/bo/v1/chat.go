package bo

import "time"

type Snapshot struct {
	Count    int       `json:"count"`
	LastAt   time.Time `json:"lastAt"`
	LastMsg  []string  `json:"lastMsg"`
	Username string    `json:"username"`
	UID      int       `json:"uid"`
}
