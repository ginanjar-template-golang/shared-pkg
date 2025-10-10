package translator

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type Translator struct {
	messages map[string]string
}

var (
	// Global default translator (English)
	globalTranslator *Translator
	mu               sync.RWMutex
)

// ================== Embed JSON files ==================

// Errors
//
//go:embed messages/errors/en.json
var enErrorsJSON []byte

//go:embed messages/errors/id.json
var idErrorsJSON []byte

// Success
//
//go:embed messages/success/en.json
var enSuccessJSON []byte

//go:embed messages/success/id.json
var idSuccessJSON []byte

// ================== Helper Functions ==================

// newTranslatorFromBytes membuat Translator dari JSON bytes
func newTranslatorFromBytes(bytesList ...[]byte) *Translator {
	messages := make(map[string]string)
	for _, b := range bytesList {
		var tmp map[string]string
		if err := json.Unmarshal(b, &tmp); err != nil {
			fmt.Printf("[translator] failed to unmarshal JSON: %v\n", err)
			continue
		}
		for k, v := range tmp {
			messages[k] = v
		}
	}
	return &Translator{messages: messages}
}

// newTranslator load semua messages untuk satu bahasa
func newTranslator(lang string) *Translator {
	switch strings.ToLower(lang) {
	case "id", "id-id":
		return newTranslatorFromBytes(idErrorsJSON, idSuccessJSON)
	default:
		return newTranslatorFromBytes(enErrorsJSON, enSuccessJSON)
	}
}

// ================== Public API ==================

// InitGlobalTranslator set default global translator
func InitGlobalTranslator(lang string) {
	t := newTranslator(lang)
	mu.Lock()
	defer mu.Unlock()
	globalTranslator = t
	fmt.Printf("[translator] global translator initialized with lang=%s\n", lang)
}

// GetGlobalTranslator thread-safe, fallback English
func GetGlobalTranslator() *Translator {
	mu.RLock()
	defer mu.RUnlock()
	if globalTranslator == nil {
		// fallback default English
		t := newTranslator("en")
		mu.RUnlock()
		mu.Lock()
		globalTranslator = t
		mu.Unlock()
		mu.RLock()
		fmt.Println("[translator] global translator not set, initialized default English")
	}
	return globalTranslator
}

// GetMessageGlobal ambil message dari global translator
func GetMessageGlobal(key string) string {
	t := GetGlobalTranslator()
	if msg, ok := t.messages[key]; ok {
		return msg
	}
	return key
}

// GetMessageByLang ambil message sesuai bahasa (per request)
func GetMessageByLang(key string, lang ...string) string {
	selectedLang := "en"
	if len(lang) > 0 && lang[0] != "" {
		selectedLang = lang[0]
	}
	t := newTranslator(selectedLang)
	if msg, ok := t.messages[key]; ok {
		return msg
	}
	return key
}
