package service

import (
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"math/rand"
	"time"
)

type VideoDBContext struct {
	VideosDBHandle *sql.DB
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateVideoId() string {
	return StringWithCharset(48, charset)
}

func GeneratePlaybackId() string {
	return StringWithCharset(48, charset)
}

func (context *VideoDBContext) CreatePlayback(draftPlayback Playback) error {
	stmt := "INSERT INTO playbacks (playback_id, video_id, playback_policy, create_timestamp)" +
		" VALUES ($1, $2, $3, current_timestamp);"

	_, err := context.VideosDBHandle.Exec(stmt, draftPlayback.PlaybackId, draftPlayback.VideoId, draftPlayback.Policy.String())
	return err
}

func (context *VideoDBContext) CreateVideo(draftVideo Video) (string, error) {

	generatedVideoId := GenerateVideoId()
	stmt := "INSERT INTO videos (video_id, title, account_id, video_state, max_width, max_height, max_frame_rate, " +
		"duration_milliseconds, filepath, create_timestamp, update_timestamp) VALUES " +
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, current_timestamp, current_timestamp);"

	fmt.Println("Starting Video creation")
	_, err := context.VideosDBHandle.Exec(stmt, generatedVideoId,
		draftVideo.Title, draftVideo.AccountId, draftVideo.VideoState.String(), draftVideo.VideoProperties.MaxWidth,
		draftVideo.VideoProperties.MaxHeight, draftVideo.VideoProperties.MaxFrameRate, draftVideo.VideoProperties.DurationMs,
		draftVideo.Filepath)

	if err != nil {
		fmt.Printf("Video creation failed: %s", err.Error())
	}
	// TODO(): Better error handling and retry for transient errors.
	return generatedVideoId, err
}

func (context *VideoDBContext) GetVideo(videoId string) (Video, error) {
	video := Video{VideoProperties: &VideoProperties{}}
	var videoStateStr string
	var createTime time.Time
	var updateTime time.Time
	err := context.VideosDBHandle.QueryRow("SELECT * from Videos where video_id = $1;", videoId).
		Scan(&video.VideoId, &video.Title, &video.AccountId, &videoStateStr,
			&video.VideoProperties.MaxWidth, &video.VideoProperties.MaxHeight,
			&video.VideoProperties.MaxFrameRate, &video.VideoProperties.DurationMs,
			&video.Filepath, &createTime, &updateTime)
	if videoStateStr == "PROCESSING" {
		video.VideoState = VideoState_PROCESSING
	} else {
		video.VideoState = VideoState_READY
	}
	video.CreateTimestamp, _ = ptypes.TimestampProto(createTime)
    video.UpdateTimestamp, _ = ptypes.TimestampProto(updateTime)
	return video, err;
}

func (context *VideoDBContext) GetVideosByAccount(accountId int64) ([]Video, error) {
	var videos []Video
	rows, err := context.VideosDBHandle.Query("SELECT * from Videos where account_id = $1 limit 200", accountId)
	if err != nil {
		return videos, err
	}
	defer rows.Close()
	fmt.Println("dB reader")
	for rows.Next() {
		video := Video{VideoProperties: &VideoProperties{}}
		var videoStateStr string
		var createTime time.Time
		var updateTime time.Time
		err = rows.Scan(&video.VideoId, &video.Title, &video.AccountId, &videoStateStr,
			&video.VideoProperties.MaxWidth, &video.VideoProperties.MaxHeight,
			&video.VideoProperties.MaxFrameRate, &video.VideoProperties.DurationMs,
			&video.Filepath, &createTime, &updateTime)

		if videoStateStr == "PROCESSING" {
			video.VideoState = VideoState_PROCESSING
		} else {
			video.VideoState = VideoState_READY
		}
		video.CreateTimestamp, _ = ptypes.TimestampProto(createTime)
		video.UpdateTimestamp, _ = ptypes.TimestampProto(updateTime)
		if err == nil {
			fmt.Println(video.Title)
			videos = append(videos, video)
		}
	}
	return videos, nil
}

func (context *VideoDBContext) UpdateVideo(draftVideo Video) error {
	stmt := "UPDATE Videos set title = $2, video_state = $3, file_path = $4, update_timestamp = current_timestamp where video_id = $1;"
	_, err := context.VideosDBHandle.Exec(stmt, draftVideo.VideoId, draftVideo.Title, draftVideo.VideoState.String(), draftVideo.Filepath);
	return err
}
