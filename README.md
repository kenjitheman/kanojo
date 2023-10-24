## image generator tgbot

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="250" alt="go logo"  />
</div>

- **works only if you have some gas on your openai account (to generate images)**

## project structure:

```go
.
├── bot
│   ├── bot.go
│   ├── keyboards.go
│   └── vars.go
├── core
│   ├── core.go
│   └── core_test.go
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── openai
│   ├── openai.go
│   └── openai_test.go
└── README.md
```

## installation

```shell
git clone https://github.com/kenjitheman/kanojo
```

## usage

- create .env file and inside you should create env variable with your api key:

```.env
TELEGRAM_API_TOKEN=YOUR_TOKEN
OPENAI_API_TOKEN=YOUR_API_TOKEN
```

- then you should uncomment commented lines in bot.go
	- **( ! you need uncomment commented lines only if you using this way !)**

```go
// "github.com/joho/godotenv"
```

```go
//err := godotenv.Load("../.env")
//if err != nil {
    //fmt.Println("[ERROR] error loading .env file")
    //log.Panic(err)
//}
```

- you can also run it using docker:
	- you need to paste your api key in dockerfile:

```dockerfile
ENV TELEGRAM_API_TOKEN=YOUR_TOKEN
ENV OPENAI_API_TOKEN=YOUR_API_TOKEN
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

- please make sure to update tests as appropriate

## license

- [MIT](https://choosealicense.com/licenses/mit/)
