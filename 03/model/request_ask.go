package model

type Request struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}
