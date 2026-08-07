package service

import (
	"bytes"
	"codewars-pretty-stats/internal/config"
	"encoding/json"
	"net/http"
	"time"
)

func AxiomLog(username string, size float64, cfg *config.AxiomConfig) {
	payload := []map[string]interface{}{
		{
			"_time":    time.Now().Format(time.RFC3339),
			"username": username,
			"size":     size,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, cfg.AxiomURL, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+cfg.AxiomToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return
	}
}
