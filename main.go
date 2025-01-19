package main

import (
    "bufio"
    "context"
    "fmt"
    "os"

    _ "github.com/joho/godotenv/autoload"
    "github.com/neurocult/agency"
    "github.com/neurocult/agency/providers/openai"
)

func main() {
    // إنشاء الـ Provider
    assistant := openai.New(openai.Params{Key: os.Getenv("OPENAI_API_KEY")}).TextToText(openai.TextToTextParams{Model: "gpt-4o-mini"}).SetPrompt("You are a helpful assistant.")

    messages := []agency.Message{}
    reader := bufio.NewReader(os.Stdin)
    ctx := context.Background()

    for {
        fmt.Print("User: ")

        text, err := reader.ReadString('\n')
        if err != nil {
            panic(err)
        }

        // إنشاء رسالة المستخدم
        input := agency.NewTextMessage(agency.UserRole, text)
        
        // تنفيذ العملية
        answer, err := assistant.SetMessages(messages).Execute(ctx, input)
        if err != nil {
            panic(err)
        }

        // طباعة الرد
        fmt.Println("Assistant:", string(answer.Content()))

        // إضافة الرسائل إلى السجل
        messages = append(messages, input, answer)
    }
}
