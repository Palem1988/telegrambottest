package bot

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	api "telegrambottest/src/bipdev"
	stct "telegrambottest/src/bipdev/structs"
	vocab "telegrambottest/src/bot/vocabulary"
	"telegrambottest/src/db"
	"time"

	//strt "bipbot/src/bipdev/structs"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	startCommand    = "start"
	priceCommand    = "price"
	buyCommand      = "buy"
	sellCommand     = "sell"
	salesCommand    = "lookat"
	getMainMenu     = "getmainmenu"
	settingsMenu    = "settings"
	language        = "language"
	engvocabCommand = "englanguage"
	rusvocabCommand = "ruslanguage"
)

var (
	commands    = make(map[int]string)
	CommandInfo = make(map[int]string)
	CoinToSell  = make(map[int]string)
	PriceToSell = make(map[int]float64)
)

type Dialog struct {
	ChatId    int64
	UserId    int
	MessageId int
	Callback  string
	Text      string
	language  string
	Command   string
}

// Bot is struct for Bot:   - Token: secret token from .env
//							- Api:   Struct App for Rest Api methods
//							- DB:    Postgres DB fro users and user's loots.
//							- Bot:	 tgbotapi Bot(token)
//							- Dlg:   For dialog struct

type Bot struct {
	Token string
	Api   *api.App
	DB    *db.DataBase
	Bot   *tgbotapi.BotAPI
	Dlg   *Dialog
}

//
func InitBot(config stct.Config, dbsql *sql.DB) *Bot {

	b := Bot{
		Token: config.Token,
		DB:    &db.DataBase{},
		Dlg:   &Dialog{},
	}

	// Create table if not exists
	db, err := db.InitDB(dbsql)
	if err != nil {
		log.Fatal(err)
	}

	b.DB = db
	// Define URL
	b.Api = api.InitApp(config.URL)

	// Create new bot
	bot, err := tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		log.Fatal(err)
	}
	b.Bot = bot

	return &b
}

// Run is starting bot.
func (b *Bot) Run() {

	//Set update timeout
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Get updates from bot
	updates, _ := b.Bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		dialog, exist := b.assembleUpdate(update)
		if !exist {
			continue
		}

		b.Dlg = dialog

		if update.Message != nil && update.Message.ReplyToMessage != nil {
			if dialog.Command == "buy" {
				b.Buy()
				continue
			} else if dialog.Command == "sell" {
				b.Sell()
				continue
			}
		}

		if botCommand := b.getCommand(update); botCommand != "" {
			b.RunCommand(botCommand)
			continue
		}

	}
}

// assembleUpdate
func (b *Bot) assembleUpdate(update tgbotapi.Update) (*Dialog, bool) {
	dialog := &Dialog{}
	if update.Message != nil {
		dialog.language = b.DB.GetLanguage(update.Message.Chat.ID)
		dialog.ChatId = update.Message.Chat.ID
		dialog.MessageId = update.Message.MessageID
		dialog.UserId = int(update.Message.Chat.ID)
		dialog.Text = update.Message.Text
	} else if update.CallbackQuery != nil {
		dialog.language = b.DB.GetLanguage(update.CallbackQuery.Message.Chat.ID)
		dialog.ChatId = update.CallbackQuery.Message.Chat.ID
		dialog.MessageId = update.CallbackQuery.Message.MessageID
		dialog.Callback = update.CallbackQuery.ID
		dialog.UserId = int(update.CallbackQuery.Message.Chat.ID)
		dialog.Text = ""
	} else {
		dialog.language = "en"
		return dialog, false
	}

	command, isset := commands[dialog.UserId]
	if isset {
		dialog.Command = command
	} else {
		dialog.Command = ""
	}

	return dialog, true
}

// getCommand returns command from telegram update
func (b *Bot) getCommand(update tgbotapi.Update) string {
	if update.Message != nil {
		if update.Message.IsCommand() {
			fmt.Println("command: ", update.Message.Command())
			return update.Message.Command()
		}
	} else if update.CallbackQuery != nil {
		return update.CallbackQuery.Data
	}

	return ""
}

// RunCommand executes the input command.
func (b *Bot) RunCommand(command string) {
	commands[b.Dlg.UserId] = command
	switch command {

	// "/Start" interacting with the bot, bot description and available commands.
	case startCommand:
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Hello", b.Dlg.language))
		msg.ReplyMarkup = b.newVocabuageKeybord()
		b.Bot.Send(msg)

	case settingsMenu:
		kb := b.newVocabuageKeybord()
		msg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:      b.Dlg.ChatId,
				MessageID:   b.Dlg.MessageId,
				ReplyMarkup: &kb,
			},
			Text: vocab.GetTranslate("Settings", b.Dlg.language),
		}

		b.Bot.Send(msg)

	// engvocabCommand sets english lang for user.
	case engvocabCommand:
		b.DB.SetLanguage(b.Dlg.UserId, "en")
		b.Dlg.language = "en"
		b.SendMenu()

	// rusvocabCommand sets russian lang for user.
	case rusvocabCommand:
		b.DB.SetLanguage(b.Dlg.UserId, "ru")
		b.Dlg.language = "ru"
		b.SendMenu()

	// priceCommand requests the server for the current BIP / USD rate and sends a message to user with the server responce.
	case priceCommand:
		price, err := b.Api.GetPrice()
		if err != nil {
			msg := tgbotapi.NewMessage(b.Dlg.ChatId, err.Error())
			b.Bot.Send(msg)
		}
		ans := fmt.Sprintf(vocab.GetTranslate("Now", b.Dlg.language), price)
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, ans)
		b.Bot.Send(msg)
		b.SendMenu()

	// buyCommand collects data from the user to transmit their request.
	// The user will receive the address for the deposit.
	// After he sends the money he will receive a notification from bot.
	// After the money is confirmed, he will receive another notification from bot.
	case buyCommand:
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Send minter", b.Dlg.language))
		// requests a forced response from the user to collect data to send a request to the server
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: true,
			Selective:  true,
		}
		b.Bot.Send(msg)

	// sellCommand collects data from the user to transmit their request.
	//
	case sellCommand:
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Coin price", b.Dlg.language))
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: true,
			Selective:  true,
		}
		b.Bot.Send(msg)

	// salesCommand sends a request to the database to get user's loots.
	case salesCommand:
		loots, err := b.DB.GetLoots(b.Dlg.UserId)
		if err != nil {
			fmt.Println(err)
			msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Error", b.Dlg.language))
			b.Bot.Send(msg)
		} else if len(loots) == 0 {
			msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Empty loots", b.Dlg.language))
			b.Bot.Send(msg)
		}
		b.ComposeResp(loots)
		b.SendMenu()

	case getMainMenu:
		b.SendMenu()
	}
}

func (b *Bot) SendMenu() {

	kb := b.newMainMenuKeyboard()
	newmsg := tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      b.Dlg.ChatId,
			MessageID:   b.Dlg.MessageId,
			ReplyMarkup: &kb,
		},
		Text: vocab.GetTranslate("Select", b.Dlg.language),
	}

	b.Bot.Send(newmsg)
}

// Buy is function for command "/buy".
// Requests an email from the user and Minter deposit address.
// Requests the "bitcoinDepositAddress" method with the received data.
func (b *Bot) Buy() {
	if strings.Contains(b.Dlg.Text, "@") {
		addr, err := b.Api.GetBTCDeposAddress(CommandInfo[b.Dlg.UserId], "BIP",
			b.Dlg.Text)
		if err != nil {
			b.Dlg.Command = ""
			msg := tgbotapi.NewMessage(b.Dlg.ChatId, err.Error())
			msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{
				RemoveKeyboard: true,
				Selective:      true,
			}
			b.Bot.Send(msg)
			return
		}
		ans := fmt.Sprintf(vocab.GetTranslate("BTC deposit", b.Dlg.language), addr)
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, ans)
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{
			RemoveKeyboard: true,
			Selective:      true,
		}
		b.Dlg.Command = ""
		b.Bot.Send(msg)
		go b.CheckStatusBuy(addr)
		return
	} else {
		CommandInfo[b.Dlg.UserId] = b.Dlg.Text
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Email", b.Dlg.language))
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: true,
			Selective:  true,
		}
		b.Bot.Send(msg)
		return
	}
}

// CheckStatusBuy checks depos BTC and wait 2 confirmations
func (b *Bot) CheckStatusBuy(address string) {
	timeout := time.After(2 * time.Minute)
	tick := time.Tick(3 * time.Second)
	willcoin := 0.
	for {
		select {
		case <-timeout:
			if willcoin == 0. {
				msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("timeout", b.Dlg.language))
				msg.ReplyMarkup = b.newMainMenuKeyboard()
				b.Bot.Send(msg)
				return
			} else {
				continue
			}

		case <-tick:
			stat, err := b.Api.GetBTCDepositStatus(address)
			if err != nil {
				fmt.Println(err)
				msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Error", b.Dlg.language))
				b.Bot.Send(msg)
				return
			}
			if stat.Data.WillReceive != willcoin {
				if willcoin == 0. {
					willcoin = stat.Data.WillReceive
					ans := fmt.Sprintf(vocab.GetTranslate("New deposit", b.Dlg.language), stat.Data.WillReceive)
					msg := tgbotapi.NewMessage(b.Dlg.ChatId, ans)
					b.Bot.Send(msg)
					time.Sleep(60 * time.Second)
				} else {
					ans := fmt.Sprintf(vocab.GetTranslate("Exchange is successful", b.Dlg.language), willcoin)
					msg := tgbotapi.NewMessage(b.Dlg.ChatId, ans)
					b.Bot.Send(msg)
					return
				}
			}
		}
	}
}

// Sell is function for command /sell.
func (b *Bot) Sell() {
	if len(b.Dlg.Text) > 24 {
		// checkvalidbitcoin
		CoinToSell[b.Dlg.UserId] = "MNT"
		depos, err := b.Api.GetMinterDeposAddress(b.Dlg.Text, CoinToSell[b.Dlg.UserId], PriceToSell[b.Dlg.UserId])
		if err != nil {
			msg := tgbotapi.NewMessage(b.Dlg.ChatId, err.Error())
			b.Bot.Send(msg)
			return
		}

		ans := fmt.Sprintf(vocab.GetTranslate("Minter deposit and tag", b.Dlg.language), depos.Data.Address, depos.Data.Tag)
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, ans)
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: false,
			Selective:  false,
		}
		b.Dlg.Command = ""
		b.Bot.Send(msg)
		go b.CheckStatusSell(depos.Data.Tag)
		return

	} else {
		CoinToSell[b.Dlg.UserId] = b.Dlg.Text
		msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Send BTC", b.Dlg.language))
		msg.ReplyMarkup = tgbotapi.ForceReply{
			ForceReply: true,
			Selective:  true,
		}
		b.Bot.Send(msg)
		return
	}
}

// CheckStatusSell checks status of deposit for method Sell().
func (b *Bot) CheckStatusSell(tag string) {
	timeout := time.After(2 * time.Minute)
	tick := time.Tick(3 * time.Second)
	amount := "0"
	for {
		select {
		case <-timeout:
			if amount == "0" {
				msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("timeout", b.Dlg.language))
				msg.ReplyMarkup = b.newMainMenuKeyboard()
				b.Bot.Send(msg)
				return
			} else {
				continue
			}
		case <-tick:
			taginfo, err := b.Api.GetTagInfo(tag)
			if err != nil {
				fmt.Println(err)
				msg := tgbotapi.NewMessage(b.Dlg.ChatId, vocab.GetTranslate("Error", b.Dlg.language))
				b.Bot.Send(msg)
				return
			}
			if taginfo.Data.Amount != amount {
				amount = taginfo.Data.Amount
				// Put in DB.
				b.DB.PutLoot(b.Dlg.UserId, tag, taginfo)
				ans := fmt.Sprintf(vocab.GetTranslate("New deposit for sale", b.Dlg.language), taginfo.Data.Amount, taginfo.Data.Coin, taginfo.Data.Price)
				msg := tgbotapi.NewMessage(b.Dlg.ChatId, ans)
				b.Bot.Send(msg)
				//go a.CheckLootforSell(taginfo.Data.MinterAddress)
				return
			}

		}
	}
}

// Method for sending loots in markdown style to user.
func (b *Bot) ComposeResp(loots []*stct.Loot) {
	for _, loot := range loots {
		text := fmt.Sprintf(
			"*Tag:*  %s\n"+
				"*Coin:*  %s  "+
				"   *Price:*  %v\n"+
				"*Amount:*  %s\n"+
				"*Minted address:*  %s\n"+
				"*Created at:*  %s\n"+
				"*Last sell at:*  %s",
			loot.Tag,
			loot.Coin,
			loot.Price,
			loot.Amout,
			loot.MinterAddress,
			loot.CreatedAt.Format("2006-01-02 15:04:05"),
			loot.LastSell.Format("2006-01-02 15:04:05"))

		msg := tgbotapi.NewMessage(b.Dlg.ChatId, text)
		msg.ParseMode = "markdown"
		b.Bot.Send(msg)
	}
}

// newMainMenuKeyboard is main menu keyboar: price, buy, sell, sales.
func (b *Bot) newMainMenuKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(vocab.GetTranslate("Price", b.Dlg.language), priceCommand),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(vocab.GetTranslate("Buy", b.Dlg.language), buyCommand),
			tgbotapi.NewInlineKeyboardButtonData(vocab.GetTranslate("Sell", b.Dlg.language), sellCommand),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(vocab.GetTranslate("Loots", b.Dlg.language), salesCommand),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(vocab.GetTranslate("Settings", b.Dlg.language), settingsMenu),
		),
	)
}

// vocabuageKeybord is keybouad for select vocabuage.
func (b *Bot) newVocabuageKeybord() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Русский", rusvocabCommand),
			tgbotapi.NewInlineKeyboardButtonData("🇬🇧 English", engvocabCommand),
		),
	)
}

// // newMainKeyboard is keyboard for main menu.
// func (b *Bot) newMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
// 	return tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton(vocab.GetTranslate("Price", b.Dlg.language)),
// 		),
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton(vocab.GetTranslate("Price", b.Dlg.language)),
// 		),
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton(vocab.GetTranslate("Price", b.Dlg.language)),
// 		),
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton(vocab.GetTranslate("Price", b.Dlg.language)),
// 		),
// 	)
// }

//
// func (b *Bot) AddressKeyboardHelp() tgbotapi.ReplyKeyboardMarkup {
// 	keyboard := tgbotapi.ReplyKeyboardMarkup{}
// 	addresses := b.DB.GetAddresses(Minter / BTC)
// 	for _, addr := range addresses {
// 		var row []tgbotapi.KeyboardButton
// 		btn := tgbotapi.NewKeyboardButton(addr)
// 		row = append(row, btn)
// 		keyboard.Keyboard = append(keyboard.Keyboard, row)
// 	}
// 	return keyboard
// }

// func (a *App) CheckLootforSell(addr string) {
// 	tick := time.Tick(1 * time.Hour)
// 	lenght := 0
// 	for {
// 		select {
// 		case <-tick:
// 			history, err := a.MinterAddressHistory(addr)
// 			if err != nil {
// 				log.Fatal(err)
// 				return
// 			}
// 			if len(history.Data) > lenght {

// 			}

// 		}
// 	}
// }
