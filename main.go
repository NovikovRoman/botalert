package main

import (
	"flag"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	token = flag.String("token", "", "Токен бота.")
	proxy = flag.String("proxy", "", "Прокси (socks5://[user]:[pass]@[host]:[port]). Необязательно.")

	to      = flag.Int64("to", 0, "ID получателя.")
	subject = flag.String("subject", "", "Тема сообщения. Необязательно.")
	message = flag.String("message", "", "Текст сообщения.")

	help = flag.Bool("h", false, "Помощь.")
)

func main() {
	flag.Parse()

	if *help || *token == "" || *to == 0 || *message == "" {
		outHelp()
		os.Exit(0)
	}

	tbTransport := &http.Transport{}
	if *proxy != "" {
		tbProxyURL, err := url.Parse(*proxy)
		if err != nil {
			log.Fatalln(err)
		}

		tbTransport.Proxy = http.ProxyURL(tbProxyURL)
	}

	client := &http.Client{Transport: tbTransport}
	bot, err := tgbotapi.NewBotAPIWithClient(*token, client)
	if err != nil {
		log.Fatalln(err)
	}

	text := ""
	if *subject != "" {
		text += "*" + *subject + "*\n"
	}

	msg := tgbotapi.NewMessage(*to, text+*message)
	msg.ParseMode = "markdown"
	if _, err := bot.Send(msg); err != nil {
		log.Fatalln(err)
	}
}

func outHelp() {
	fmt.Print(`
Ключи:
	-token=…   - токен бота (000:XX-XX-XX-XX).
	-proxy=…   - socks-прокси. Необязательно.
	-to=…      - ID получателя (123456789, -123456789).
	-subject=… - тема сообщения. Необязательно.
	-message=… - текст сообщения.
`)
}
