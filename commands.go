package main

import (
	"errors"
	"github.com/imbaggaarm/fb-stranger-bot/model"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	kCommandStart      string = "#start"
	kCommandEnd        string = "#end"
	kCommandSetGender  string = "#set gender"
	kCommandReport     string = "#report"
	kCommandHelp       string = "#help"
	kCommandGetStarted string = "#get started"
)

func commandHelp(user *model.User) {
	// Send text message
	text := "#start - để bắt đầu chat với người lạ" +
		"\n\n#end - khi bạn muốn end chat." +
		"\n\n#safe mode on|off - để bật/tắt safe mode." +
		" Khi safe mode on, người lạ sẽ không thể gửi ảnh" +
		", audio cho bạn, và ngược lại. Mặc định, safe mode sẽ được bật." +
		"\n\n#report - tố cáo hành vi xấu hoặc lỗi của bot." +
		"\n\n#help - để xem các câu lệnh."
	sendTextMessage(user.ChatID, text)
}

func commandStartReport(user *model.User) {
	// Send generic message
	sendStartReportMessage(user.ChatID)
}

func commandReport(user *model.User, command string) {
	// Set report
	reportCase, err := getCommandValue(command)
	if err != nil {
		sendErrorMessage(user.ChatID)
		return
	}
	if reportCase == "bug" {
		model.CreateReport(1, user.ID, reportCase)
		sendTitleMessage(user.ChatID, "Đã report lỗi", "Cảm ơn bạn nhiều ạ ^^.")
		return
	}
	if user.MatchChatID != nil {
		createReport(*user.MatchChatID, user, reportCase)
	} else if user.PreviousMatch != nil {
		createReport(*user.PreviousMatch, user, reportCase)
	} else {
		sendTitleMessage(user.ChatID, "Không thể report", "Vì bạn chưa nói chuyện với ai cả.")
	}
}

func createReport(reportedChatID string, reporter *model.User, reportCase string) {
	reported, err := model.RetrieveUser(reportedChatID)
	if err != nil {
		sendErrorMessage(reporter.ChatID)
		return
	}
	model.CreateReport(reported.ID, reporter.ID, reportCase)
	sendTitleMessage(reporter.ChatID, "Đã report Cá", "Ban quản trị sẽ xem xét report trong thời gian sớm nhất, cảm ơn bạn nhiều.")
}

func commandStartSetGender(user *model.User) {
	// Send quick replies
	sendSetGenderQuickReplies(user.ChatID, "Bạn hãy chọn giới tính của mình dưới đây nhé:")
}

func commandSetGender(user *model.User, command string) {
	// Set gender to db
	gender, err := getCommandValue(command)
	if err != nil {
		sendErrorMessage(user.ChatID)
		return
	}
	// Send result
	dbGender, err := getDBGenderValue(gender)
	if err != nil {
		//send error message
		sendErrorMessage(user.ChatID)
		return
	}
	if !user.UpdateGender(dbGender) {
		sendErrorMessage(user.ChatID)
	} else {
		genderStr, _ := getGenderStringFrom(dbGender)
		sendDidUpdateGender(user.ChatID, genderStr)
	}
}

func commandStartChat(user *model.User) {
	// TODO: Alert you have to end first
	// Send quick replies
	sendStartChatQuickReplies(user.ChatID, "Bạn muốn chat với ai nè:")
}

func commandMatchChat(user *model.User, command string) {
	// TODO: Alert you have to end first
	if user.Gender == nil {
		sendSetGenderQuickReplies(user.ChatID, "Bạn phải set giới tính của mình trước:")
		return
	}
	// Start matching
	gender, err := getCommandValue(command)
	if err != nil {
		sendErrorMessage(user.ChatID)
		return
	}

	dbGender, err := getDBGenderValue(gender)
	if err != nil {
		sendErrorMessage(user.ChatID)
		return
	}
	if user.Available && user.MatchChatID != nil && *user.GenderNeedMatch == dbGender {
		sendTitleMessage(user.ChatID, "Đang thả thính...", "Cá đang bơi lội tung tăng chưa về, bạn ráng chờ xíu nghen.")
		return
	}
	user.UpdateMatching(dbGender)
	//get available user
	matchUser := user.GetAvailableUser(dbGender)
	if matchUser != nil {
		user.Match(matchUser)
		sendTitleMessage(user.ChatID, "Cá đã ăn thính", "Hãy chăm sóc Cá cẩn thận nhé ^^.")
		sendTitleMessage(matchUser.ChatID, "Cá đã ăn thính", "Hãy chăm sóc Cá cẩn thận nhé ^^.")
		sendTextMessage(matchUser.ChatID, getRandomGreeting())
	} else {
		sendTitleMessage(user.ChatID, "Đang thả thính...", "Cá đang bơi lội tung tăng chưa về, bạn ráng chờ xíu nghen.")
	}
}

func getRandomGreeting() string {
	greetings := []string{"Hi", "Hi", "Hello", "Hiiii", "Helloooo", "Chào cá!", "Chào cậu", "Chào người lạ", "Chào đằng ấy nhé", "Hi cá", "Hello cá",
		"Hallo cá", "Chào bạn", "Chào đằng ấy", "Chào c", "Hi c", "Hello c"}
	rand.Seed(time.Now().Unix())
	return greetings[rand.Intn(len(greetings))]
}

func commandEndChat(user *model.User) {
	matchChatID := user.MatchChatID
	if matchChatID == nil {
		sendTapToStart(user.ChatID, "Bạn đang không ở trong cuộc trò chuyện nào.")
		return
	}
	// End chat
	user.EndConversation()
	// Send result
	sendChatEnded(user.ChatID, "Cuộc trò chuyện đã kết thúc.")
	sendChatEnded(*matchChatID, "Cá đã bơi đi :((")
}

func commandSetSafeMode(user *model.User, command string) {
	// Set safe mode
	value, err := getCommandValue(command)
	if err != nil {
		sendErrorMessage(user.ChatID)
		return
	}
	isOn := value == "on"
	// Send result
	if user.UpdateSafeMode(isOn) {
		var title, subtitle string
		if isOn {
			title = "Safe mode đã được bật."
			subtitle = "Bạn sẽ được an toàn hơn nhiều đó!"
		} else {
			title = "Safe mode đã được tắt."
			subtitle = "Giờ đây bạn có thể nhận được ảnh, audio từ người lạ."
		}
		sendTitleMessage(user.ChatID, title, subtitle)
	} else {
		//Failed
		sendErrorMessage(user.ChatID)
	}
}

func commandGetStarted(user *model.User) {
	// Send welcome message
	if user.Gender == nil {
		sendTitleMessage(user.ChatID, "Chào mừng bạn đến với VNU Chatbot", "Dòng sông cá dành riêng cho sinh viên ^^")
		sendSetGenderQuickReplies(user.ChatID, "Đầu tiên, bạn phải set giới tính của mình:")
	} else {
		// TODO: Alert if user still in a conversation
		sendWelcomeBackMessage(user.ChatID)
	}
}

func getCommandValue(command string) (string, error) {
	commands := strings.Split(command, " ")
	if len(commands) < 2 {
		log.Println("unexpected error has occurred")
		return "", errors.New("unexpected error has occurred")
	}
	value := commands[len(commands)-1]
	return value, nil
}

func getDBGenderValue(strGender string) (uint, error) {
	switch strGender {
	case "male":
		return 0, nil
	case "female":
		return 1, nil
	case "les":
		return 2, nil
	case "gay":
		return 3, nil
	case "all":
		return 4, nil
	default:
		return 0, errors.New("gender not found")
	}
}

func getGenderStringFrom(intGender uint) (string, error) {
	switch intGender {
	case 0:
		return "NAM", nil
	case 1:
		return "NỮ", nil
	case 2:
		return "LES", nil
	case 3:
		return "GAY", nil
	default:
		return "", errors.New("gender not found")
	}
}
