package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	//"github.com/joho/godotenv"
	"github.com/kenjitheman/kanojo/openai"
	"os"
)

func Start() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//fmt.Println("[ERROR] error loading .env file")
	//}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		fmt.Println("[ERROR] error creating bot: ", err)
	}

	bot.Debug = true
	isBotRunning = false

	fmt.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start", "/start":
			isBotRunning = true
			if isBotRunning {
				msg.Text = AlreadyRunningMsg
			}
			msg.Text = StartMsg
			msg.ReplyMarkup = generalKeyboard
		case "help", "/help":
			if isBotRunning {
				msg.Text = HelpMsg
				msg.ReplyMarkup = generalKeyboard
			}
		case "generator", "/generator":
			if isBotRunning {
				msg.Text = YourPromptMsg
				bot.Send(msg)

				for {
					response := <-updates
					if response.Message == nil {
						continue
					}
					if response.Message.Chat.ID != update.Message.Chat.ID {
						continue
					}

					prompt := response.Message.Text
					imageUrl := openai.GenerateImage(prompt)
					msg.Text = imageUrl

					break
				}
			}

		case "stop", "/stop":
			if isBotRunning {
				msg.Text = StopMsg
				msg.ReplyMarkup = startKeyboard
				isBotRunning = false
			}
			isBotRunning = false
		case "bug_report", "/bug_report", "bug report":
			if isBotRunning {
				msg.Text = PlsDescribeMsg
				bot.Send(msg)
				for {
					response := <-updates
					if response.Message == nil {
						continue
					}
					if response.Message.Chat.ID != update.Message.Chat.ID {
						continue
					}
					description := response.Message.Text
					msg.Text = ThxMsg
					bot.Send(msg)
					supportMsg := tgbotapi.NewMessage(
						creatorChatID,
						fmt.Sprintf(
							BugReportMsg,
							update.Message.From.UserName,
							description,
						),
					)
					bot.Send(supportMsg)
					break
				}
				continue
			}
			isBotRunning = false

		default:
			if isBotRunning {
				msg.Text = IdkMsg
			}
		}
		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("[ERROR] error sending message")
		}
	}
}
