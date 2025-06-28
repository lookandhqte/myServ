package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(nums []int, target int) []int {
	resultArr := make([]int, 0, 2)
	// значение число, ключ айди в массиве
	mapForSearch := make(map[int]int, 0)
	for id, val := range nums {
		mapForSearch[val] = id
	}
	targetNum := 0
	for id, val := range nums {
		targetNum = target - val
		if idOfVal, ok := mapForSearch[targetNum]; ok {
			resultArr = append(resultArr, id)
			resultArr = append(resultArr, idOfVal)
			break
		}
	}

	return resultArr[:]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter target: ")
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Enter numbers (space separated): ")
	scanner.Scan()
	input := scanner.Text()

	strNums := strings.Split(input, " ")
	nums := make([]int, 0, len(strNums))
	for _, val := range strNums {
		numVal, _ := strconv.Atoi(val)
		nums = append(nums, numVal)
	}
	// numsArr := []int{1, 0, 15, 2, 89, 98}
	//[0, 5]
	fmt.Print("Result: ")
	fmt.Print(twoSum(nums[:], target))
}

// func isPalindrome(x int) bool {
// 	var reversedNum = []int{}
//     if x < 0 {
// 		return false
// 	}
// 	copyNum := x
// 	for copyNum != 0 {
// 		if copyNum < 10 {
// 			reversedNum = append(reversedNum, copyNum)
// 			break
// 		}
// 		reversedNum = append(reversedNum, copyNum%10)
// 		copyNum = copyNum / 10
// 	}
// 	reversed := 0
// 	diff := 1
// 	for i := len(reversedNum) - 1; i >= 0; i-- {
// 		reversed += reversedNum[i] * diff
// 		diff *= 10
// 	}
// 	if reversed == x {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// package main

// import (
// 	"context"
// 	"crypto/tls"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// 	openai "github.com/sashabaranov/go-openai"
// )

// // Структура для инструментов из instruments.json
// type Instrument struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// // Загрузить инструменты из JSON файла
// func loadInstruments(path string) ([]Instrument, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var instruments []Instrument
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&instruments)
// 	return instruments, err
// }

// // Конвертировать голосовое сообщение в текст с помощью Whisper
// func transcribeVoice(bot *tgbotapi.BotAPI, fileID string) (string, error) {
// 	// Скачивание файла из Telegram
// 	fileURL, err := bot.GetFileDirectURL(fileID)
// 	if err != nil {
// 		return "", err
// 	}

// 	resp, err := http.Get(fileURL)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// Использование Whisper API для транскрипции
// 	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
// 	req := openai.AudioRequest{
// 		Model:    openai.Whisper1,
// 		FilePath: "audio.ogg", // Имя не важно для потока
// 		Reader:   resp.Body,
// 	}
// 	response, err := client.CreateTranscription(context.Background(), req)
// 	if err != nil {
// 		return "", err
// 	}
// 	return response.Text, nil
// }

// // Сформировать промпт с инструментами для DeepSeek
// func buildPrompt(userInput string, instruments []Instrument) string {
// 	var builder strings.Builder

// 	// Системный промпт с описанием инструментов
// 	builder.WriteString("Ты ИИ-ассистент с доступом к инструментам. Вот список доступных инструментов:\n")
// 	for _, inst := range instruments {
// 		builder.WriteString(fmt.Sprintf("- %s: %s\n", inst.Name, inst.Description))
// 	}

// 	builder.WriteString("\nИнструкции:\n1. Определи какие инструменты нужны для решения задачи\n")
// 	builder.WriteString("2. Объясни как ты будешь их использовать\n3. Ответь на запрос пользователя\n\n")
// 	builder.WriteString(fmt.Sprintf("Запрос пользователя: %s", userInput))

// 	return builder.String()
// }

// // Отправить запрос в DeepSeek API
// func queryDeepSeek(prompt string) (string, error) {
// 	client := openai.NewClient(os.Getenv("DEEPSEEK_API"))

// 	resp, err := client.CreateChatCompletion(
// 		context.Background(),
// 		openai.ChatCompletionRequest{
// 			Model: openai.GPT3Dot5Turbo,
// 			Messages: []openai.ChatCompletionMessage{
// 				{
// 					Role:    openai.ChatMessageRoleUser,
// 					Content: prompt,
// 				},
// 			},
// 		},
// 	)

// 	if err != nil {
// 		return "", err
// 	}

// 	return resp.Choices[0].Message.Content, nil
// }

// func projectHandler(w http.ResponseWriter, r *http.Request) {
// 	projectName := r.URL.Path[len("/project/"):]

// 	// Здесь можно загружать данные проекта из БД
// 	// и динамически генерировать страницу

// 	html := fmt.Sprintf(`
// 	<!DOCTYPE html>
// 	<html>
// 	<head>
// 		<title>Проект %s</title>
// 	</head>
// 	<body>
// 		<h1>Страница проекта: %s</h1>
// 		<p>Детальное описание проекта...</p>
// 		<a href="/">На главную</a>
// 	</body>
// 	</html>
// 	`, projectName, projectName)

// 	fmt.Fprint(w, html)
// }

// func main() {
// 	// Обработчик статических файлов (изображения, CSS, JS)
// 	fs := http.FileServer(http.Dir("./static"))
// 	// http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	http.Handle("/static/", http.StripPrefix("/static/",
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			if strings.HasPrefix(r.URL.Path, ".") {
// 				http.NotFound(w, r)
// 				return
// 			}
// 			w.Header().Set("Cache-Control", "public, max-age=31536000")
// 			fs.ServeHTTP(w, r)
// 		}),
// 	))
// 	// Общий мультиплексор для обоих серверов
// 	mux := http.NewServeMux()

// 	// Основные обработчики
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "index.html")
// 	})
// 	mux.HandleFunc("/project/", projectHandler)

// 	// HTTPS-сервер (основной)
// 	go func() {
// 		server := &http.Server{
// 			Addr:         ":443",
// 			Handler:      mux, // используем общий обработчик
// 			ReadTimeout:  10 * time.Second,
// 			WriteTimeout: 10 * time.Second,
// 			TLSConfig: &tls.Config{
// 				MinVersion: tls.VersionTLS12,
// 			},
// 		}
// 		log.Println("HTTPS сервер запущен на :443")
// 		log.Fatal(server.ListenAndServeTLS("fullchain.pem", "privkey.pem"))
// 	}()

// 	// HTTP-сервер (ТОЛЬКО для редиректа)
// 	redirector := &http.Server{
// 		Addr: ":80",
// 		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// Редирект на HTTPS-версию того же URL
// 			target := "https://" + r.Host + r.URL.String()
// 			http.Redirect(w, r, target, http.StatusMovedPermanently)
// 		}),
// 	}
// 	log.Println("HTTP сервер запущен на :80 (только редирект на HTTPS)")
// 	log.Fatal(redirector.ListenAndServe())
// 	// Загрузка переменных окружения
// 	godotenv.Load()

// 	// Загрузка инструментов
// 	instruments, err := loadInstruments("instruments.json")
// 	if err != nil {
// 		log.Fatalf("Ошибка загрузки инструментов: %v", err)
// 	}

// 	// Инициализация Telegram бота
// 	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API"))
// 	if err != nil {
// 		log.Fatalf("Ошибка инициализации бота: %v", err)
// 	}

// 	bot.Debug = true
// 	log.Printf("Авторизован как %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60
// 	updates := bot.GetUpdatesChan(u)

// 	// Обработка входящих сообщений
// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}

// 		var userText string
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

// 		// Обработка голосовых сообщений
// 		if update.Message.Voice != nil {
// 			text, err := transcribeVoice(bot, update.Message.Voice.FileID)
// 			if err != nil {
// 				msg.Text = "Ошибка распознавания голоса: " + err.Error()
// 				bot.Send(msg)
// 				continue
// 			}
// 			userText = text
// 		} else {
// 			userText = update.Message.Text
// 		}

// 		// Формирование промпта с инструментами
// 		prompt := buildPrompt(userText, instruments)

// 		// Отправка запроса в DeepSeek
// 		response, err := queryDeepSeek(prompt)
// 		if err != nil {
// 			msg.Text = "Ошибка запроса к DeepSeek: " + err.Error()
// 			bot.Send(msg)
// 			continue
// 		}

// 		// Отправка ответа пользователю
// 		msg.Text = response
// 		bot.Send(msg)
// 	}
// }
