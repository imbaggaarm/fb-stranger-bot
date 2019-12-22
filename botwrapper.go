package main

import "github.com/imbaggaarm/go-messenger"

func sendTextMessage(recipientID, text string) {
	bot.SendTextMessage(recipientID, text)
}

func sendGenericMessage(recipientID string, elements []messenger.Element) {
	bot.SendGenericMessage(recipientID, elements)
}

func sendButtonMessage(recipientID, text string, buttons []messenger.Button) {
	bot.SendButtonMessage(recipientID, text, buttons)
}

func sendTitleMessage(recipientID, title, subtitle string) {
	elements := []messenger.Element{
		{
			Title:    title,
			Subtitle: subtitle,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendQuickReplies(recipientID, text string, quickReplies []messenger.QuickReply) {
	bot.SendQuickReplies(recipientID, text, quickReplies)
}

func sendSetGenderQuickReplies(recipientID, text string) {
	quickReplies := []messenger.QuickReply{
		{
			ContentType: "text",
			Title:       "Nam",
			Payload:     "#gender male",
		},
		{
			ContentType: "text",
			Title:       "Nữ",
			Payload:     "#gender female",
		},
		{
			ContentType: "text",
			Title:       "Les",
			Payload:     "#gender les",
		},
		{
			ContentType: "text",
			Title:       "Gay",
			Payload:     "#gender gay",
		},
	}
	sendQuickReplies(recipientID, text, quickReplies)
}

func sendStartChatQuickReplies(recipientID string, text string) {
	quickReplies := []messenger.QuickReply{
		{
			ContentType: "text",
			Title:       "Nam",
			Payload:     "#start male",
		},
		{
			ContentType: "text",
			Title:       "Nữ",
			Payload:     "#start female",
		},
		{
			ContentType: "text",
			Title:       "Les",
			Payload:     "#start les",
		},
		{
			ContentType: "text",
			Title:       "Gay",
			Payload:     "#start gay",
		},
		{
			ContentType: "text",
			Title:       "Bất kì",
			Payload:     "#start all",
		},
	}
	sendQuickReplies(recipientID, text, quickReplies)
}

func sendChatEnded(recipientID, title string) {
	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "Report cá",
			Payload: "#report",
		},
		{
			Type:    "postback",
			Title:   "Câu cá mới",
			Payload: "#start",
		},
	}

	elements := []messenger.Element{
		{
			Title:    title,
			Subtitle: "Câu cá mới bạn nha...",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendWelcomeBackMessage(recipientID string) {
	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "Hướng dẫn",
			Payload: "#help",
		},
		{
			Type:    "postback",
			Title:   "Câu cá mới",
			Payload: "#start",
		},
	}
	elements := []messenger.Element{
		{
			Title:    "Welcome back ^^",
			Subtitle: "Cùng câu cá với mọi người nhé...",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendTapToStart(recipientID, subtitle string) {
	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "Câu cá mới",
			Payload: "#start",
		},
	}

	elements := []messenger.Element{
		{
			Title:    "Câu cá nào bạn ơiiii",
			Subtitle: subtitle,
			Buttons:  buttons,
		},
	}

	sendGenericMessage(recipientID, elements)
}

func sendOutOfWaitingRoom(recipientID string) {
	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "Câu cá mới",
			Payload: "#start",
		},
	}

	elements := []messenger.Element{
		{
			Title:    "Đã thoát khỏi phòng chờ",
			Subtitle: "Câu cá là phải kiên nhẫn nè :((",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendDidUpdateGender(recipientID, gender string) {
	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "Hướng dẫn sử dụng",
			Payload: "#help",
		},
		{
			Type:    "postback",
			Title:   "Thả thính",
			Payload: "#start",
		},
	}

	elements := []messenger.Element{
		{
			Title:    "Giới tính của bạn là " + gender,
			Subtitle: "Bạn chỉ thay đổi được giới tính của mình sau 30 ngày nữa.",
			Buttons:  buttons,
		},
	}

	sendGenericMessage(recipientID, elements)
}

func sendTurnOffSafeModeMessage(recipientID string) {
	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "Tắt safe mode",
			Payload: "#safe mode off",
		},
	}
	elements := []messenger.Element{
		{
			Title:    "Người lạ vừa gửi bạn một tin nhắn có đính kèm",
			Subtitle: "Nếu bạn muốn nhận tin nhắn này, hãy tắt safe mode nhé.",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendErrorMessage(recipientID string) {
	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "Báo cáo lỗi",
			Payload: "#report bug",
		},
	}
	sendButtonMessage(recipientID, "Đã có lỗi xảy ra. Vui lòng báo cáo lỗi nhé, xin lỗi bạn nhiều.", buttons)
}

func setGetStarted() {
	gsPayload := messenger.Payload{GetStarted: &messenger.GetStarted{Payload: kCommandGetStarted}}
	bot.SetGetStarted(gsPayload)
}

func removeGetStarted() {
	bot.RemoveGetStarted()
}

func setPersistentMenu() {
	// TODO: Build help url
	// TODO: Set button types in library

	buttons := []messenger.Button{
		{
			Type:    "postback",
			Title:   "End chat",
			Payload: "#end",
		},
		{
			Type:    "postback",
			Title:   "Hướng dẫn sử dụng",
			Payload: "#help",
			URL:     "",
		},
		{
			Type:    "postback",
			Title:   "Report",
			Payload: "#report",
		},
	}
	pmPayload := messenger.Payload{PersistentMenu: []messenger.PersistentMenu{
		{
			Locale:                "vn_vi",
			ComposerInputDisabled: false,
			CallToActions:         buttons,
		},
	}}
	bot.SetPersistentMenu(pmPayload)
}

func removePersistentMenu() {
	bot.RemovePersistentMenu()
}
