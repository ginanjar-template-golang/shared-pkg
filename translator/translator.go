package translator

import (
	"encoding/json"
	"fmt"
	"os"
)

type Translator struct {
	messages map[string]string
}

// Load translation file
func NewTranslator(filePath string) *Translator {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("translator file not found: %s", filePath))
	}
	var msg map[string]string
	json.Unmarshal(data, &msg)
	return &Translator{messages: msg}
}

func NewTranslatorFromBytes(data []byte) *Translator {
	var msg map[string]string
	if err := json.Unmarshal(data, &msg); err != nil {
		panic(fmt.Sprintf("failed to unmarshal translator json: %v", err))
	}
	return &Translator{messages: msg}
}

func (t *Translator) T(key string) string {
	if t == nil {
		return key
	}
	if val, ok := t.messages[key]; ok {
		return val
	}
	return key
}
