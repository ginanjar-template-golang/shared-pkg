package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var messages = map[string]map[string]string{}

// Load all locale files (en.json, id.json)
func InitLocales(basePath string) error {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return fmt.Errorf("failed to read locales: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			locale := f.Name()[:len(f.Name())-5]
			path := filepath.Join(basePath, f.Name())

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var data map[string]string
			if err := json.Unmarshal(content, &data); err != nil {
				return err
			}

			messages[locale] = data
		}
	}
	return nil
}

// Translate code ke message sesuai locale
func T(locale, code string) string {
	if msg, ok := messages[locale][code]; ok {
		return msg
	}
	// fallback ke English
	if msg, ok := messages["en"][code]; ok {
		return msg
	}
	return code
}
