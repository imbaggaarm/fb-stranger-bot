package main

import . "github.com/imbaggaarm/go-messenger"

func sendTextMessage(recipientID, text string) {
	bot.SendTextMessage(recipientID, text)
}

func sendGenericMessage(recipientID string, elements []Element) {
	bot.SendGenericMessage(recipientID, elements)
}

func sendButtonMessage(recipientID, text string, buttons []Button) {
	bot.SendButtonMessage(recipientID, text, buttons)
}

func sendStartReportMessage(recipientID string) {
	buttons := []Button{
		{
			Type:    "postback",
			Title:   "18+",
			Payload: "#report 18+",
		},
		{
			Type:    "postback",
			Title:   "Bất lịch sự",
			Payload: "#report impolite",
		},
		{
			Type:    "postback",
			Title:   "Lỗi chatbot",
			Payload: "#report bug",
		},
	}
	sendButtonMessage(recipientID, "Hãy chọn nội dung dưới đây để báo cáo nhé:", buttons)
}

func sendTitleMessage(recipientID, title, subtitle string) {
	elements := []Element{
		{
			Title:    title,
			Subtitle: subtitle,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendQuickReplies(recipientID, text string, quickReplies []QuickReply) {
	bot.SendQuickReplies(recipientID, text, nil, quickReplies)
}

func sendSetGenderQuickReplies(recipientID, text string) {
	quickReplies := []QuickReply{
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
	quickReplies := []QuickReply{
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
	buttons := []Button{
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

	elements := []Element{
		{
			Title:    title,
			Subtitle: "Câu cá mới bạn nha...",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendWelcomeBackMessage(recipientID string) {
	buttons := []Button{
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
	elements := []Element{
		{
			Title:    "Welcome back ^^",
			Subtitle: "Cùng câu cá với mọi người nhé...",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendTapToStart(recipientID, subtitle string) {
	buttons := []Button{
		{
			Type:    "postback",
			Title:   "Câu cá mới",
			Payload: "#start",
		},
	}

	elements := []Element{
		{
			Title:    "Câu cá nào bạn ơiiii",
			Subtitle: subtitle,
			Buttons:  buttons,
		},
	}

	sendGenericMessage(recipientID, elements)
}

func sendOutOfWaitingRoom(recipientID string) {
	buttons := []Button{
		{
			Type:    "postback",
			Title:   "Câu cá mới",
			Payload: "#start",
		},
	}

	elements := []Element{
		{
			Title:    "Đã thoát khỏi phòng chờ",
			Subtitle: "Câu cá là phải kiên nhẫn nè :((",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendDidUpdateGender(recipientID, gender string) {
	buttons := []Button{
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

	elements := []Element{
		{
			Title:    "Giới tính của bạn là " + gender,
			Subtitle: "Bạn chỉ thay đổi được giới tính của mình sau 30 ngày nữa.",
			Buttons:  buttons,
		},
	}

	sendGenericMessage(recipientID, elements)
}

func sendTurnOffSafeModeMessage(recipientID string) {
	buttons := []Button{
		{
			Type:    "postback",
			Title:   "Tắt safe mode",
			Payload: "#safe mode off",
		},
	}
	elements := []Element{
		{
			Title:    "Người lạ vừa gửi bạn một tin nhắn có đính kèm",
			Subtitle: "Nếu bạn muốn nhận tin nhắn này, hãy tắt safe mode nhé.",
			Buttons:  buttons,
		},
	}
	sendGenericMessage(recipientID, elements)
}

func sendErrorMessage(recipientID string) {
	buttons := []Button{
		{
			Type:    "postback",
			Title:   "Báo cáo lỗi",
			Payload: "#report bug",
		},
	}
	sendButtonMessage(recipientID, "Đã có lỗi xảy ra. Vui lòng báo cáo lỗi nhé, xin lỗi bạn nhiều.", buttons)
}

func setGetStarted() {
	gsPayload := Payload{GetStarted: &GetStarted{Payload: kCommandGetStarted}}
	bot.SetGetStarted(gsPayload)
}

func removeGetStarted() {
	bot.RemoveGetStarted()
}

func setPersistentMenu() {
	// TODO: Build help url
	// TODO: Set button types in library

	buttons := []Button{
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
	pmPayload := Payload{PersistentMenu: []PersistentMenu{
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
