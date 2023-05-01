package myBot

import (
	"AquaBot/packages/graph"
	"AquaBot/packages/structs"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type MyBot struct{
	bot * tgbotapi.BotAPI
	vars * structs.TypeVars
	notifiedChats []int64
}

func New(token string, vars *structs.TypeVars)* MyBot{
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	return &MyBot{bot :bot, vars : vars}
}



func (mb *MyBot) Start() {
	go func(){ // Идея такая бот: бот будет помнить последних 5 написавших
					//человек и слать им сообщение в случае аварии
		var canSand bool


		for{
			mb.vars.AlarmsMutex.Lock()
			if mb.vars.Alarms["Котёл1"]  == 0 {
				canSand = true
			}
			if canSand && mb.vars.Alarms["Котёл1"] != 0{
				canSand = false
				for _, chatID := range mb.notifiedChats{
					msg := tgbotapi.NewMessage(chatID, "Авария котла 1")
					mb.bot.Send(msg)
				}
			}
			mb.vars.AlarmsMutex.Unlock()
			time.Sleep(time.Second * 60)
		}
	}()
	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates := mb.bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil { // ignore non-Message updates
				continue
			}
			mb.ChatsToNotifie(update.Message.Chat.ID)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			var path tgbotapi.FilePath = "./output.png"
			msgPic := tgbotapi.NewPhoto(update.Message.Chat.ID, path)

			switch update.Message.Text {
			case "/start":
				msg.ReplyMarkup = numericKeyboard
			case "Система":
				var errors string
				for key, value := range structs.Vars.Alarms{
					if value > 0{
						errors += fmt.Sprintf(" %v ", key)
					}
				}

				msg.Text = fmt.Sprintf("Cистема:\n Связь с ПЛК: %v\n Аварии:%v\n Температура снаружи: %v\n", mb.vars.ConnectionState, errors, mb.vars.TempOutdoorNow)
			case "Котлы":
				msg.Text = fmt.Sprintf("КОТЛЫ:\n  Котёл 1: %v\n  Котёл 2: %v\n", mb.vars.Boiler1State, mb.vars.Boiler2State)
			case "Отопление":
				var array []float64
				mb.vars.HTMutex.Lock()
				array = mb.vars.HTArray
				mb.vars.HTMutex.Unlock()
				msg.Text = fmt.Sprintf("ОТОПЛЕНИЕ: \n Температура сейчас: %v\n", array[11])
				graph.Draw("Температура Отопление", array)
			case "ГВС":
				var array []float64
				mb.vars.GVSTMutex.Lock()
				array = mb.vars.GVSTArray
				mb.vars.GVSTMutex.Unlock()
				msg.Text = fmt.Sprintf("ГВС: \n Температура сейчас: %v\n", array[11])
				graph.Draw("Температура ГВС", array)
			case "Подпитка":
				msg.Text = fmt.Sprintf("ПОДПИТКА:\n Давление в системе: %v\n Включений за час: %v\n", mb.vars.SystemPress, mb.vars.RechargeCount)

			case "СБРОС":
				msg.Text = fmt.Sprintf("Аварии сброшены!")
			default:
				msg.Text = fmt.Sprintf("Команда не распознана..")
			}

			if _, err := mb.bot.Send(msg); err != nil {
				log.Panic(err)
			}
			if update.Message.Text == "Отопление" || update.Message.Text == "ГВС"{
				if _, err := mb.bot.Send(msgPic); err != nil {
					log.Panic(err)
				}
			}
		}
	}()
}

var (
	numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Система"),
			tgbotapi.NewKeyboardButton("Котлы"),
			tgbotapi.NewKeyboardButton("Отопление"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ГВС"),
			tgbotapi.NewKeyboardButton("Подпитка"),
			tgbotapi.NewKeyboardButton("СБРОС"),
		),
	)
)

func (mb *MyBot) ChatsToNotifie(chatID int64){
	for _, chID := range mb.notifiedChats{
		if chatID == chID{
			return
		}
	}
	mb.notifiedChats = append(mb.notifiedChats, chatID )
}