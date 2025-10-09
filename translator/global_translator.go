package translator

import (
	"fmt"
	"sync"
)

var (
	globalTranslator *Translator
	mu               sync.RWMutex
)

func SetGlobalTranslator(t *Translator) {
	mu.Lock()
	defer mu.Unlock()
	globalTranslator = t
}

func GetGlobalTranslator() *Translator {
	mu.RLock()
	defer mu.RUnlock()
	if globalTranslator == nil {
		fmt.Println("[translator] warning: global translator not set, using default English")
		globalTranslator = NewTranslator("./translator/messages/en.json")
	}
	return globalTranslator
}

func InitDefaultTranslator(lang string) *Translator {
	var t *Translator
	switch lang {
	case "id":
		t = NewTranslator("./translator/messages/id.json")
	default:
		t = NewTranslator("./translator/messages/en.json")
	}
	SetGlobalTranslator(t)
	return t
}
