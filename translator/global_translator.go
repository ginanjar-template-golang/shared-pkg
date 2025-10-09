package translator

import (
	_ "embed"
	"fmt"
	"sync"
)

var (
	globalTranslator *Translator
	mu               sync.RWMutex
)

//go:embed messages/en.json
var enJSON []byte

//go:embed messages/id.json
var idJSON []byte

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
		globalTranslator = NewTranslatorFromBytes(enJSON)
	}
	return globalTranslator
}

func InitDefaultTranslator(lang string) *Translator {
	var t *Translator
	switch lang {
	case "id":
		t = NewTranslatorFromBytes(idJSON)
	default:
		t = NewTranslatorFromBytes(enJSON)
	}
	SetGlobalTranslator(t)
	return t
}
