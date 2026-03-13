package middleware

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func Invoke() {
	// 初始化客户端 (兼容 OpenAI 协议的模型，如 DeepSeek, Moonshot 等都可以用)
	llm, err := openai.New(
		openai.WithBaseURL("https://api.deepseek.com/v1"), // 替换为国内大模型地址
		openai.WithToken(""),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	prompt := "你是玄幻小说家，请续写：萧炎大喝一声，手中的尺子..."

	// 流式生成
	resp, err := llm.Call(ctx, prompt,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk)) // 这里可以通过 SSE 发送给前端
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESP : ", resp)
}
