package model

import (
	"errors"
	"time"
)

type User struct {
	ID                 uint64
	ChatID             string
	Gender             *uint
	LastActivity       time.Time
	MatchChatID        *string
	GenderNeedMatch    *uint
	Available          bool
	RegisterDate       time.Time
	PreviousMatch      *string
	SafeMode           bool
	BannedUntil        *time.Time
	WaitingTimestamp   *uint64
	LastGenderChanging *time.Time
}

func RetrieveUser(chatID string) (*User, error) {
	user := &User{ChatID: chatID}
	GetDB().Where("chat_id = ?", chatID).First(user)
	if user.ID == 0 {
		// User hasn't existed, create user
		createUser(user)
	}
	if user.ID <= 0 {
		return nil, errors.New("database error")
	}
	return user, nil
}

func (user *User) RetrievePartner() (*User, error) {
	return RetrieveUser(*user.MatchChatID)
}

func createUser(user *User) bool {
	now := time.Now()
	user.RegisterDate = now
	user.LastActivity = now
	err := GetDB().Save(user)
	return err == nil
}

func (user *User) UpdateMatching(genderNeedMatch uint) bool {
	waitingTimestamp := uint64(time.Now().Unix())
	err := GetDB().Model(user).Updates(map[string]interface{}{"gender_need_match": genderNeedMatch, "available": true, "waiting_timestamp": waitingTimestamp}).Error
	return err == nil
}

func (user *User) UpdateAvailable(available bool) bool {
	user.Available = available
	err := GetDB().Save(user).Error
	return err == nil
}

func (user *User) UpdateSafeMode(isSafeMode bool) bool {
	user.SafeMode = isSafeMode
	err := GetDB().Save(user).Error
	return err == nil
}

func (user *User) UpdateGender(gender uint) bool {
	user.Gender = &gender
	lastGenderChanging := time.Now()
	user.LastGenderChanging = &lastGenderChanging
	err := GetDB().Save(user).Error
	return err == nil
}

func (user *User) GetAvailableUser(genderNeedMatch uint) *User {
	//matchUser := &User{}
	var results []*User
	if genderNeedMatch == 4 {
		GetDB().
			Where("chat_id != ? AND available = 1 AND match_chat_id IS NULL AND (gender_need_match = ? OR gender_need_match = 4)", user.ChatID, user.Gender).
			Order("waiting_timestamp").
			Limit(1).
			Find(&results)
	} else {
		GetDB().
			Where("chat_id != ? AND available = 1 AND match_chat_id IS NULL AND (gender_need_match = ? OR gender_need_match = 4) AND gender = ? ", user.ChatID, user.Gender, user.GenderNeedMatch).
			Order("waiting_timestamp").
			Limit(1).
			Find(&results)
	}
	if len(results) < 1 {
		return nil
	}
	return results[0]
}

func (user *User) Match(matchUser *User) {
	GetDB().Exec("UPDATE users SET match_chat_id = ? WHERE chat_id = ?", matchUser.ChatID, user.ChatID)
	GetDB().Exec("UPDATE users SET match_chat_id = ? WHERE chat_id = ?", user.ChatID, matchUser.ChatID)
}

func (user *User) EndConversation() bool {
	if user.MatchChatID == nil {
		return true
	}
	GetDB().Exec("UPDATE users SET match_chat_id = NULL, available = 0, previous_match = ? WHERE chat_id = ?", user.MatchChatID, user.ChatID)
	GetDB().Exec("UPDATE users SET match_chat_id = NULL, available = 0, previous_match = ? WHERE chat_id = ?", user.ChatID, user.MatchChatID)
	return true
}
