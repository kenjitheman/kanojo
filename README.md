<h2 align="center">kanojo - image generator tgbot</h2>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="250" alt="go logo"  />
  <img width="10" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="250" alt="docker logo"  />
</div>

##
- **works only if you have some gas on your openai account (to generate images)**
##

## project structure:

```
├── cmd
│   └── main.go
├── core
│   └── core.go
├── Dockerfile
├── go.mod
├── go.sum
├── openai
│   └── openai.go
├── README.md
└── tg
    └── tg.go
```

## installation

```
git clone https://github.com/kenjitheman/kanojo
```

## usage

- create .env file and inside you should create env variable with your api key:

```
TELEGRAM_API_TOKEN=YOUR_TOKEN
OPENAI_API_TOKEN=YOUR_API_TOKEN
```

- then you should uncomment commented lines in tg/tg.go
	- **( ! you need uncomment commented lines only if you using this way !)**

```
package tg

import (
	"fmt"
	"log"
	"os"

	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	// "github.com/joho/godotenv"

)

var (
	isBotRunning  bool
	creatorChatID int64
)

func Start() {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	fmt.Println("[ERROR] error loading .env file")
	// 	log.Panic(err)
	// }
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
```

- you can also run it using docker:
	- you need to paste your api key in dockerfile:

```
ENV TELEGRAM_API_TOKEN=YOUR_TOKEN
ENV OPENAI_API_TOKEN=YOUR_API_TOKEN
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
