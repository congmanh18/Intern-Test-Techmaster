package groq

import (
	"03/model"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Gọi Groq API
func CallGroqAPI(url, apiKey, prompt, chatModel string) (string, error) {
	// Chuẩn bị dữ liệu JSON
	requestBody, err := json.Marshal(model.GroqRequest{
		Messages: []model.Message{{Role: "user", Content: prompt}},
		Model:    chatModel,
	})
	if err != nil {
		return "", err
	}

	// Gửi request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Đọc phản hồi
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// In response để debug
	log.Println("🔹 Response từ API:", string(body))

	// Parse JSON response
	var response model.GroqResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	// Trả về nội dung chat
	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}
	return "Không có phản hồi từ API", nil
}
