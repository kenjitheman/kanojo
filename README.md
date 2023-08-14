<h2 align="center">fantajier - image generator tgbot</h2>

###

<img align="right" height="200" src="https://media.tenor.com/NBS9cKKEeC8AAAAC/rainbow.gif"  />

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="25" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

<div align="center">
  <img src="https://img.shields.io/static/v1?message=fantajier&logo=telegram&label=&color=909ef7&logoColor=white&labelColor=&style=for-the-badge" height="30" alt="telegram logo"  />
</div>

## project structure:

```
├── cmd
│   └── main.go
├── Dockerfile
├── gdrive
│   └── gdrive.go
├── go.mod
├── go.sum
├── openai
│   └── openai.go
├── README.md
└── tg
    └── bot.go
```

## installation

use git clone:

```
git clone https://github.com/kenjitheman/fantajier
```

## usage

- create .env file and inside you should create env variable with your api key:

```
TELEGRAM_API_TOKEN=YOUR_TOKEN
OPENAI_API_TOKEN=YOUR_API_TOKEN
```

- then you should uncomment commented lines in tg/tg.go ( ! you need uncomment
  commented lines only if you using this way !) ->

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

### you can also run it using docker ->

- you need to paste your api key in dockerfile ->

```
ENV TELEGRAM_API_TOKEN=YOUR_TOKEN
ENV OPENAI_API_TOKEN=YOUR_API_TOKEN
```

- then run it ->

```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## contributing

- pull requests are welcome. For major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
