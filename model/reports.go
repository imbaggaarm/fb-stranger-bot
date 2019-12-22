package model

import "time"

type Report struct {
	ID         uint64
	UserID     uint64
	ReporterID uint64
	CreatedAt  time.Time
	Report     string
}

func CreateReport(userID, reporterID uint64, content string) bool {

	report := Report{
		UserID:     userID,
		ReporterID: reporterID,
		CreatedAt:  time.Now(),
		Report:     content,
	}
	GetDB().Create(&report)
	return report.ID >= 1
}
