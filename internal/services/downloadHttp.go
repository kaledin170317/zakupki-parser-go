package services

import (
	"fmt"
	"io"
	"net/http"
)

func FetchHTML(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %w", err)
	}

	// Установим заголовок User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; ZakupkiBot/1.0)")

	resp, err := getClient().Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("неожиданный статус: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения тела ответа: %w", err)
	}

	return string(body), nil
}
