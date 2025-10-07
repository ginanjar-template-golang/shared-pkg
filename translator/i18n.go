package translator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	cache = make(map[string]map[string]string)
	mu    sync.RWMutex
)

// Load locale dari JSON
func loadLocale(lang string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := cache[lang]; exists {
		return nil
	}

	path := filepath.Join("shared-pkg", "translator", "locales", fmt.Sprintf("%s.json", lang))
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to load locale %s: %v", lang, err)
	}

	var data map[string]string
	if err := json.Unmarshal(file, &data); err != nil {
		return fmt.Errorf("invalid json in locale %s: %v", lang, err)
	}

	cache[lang] = data
	return nil
}

// T = translate
func T(lang, code string) string {
	if err := loadLocale(lang); err != nil {
		_ = loadLocale("en")
		if msg, ok := cache["en"][code]; ok {
			return msg
		}
		return code
	}

	mu.RLock()
	defer mu.RUnlock()

	if msg, ok := cache[lang][code]; ok {
		return msg
	}
	return code
}
