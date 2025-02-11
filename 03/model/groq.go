package model

// Cấu trúc yêu cầu API
type GroqRequest struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
}

// Tin nhắn gửi đi
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Cấu trúc phản hồi từ Groq API
type GroqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type ResponseData struct {
	Dialogue        string
	ImportantWords  string
	TranslatedWords string
}
