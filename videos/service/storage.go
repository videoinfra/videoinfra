package storage

import (
	"database/sql"
	"math/rand"
	"time"
	videoApi "videoinfra/videos/service"
)

type VideoDBContext struct {
	VideosDBHandle *sql.DB
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateVideoId(length int) string {
	return StringWithCharset(length, charset)
}

func (context *VideoDBContext) CreateVideo(draftVideo videoApi.Video) error {

	generatedVideoId := GenerateVideoId(32)
	stmt := `INSERT INTO Videos (video_id, title, account_id, video_state, file_path, create_timestamp, update_timestamp) 
    VALUES(?, ?, ?, ?, ?, current_timestamp, current_timestamp)`

	_, err := context.VideosDBHandle.Exec(stmt, generatedVideoId, draftVideo.Title, draftVideo.AccountId, draftVideo.VideoState, draftVideo.Filepath)
	// TODO(): Better error handling and retry for transient errors.
	return err
}

func (context *VideoDBContext) GetVideo(videoId string) (videoApi.Video, error) {
	video := videoApi.Video{}
	err := context.VideosDBHandle.QueryRow("SELECT * from Videos where video_id = ?", videoId).
		Scan(&video.VideoId, &video.Title, &video.AccountId, &video.VideoState, &video.Filepath, &video.CreateTimestamp, &video.UpdateTimestamp)
    return video, err;
}

func (context *VideoDBContext) UpdateVideo(draftVideo videoApi.Video) (error) {
	stmt := `UPDATE Videos set title = $2, video_state = $3, file_path = $4, update_timestamp = current_timestamp where video_id = $1;`
	_, err := context.VideosDBHandle.Exec(stmt, draftVideo.VideoId, draftVideo.Title, draftVideo.VideoState, draftVideo.Filepath);
	return err
}
