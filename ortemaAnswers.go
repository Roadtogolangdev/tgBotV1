package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"strings"
)

const TOKEN = "7139983209:AAEJOD-EZnFgl5abmWpujmbiqgO3k7pqQtg"

var bot *tgbotapi.BotAPI
var chatID int64

func connectTG() {
	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN) // Используем присваивание вместо короткой декларации
	if err != nil {
		log.Panic("Чето отвалилось, скорее всего телега", err)
	}
}

func sendMsg(msg string) {
	msgConfig := tgbotapi.NewMessage(chatID, msg)
	bot.Send(msgConfig)
}

var myNames = []string{"артемчик", "артем", "артемида", "тема"}

func isMsgForMe(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}
	msgToMe := strings.ToLower(update.Message.Text)
	for _, name := range myNames {
		if strings.Contains(msgToMe, name) {
			return true
		}
	}
	return false
}

var answers = []string{
	"Когда работаешь над сложным проектом весь день, а потом понимаешь, что проще всего... выпить пивка.",
	"Учился кодить, учился, и наконец понял: главное в жизни — это вовремя выпить пивка.",
	"Провел час, решая баг, и тут понял, что пора... выпить пивка.",
	"Смотрел тут видео про продуктивность и осознал, что лучший способ её повысить — это выпить пивка.",
	"Друзья, жизнь коротка, и нет ничего лучше, чем... выпить пивка.",
	"Стал гуру в Go, и понял, что в финале все равно буду... выпить пивка.",
	"Искал смысл жизни, и в конце концов нашел его: выпить пивка.",
	"После долгих размышлений и философских споров решили: лучшее решение — это выпить пивка.",
	"Борешься с трудностями? Не забудь важное правило: всегда можно... выпить пивка.",
	"Выучил все алгоритмы и структуры данных, а затем осознал, что главное в жизни — это... выпить пивка.",
}

func getMyAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatID, getMyAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectTG()
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start" {
			chatID = update.Message.Chat.ID
			sendMsg("Ты можешь задать мне вопрос, но скорее всего после этого мы пойдем пить пивко. Например, \"Артем, ты уже начал работу над той штукой?\" ")
		} else if update.Message != nil && isMsgForMe(&update) {
			sendAnswer(&update)
		} else if update.Message != nil && !isMsgForMe(&update) {
			sendMsg("Сообщение не содержит имени, но все равно... " + getMyAnswer())
		}
	}
}
