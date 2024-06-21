package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/sashabaranov/go-openai"
	"github.com/subosito/gotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	_ = gotenv.Load(".env")
	client := openai.NewClient(os.Getenv("OPENAI_TOKEN"))
	chatCompletionStream(client)
	//batchesApi(client)
}

func chatCompletionStream(client *openai.Client) {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 设置响应头
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.Header().Set("Transfer-Encoding", "chunked")
		writer.Header().Set("Cache-Control", "no-cache")
		flusher, ok := writer.(http.Flusher)
		if !ok {
			http.Error(writer, "Streaming not supported!", http.StatusInternalServerError)
			return
		}

		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 50,
			Stream:    true,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: request.URL.Query().Get("msg"),
				},
			},
		}
		stream, err := client.CreateChatCompletionStream(request.Context(), req)
		if err != nil {
			fmt.Printf("CreateChatCompletionStream error: %v\n", err)
			return
		}

		defer stream.Close()
		for {
			select {
			case <-request.Context().Done():
				return
			default:
				response, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					return
				} else if err != nil {
					log.Printf("stream error: %v\r\n", err)
					return
				}
				content := response.Choices[0].Delta.Content
				//需要 \r\n 结尾 https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding#chunked
				writer.Write([]byte(content + "\n"))
				flusher.Flush()
				time.Sleep(200 * time.Millisecond)
			}
		}
	})
	srv := khttp.NewServer(khttp.Address(":8003"), khttp.Timeout(30*time.Second))
	srv.HandleFunc("/chat", handler)

	app := kratos.New(kratos.Server(srv))
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func batchesApi(client *openai.Client) {
	ctx := context.Background()
	//createBatch(ctx, client)

	batchID := ""
	batch, err := client.RetrieveBatch(ctx, batchID)
	if err != nil {
		return
	}
	fmt.Println(batch.Status)
	printResponse(ctx, client, &batch.InputFileID)
	printResponse(ctx, client, batch.OutputFileID)
	printResponse(ctx, client, batch.ErrorFileID)
}

func printResponse(ctx context.Context, client *openai.Client, fileID *string) {
	if fileID != nil {
		content, err := client.GetFileContent(ctx, *fileID)
		if err != nil {
			return
		}
		all, _ := io.ReadAll(content)
		fmt.Println(string(all))
	}
}

func createBatch(ctx context.Context, client *openai.Client) {
	req := openai.CreateBatchWithUploadFileRequest{
		Endpoint: openai.BatchEndpointChatCompletions,
	}
	comments := []string{
		"it's a good bike but if you have a problem after the sale they either do not respond to you or the parts are not available",
		"I ordered 2 Mars 2.0.A blue and an Orange.Blue came first and had shipping damage to the seat post.It came with a flip seat.The Orange came  about 10 days later and didnt have a flip seat.I notified customer service about both issues.They shipped a new seat post but it will not fit the blue bike because it is for a non flip seat.I am still waiting for a fix both both of these problems.\nI do not like the fact that the throttle cannot be used without the peddle assist being on.At time I feel the peddle assist is dangerous.You better not try to make a turn with the peddle assist on.",
		"This was my first E-bike. Love it so far, it has plenty power and range. I use it for hunting on our land. Works well for me, I am very satisfied.",
		"I would definitely recommend this bike. Easy to use. Great battery life, quick delivery!",
		"Slight difficulty setting up bike but it’s perfect and love it’s speed and power",
	}
	prompt := "请分析以下产品评论，并提取出用户提及到的维度值及原因。\n\n\n评论示例：\n```\n这款耳机音质非常好，适合喜欢听音乐的人使用。我每天在通勤路上都会戴着它，降噪效果也很棒。售后服务也非常到位，客服很有耐心地解决了我的问题。唯一的缺点是佩戴时间长了耳朵会有点疼。\n```\n\n预期JSON输出示例：\n```json\n{\n    \"dimensions\": [\n        {\n            \"dimension\": \"使用场景\",\n            \"value\": \"通勤路上\",\n            \"reason\": \"用户每天在通勤路上都会戴着它\"\n        },\n        {\n            \"dimension\": \"使用人群\",\n            \"value\": \"喜欢听音乐的人\",\n            \"reason\": \"用户喜欢听音乐\"\n        },\n        {\n            \"dimension\": \"产品体验-正向观点\",\n            \"value\": \"音质非常好\",\n            \"reason\": \"用户认为耳机音质非常好\"\n        },\n        {\n            \"dimension\": \"产品体验-正向观点\",\n            \"value\": \"降噪效果很棒\",\n            \"reason\": \"用户认为降噪效果很棒\"\n        },\n        {\n            \"dimension\": \"产品体验-负向观点\",\n            \"value\": \"佩戴时间长了耳朵会疼\",\n            \"reason\": \"用户认为佩戴时间长了耳朵会有点疼\"\n        },\n    ]\n}\n}\n```\n请根据以上示例进行分析，并返回JSON格式的结果。"
	for i, comment := range comments {

		req.AddChatCompletion(fmt.Sprintf("req-%d", i), openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: prompt},
				{Role: openai.ChatMessageRoleUser, Content: comment},
			},
			MaxTokens: 2000,
		})
	}
	response, err := client.CreateBatchWithUploadFile(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(response)
}
