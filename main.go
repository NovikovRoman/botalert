package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"net/url"
)

var (
	token   = kingpin.Arg("token", "Токен бота (000:XX-XX-XX-XX).").Required().String()
	to      = kingpin.Arg("to", "ID получателя (123456789, -123456789).").Required().Int64()
	message = kingpin.Arg("msg", "Текст сообщения.").Required().String()

	proxy   = kingpin.Flag("proxy", "Socks-прокси.").Short('p').String()
	subject = kingpin.Flag("subject", "Тема сообщения.").Short('s').String()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

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
