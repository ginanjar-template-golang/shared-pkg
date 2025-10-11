package utils

import (
	"sync"

	"github.com/google/uuid"
)

var (
	requestID string
	mu        sync.RWMutex
)

// SetRequestID menyimpan Request ID (biasanya dipanggil di middleware)
func SetRequestID(id string) {
	mu.Lock()
	defer mu.Unlock()
	requestID = id
}

func GetRequestID() string {
	mu.RLock()
	id := requestID
	mu.RUnlock()

	if id == "" {
		id = NewRequestID()
		SetRequestID(id)
	}

	return id
}

func NewRequestID() string {
	return uuid.NewString()
}

func ResetRequestID() {
	mu.Lock()
	defer mu.Unlock()
	requestID = ""
}
