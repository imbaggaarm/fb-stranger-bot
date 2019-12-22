package main

import (
	"github.com/imbaggaarm/fb-stranger-bot/model"
	. "github.com/imbaggaarm/go-messenger"
	"strings"
)

func handleWebhookEvent(event WebhookEvent) {
	messaging := event.Entry[0].Messaging
	message := (*messaging)[0]

	senderID := message.Sender.ID

	// Get user
	user, err := model.RetrieveUser(senderID)
	if err != nil {
		sendErrorMessage(senderID)
		return
	}

	if message := message.Message; message != nil {
		handleMessage(user, *message)
		return
	}

	if postBack := message.Postback; postBack != nil {
		handlePostBack(user, *postBack)
		return
	}
}

func handleMessage(user *model.User, message WebhookMessage) {
	// Handle commands
	if text := message.Text; text != nil && strings.HasPrefix(*text, "#") {
		handleCommands(user, *text)
		return
	}
	// Handle quick reply
	if quickReply := message.QuickReply; quickReply != nil {
		payload := quickReply.Payload
		handleCommands(user, payload)
		return
	}
	// Handle normal message
	handleNormalMessage(user, message)
}

func handlePostBack(user *model.User, postBack Postback) {
	payload := postBack.Payload
	handleCommands(user, payload)
}

func handleCommands(user *model.User, command string) {
	switch command {
	case "#start":
		commandStartChat(user)
	case "#start male", "#start female", "#start les", "#start gay", "#start all":
		commandMatchChat(user, command)
	case "#end":
		commandEndChat(user)
	case "#set gender":
		commandStartSetGender(user)
	case "#gender male", "#gender female", "#gender les", "#gender gay":
		commandSetGender(user, command)
	case "#help":
		commandHelp(user)
	case "#report":
		commandStartReport(user)
	case "#report bug", "#report 18+", "#report impolite":
		commandReport(user, command)
	case "#safe mode on", "#safe mode off":
		commandSetSafeMode(user, command)
	case "#get started":
		commandGetStarted(user)
	default:
		handleInvalidCommand(user)
	}
}

func handleInvalidCommand(user *model.User) {
	buttons := []Button{
		{
			Type:    "postback",
			Title:   "Hướng dẫn sử dụng",
			Payload: "#help",
		},
	}
	sendButtonMessage(user.ChatID, "Không có câu lệnh nào trùng khớp.\n", buttons)
}

func handleNormalMessage(user *model.User, message WebhookMessage) {
	if user.MatchChatID == nil {
		sendTapToStart(user.ChatID, "Bạn đang không ở trong cuộc trò chuyện nào!")
		return
	}
	//send message between users
	partner, err := user.RetrievePartner()
	if err != nil {
		sendErrorMessage(user.ChatID)
		return
	}
	text := message.Text
	if text != nil {
		sendTextMessage(partner.ChatID, *text)
		return
	}
	attachments := message.Attachments
	if attachments != nil {
		if partner.SafeMode {
			sendTitleMessage(user.ChatID, "Người lạ đang bật safe mode", "Hãy nhắn người lạ tắt safe mode nếu muốn nhận tin nhắn này nhé.")
			sendTurnOffSafeModeMessage(partner.ChatID)
			return
		}
		attachment := (*attachments)[0]
		switch attachment.Type {
		case AttachmentTypeAudio:
			url := attachment.Payload.URL
			bot.SendAudioUrl(partner.ChatID, url)
		case AttachmentTypeImage:
			url := attachment.Payload.URL
			bot.SendImageUrl(partner.ChatID, url)
		default:
			sendTextMessage(user.ChatID, "Bạn chỉ có thể gửi 1 ảnh hoặc audio")
		}
		return
	}
	sendErrorMessage(user.ChatID)
}
