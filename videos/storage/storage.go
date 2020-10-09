package storage

import (
	"gorm.io/gorm"
	"time"
)

type Video struct { 
	VideoId   string
	Title     string
	AccountId string
	VideoState string
	Filepath  string
	CreateTimestamp time.Time
	UpdateTimestamp time.Time
}

type VideoDBContext struct {
	Videos_db_handle *gorm.DB
}

type VideoDraft struct {
	VideoId    string
	Title      string
	AccountId  string
	VideoState string
	Filepath   string
}

func (context *VideoDBContext) CreateVideo(draft VideoDraft) error {
	time_now := time.Now()
	video := Video{VideoId: draft.VideoId, Title: draft.Title, AccountId: draft.AccountId,
		VideoState: draft.VideoState, Filepath: draft.Filepath, CreateTimestamp: time_now, UpdateTimestamp: time_now}
	result := context.Videos_db_handle.Create(&video)
	return result.Error
}

