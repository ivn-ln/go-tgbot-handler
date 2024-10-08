package TGBotHandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Checks if the update message contains a bot command and if so, calls the corresponding registered callback.
func (handler Handler) handleCommand(update tgbotapi.Update) {
	if update.Message != nil {
		entities := update.Message.Entities
		for _, entity := range entities {
			switch entity.Type {
			case "bot_command":
				for command, callback := range handler.commands {
					received_command := update.Message.Command()
					if received_command == command {
						callback(update)
					}
				}
			}
		}
	}
}

// Adds a new command handler to the list of commands that are listened for in handleCommand.
// The first parameter is the command name, and the second is the callback that is called when the command is received.
func (handler Handler) RegisterCommand(command string, callback func(update tgbotapi.Update)) {
	handler.commands[command] = callback
}

// Starts the bot and listens for updates. It calls the corresponding callbacks for each update type.
func (handler Handler) StartPolling(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.InlineQuery != nil {
			handler.Callbacks.InlineQueryHandler(update)
			continue
		}
		if update.CallbackQuery != nil {
			handler.Callbacks.CallbackQueryHandler(update)
			continue
		}
		if update.Message.Poll != nil {
			handler.Callbacks.PollHandler(update)
			continue
		}
		if update.Message.Sticker != nil {
			handler.Callbacks.StickerHandler(update)
			continue
		}
		if update.Message.Voice != nil {
			handler.Callbacks.VoiceHandler(update)
			continue
		}
		if update.Message.VideoNote != nil {
			handler.Callbacks.VideoNoteHandler(update)
			continue
		}
		if update.Message.Contact != nil {
			handler.Callbacks.ContactHandler(update)
			continue
		}
		if update.Message.Location != nil {
			handler.Callbacks.LocationHandler(update)
			continue
		}
		if update.Message.Game != nil {
			handler.Callbacks.GameHandler(update)
			continue
		}
		if update.Message.Venue != nil {
			handler.Callbacks.VenueHandler(update)
			continue
		}
		if update.Message.Animation != nil {
			handler.Callbacks.AnimationHandler(update)
			continue
		}
		if update.Message.Video != nil {
			handler.Callbacks.VideoHandler(update)
			continue
		}
		if update.Message.Document != nil {
			handler.Callbacks.DocumentHandler(update)
			continue
		}
		if update.Message.Audio != nil {
			handler.Callbacks.AudioHandler(update)
			continue
		}
		if update.Message.Photo != nil {
			handler.Callbacks.PhotoHandler(update)
			continue
		}
		if update.Message.Command() != "" {
			handler.handleCommand(update)
		} else if update.Message.Text != "" {
			handler.Callbacks.PlaintextHandler(update)
		}
	}
}

// A struct for handling Telegram bot updates.
type Handler struct {
	commands  map[string]func(update tgbotapi.Update)
	Callbacks Callbacks
}

func NewHandler() *Handler {
	return &Handler{commands: make(map[string]func(update tgbotapi.Update)), Callbacks: *newCallbacks()}
}

// Creates a new instance of the Telegram bot API and a handler.
// It uses the provided API token and sets the debug flag to the provided value.
func NewTGBot(api_token string, debug bool) (bot *tgbotapi.BotAPI, handler *Handler, err error) {
	bot, err = tgbotapi.NewBotAPI(api_token)
	bot.Debug = debug
	handler = &Handler{commands: make(map[string]func(update tgbotapi.Update)), Callbacks: *newCallbacks()}
	return
}

type Callbacks struct {
	PlaintextHandler     func(update tgbotapi.Update)
	CallbackQueryHandler func(update tgbotapi.Update)
	InlineQueryHandler   func(update tgbotapi.Update)
	PollHandler          func(update tgbotapi.Update)
	StickerHandler       func(update tgbotapi.Update)
	VoiceHandler         func(update tgbotapi.Update)
	VideoNoteHandler     func(update tgbotapi.Update)
	ContactHandler       func(update tgbotapi.Update)
	LocationHandler      func(update tgbotapi.Update)
	GameHandler          func(update tgbotapi.Update)
	VenueHandler         func(update tgbotapi.Update)
	AnimationHandler     func(update tgbotapi.Update)
	VideoHandler         func(update tgbotapi.Update)
	DocumentHandler      func(update tgbotapi.Update)
	AudioHandler         func(update tgbotapi.Update)
	PhotoHandler         func(update tgbotapi.Update)
}

func newCallbacks() *Callbacks {
	return &Callbacks{
		PlaintextHandler:     func(update tgbotapi.Update) {},
		CallbackQueryHandler: func(update tgbotapi.Update) {},
		InlineQueryHandler:   func(update tgbotapi.Update) {},
		PollHandler:          func(update tgbotapi.Update) {},
		StickerHandler:       func(update tgbotapi.Update) {},
		VoiceHandler:         func(update tgbotapi.Update) {},
		VideoNoteHandler:     func(update tgbotapi.Update) {},
		ContactHandler:       func(update tgbotapi.Update) {},
		LocationHandler:      func(update tgbotapi.Update) {},
		GameHandler:          func(update tgbotapi.Update) {},
		VenueHandler:         func(update tgbotapi.Update) {},
		AnimationHandler:     func(update tgbotapi.Update) {},
		VideoHandler:         func(update tgbotapi.Update) {},
		DocumentHandler:      func(update tgbotapi.Update) {},
		AudioHandler:         func(update tgbotapi.Update) {},
		PhotoHandler:         func(update tgbotapi.Update) {},
	}
}
