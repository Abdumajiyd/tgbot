package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/NicoNex/echotron/v3"
	"github.com/kkdai/youtube/v2"
)

const token = "6539696441:AAHHDpsHVfCBuC0pP-IPrr--qWShNXaZd_4"

type Bot struct {
	echotron.API
	chatID int64
}

func NewBot(chatID int64) *Bot {
	return &Bot{
		API:    echotron.NewAPI(token),
		chatID: chatID,
	}
}

func (b *Bot) Update(update *echotron.Update) {
	log.Println("Received update:", update)
	if update.Message != nil {
		message := update.Message
		b.chatID = message.Chat.ID
		if message.Text != "" {
			log.Println("Received message:", message.Text)
			b.handleMessage(message.Text)
		}
	}
}

func (b *Bot) handleMessage(text string) {
	if text == "/start" {
		b.sendMessage("Assaloomu alaykum, Videoni 720p formatda yuklab olish uchun menga YouTube havolasini yuboring.")
	} else if strings.HasPrefix(text, "/download ") {
		url := strings.TrimPrefix(text, "/download ")
		log.Println("Processing download command for URL:", url)
		b.downloadAndSendVideo(url)
	} else if isValidURL(text) {
		log.Println("Processing URL:", text)
		b.downloadAndSendVideo(text)
	} else {
		log.Println("Notog'ri buyruq qabul qilindi")
		b.sendMessage("Noto'g'ri buyruq. Videoni 720p formatda yuklab olish uchun faqat YouTube havolasini o'zini yuboring.")
	}
}

func isValidURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func (b *Bot) downloadAndSendVideo(url string) {
	b.sendMessage("Video yuklab olinmoqda, kuting...")
	videoPath, err := b.downloadVideo(url)
	if err != nil {
		log.Println("Videoni yuklab olishda xatolik yuz berdi:", err)
		b.sendMessage("Video yuklab olinmadi: " + err.Error())
		return
	}
	b.sendMessage("Video muvaqffaqiyatli yuklab olindi: " + videoPath)
	b.sendVideo(videoPath)
}

func (b *Bot) downloadVideo(url string) (string, error) {
	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		return "", fmt.Errorf("error getting video info: %s", err)
	}
	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		return "", fmt.Errorf("error getting video stream: %s", err)
	}
	filePath := fmt.Sprintf("%s.mp4", video.Title)
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()
	_, err = file.ReadFrom(stream)
	if err != nil {
		return "", fmt.Errorf("error saving video: %s", err)
	}
	return filePath, nil
}

func (b *Bot) sendMessage(text string) {
	_, err := b.SendMessage(text, b.chatID, nil)
	if err != nil {
		log.Println("Failed to send message:", err)
	}
}

func (b *Bot) sendVideo(videoPath string) {
	videoFile := echotron.NewInputFilePath(videoPath)
	_, err := b.SendVideo(videoFile, b.chatID, nil)
	if err != nil {
		log.Println("Failed to send video:", err)
	}
}

func main() {
	log.Println("Starting bot...")
	dsp := echotron.NewDispatcher(token, func(chatID int64) echotron.Bot {
		return NewBot(chatID)
	})
	log.Fatal(dsp.Poll())
}


















































// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"github.com/NicoNex/echotron/v3"
// 	"github.com/kkdai/youtube/v2"
// )

// const token = "6539696441:AAHHDpsHVfCBuC0pP-IPrr--qWShNXaZd_4"

// type Bot struct {
// 	echotron.API
// 	chatID int64
// }

// func NewBot(chatID int64) *Bot {
// 	return &Bot{
// 		API:    echotron.NewAPI(token),
// 		chatID: chatID,
// 	}
// }

// func (b *Bot) Update(update *echotron.Update) {
// 	log.Println("Received update:", update)
// 	if update.Message != nil {
// 		message := update.Message
// 		b.chatID = message.Chat.ID

// 		if message.Text != "" {
// 			log.Println("Received message:", message.Text)
// 			b.handleMessage(message.Text)
// 		}
// 	}
// }

// func (b *Bot) handleMessage(text string) {
// 	if text == "/start" {
// 		b.sendMessage("Assaloomu alaykum, Videoni 720p formatda yuklab olish uchun menga YouTube havolasini yuboring.")
// 	} else if strings.HasPrefix(text, "/download ") {
// 		url := strings.TrimPrefix(text, "/download ")
// 		log.Println("Processing download command for URL:", url)
// 		b.downloadAndSendVideo(url)
// 	} else if isValidURL(text) {
// 		log.Println("Processing URL:", text)
// 		b.downloadAndSendVideo(text)
// 	} else {
// 		log.Println("Notog'ri buyruq qabul qilindi")
// 		b.sendMessage("Noto'g'ri buyruq. Videoni 720p formatda yuklab olish uchun YouTube havolasini yuboring.")
// 	}
// }

// func isValidURL(url string) bool {
// 	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
// }

// func (b *Bot) downloadAndSendVideo(url string) {
// 	b.sendMessage("Video yuklab olinmoqda, kuting...")
// 	videoPath, err := b.downloadVideo(url)
// 	if err != nil {
// 		log.Println("Videoni yuklab olishda xatolik yuz berdi:", err)
// 		b.sendMessage("Video yuklab olinmadi: " + err.Error())
// 		return
// 	}
// 	b.sendMessage("Video muvaqffaqiyatli yuklab olindi: " + videoPath)
// 	b.sendVideo(videoPath)
// }

// func (b *Bot) downloadVideo(url string) (string, error) {
// 	client := youtube.Client{}

// 	video, err := client.GetVideo(url)
// 	if err != nil {
// 		return "", fmt.Errorf("error getting video info: %s", err)
// 	}

// 	formats := video.Formats.WithAudioChannels()
// 	stream, _, err := client.GetStream(video, &formats[0])
// 	if err != nil {
// 		return "", fmt.Errorf("error getting video stream: %s", err)
// 	}

// 	filePath := fmt.Sprintf("%s.mp4", video.Title)
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return "", fmt.Errorf("error creating file: %s", err)
// 	}
// 	defer file.Close()

// 	_, err = file.ReadFrom(stream)
// 	if err != nil {
// 		return "", fmt.Errorf("error saving video: %s", err)
// 	}

// 	return filePath, nil
// }

// // func (b *Bot) downloadVideo(url string) (string, error) {
// // 	cmd := exec.Command("youtube-dl", "-f", "bestvideo[height<=720]+bestaudio/best[height<=720]", "-o", "video.%(ext)s", url)
// // 	output, err := cmd.CombinedOutput()
// // 	if err != nil {
// // 		return "", fmt.Errorf("command failed: %s\n%s", err, string(output))
// // 	}

// // 	videoPath := "video.mp4"
// // 	return videoPath, nil
// // }

// func (b *Bot) sendMessage(text string) {
// 	_, err := b.SendMessage(text, b.chatID, nil)
// 	if err != nil {
// 		log.Println("Failed to send message:", err)
// 	}
// }

// func (b *Bot) sendVideo(videoPath string) {
// 	videoFile := echotron.NewInputFilePath(videoPath)

// 	_, err := b.SendVideo(videoFile, b.chatID, nil)
// 	if err != nil {
// 		log.Println("Failed to send video:", err)
// 	}
// }

// func main() {
// 	log.Println("Starting bot...")
// 	dsp := echotron.NewDispatcher(token, func(chatID int64) echotron.Bot {
// 		return NewBot(chatID)
// 	})
// 	log.Fatal(dsp.Poll())
// }
