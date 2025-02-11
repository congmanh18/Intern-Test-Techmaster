package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/russross/blackfriday/v2"
)

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

// Đọc API Key từ .env
func loadAPIKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️ Không tìm thấy file .env, kiểm tra lại!")
	}
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ API Key chưa được thiết lập. Hãy kiểm tra file .env")
	}
	return apiKey
}

// Gọi Groq API
func callGroqAPI(prompt, model string) (string, error) {
	apiKey := loadAPIKey() // Lấy API Key từ .env

	url := "https://api.groq.com/openai/v1/chat/completions"

	// Chuẩn bị dữ liệu JSON
	requestBody, err := json.Marshal(GroqRequest{
		Messages: []Message{{Role: "user", Content: prompt}},
		Model:    model,
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
	var response GroqResponse
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

func main() {
	app := iris.New()

	// Serve static files
	app.HandleDir("/static", "./static")

	// Trang chính
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	// API xử lý prompt
	app.Post("/ask", func(ctx iris.Context) {
		type Request struct {
			Prompt string `json:"prompt"`
			Model  string `json:"model"`
		}

		var req Request
		if err := ctx.ReadJSON(&req); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": "Lỗi khi đọc dữ liệu"})
			return
		}

		response, err := callGroqAPI(req.Prompt, req.Model)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		// Chuyển Markdown thành HTML
		htmlResponse := string(blackfriday.Run([]byte(response)))

		ctx.JSON(iris.Map{"response": htmlResponse})
	})

	app.RegisterView(iris.HTML("./templates", ".html"))
	log.Println("🚀 Server đang chạy tại http://localhost:8080")
	app.Listen(":8080")
}
