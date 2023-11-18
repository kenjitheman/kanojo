## Image generator tgbot

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="250" alt="go logo"  />
</div>

- **Works only if you have some gas on your openai account (to generate images)**

## Project structure:

```go
kanojo
│
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

## Installation

```sh
git clone https://github.com/kenjitheman/kanojo
```

## Usage

- Create .env file and inside you should create env variable with your api key:

```.env
TELEGRAM_API_TOKEN=YOUR_TOKEN
OPENAI_API_TOKEN=YOUR_API_TOKEN
```

- Then you should uncomment commented lines in bot.go
	- **( ! You need uncomment commented lines only if you using this way !)**

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

- You can also run it using docker:
	- You need to paste your API key in Dockerfile:

```dockerfile
ENV TELEGRAM_API_TOKEN=YOUR_TOKEN
ENV OPENAI_API_TOKEN=YOUR_API_TOKEN
```

## Contributing

- Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## License

- [MIT](https://choosealicense.com/licenses/mit/)
