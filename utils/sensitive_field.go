package utils

import (
	"strings"
)

var sensitiveFields = []string{
	"password", "token", "access_token", "refresh_token", "authorization",
	"secret", "client_secret", "api_key", "jwt", "bearer",
}

func SanitizeMap(m map[string]any) map[string]any {
	for k, v := range m {
		lowerK := strings.ToLower(k)
		for _, sf := range sensitiveFields {
			if strings.Contains(lowerK, sf) {
				m[k] = "[REDACTED]"
				break
			}
		}
		if sub, ok := v.(map[string]any); ok {
			m[k] = SanitizeMap(sub)
		}
	}
	return m
}
