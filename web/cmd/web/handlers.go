package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
	videoApi "videoinfra/videos/service"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	app.infoLog.Println("Calling VideoAPI.GetVideosByAccount ...")
	response, err := app.videoApi.GetVideosByAccount(ctx, &videoApi.GetVideosByAccountRequest{AccountId: 1})
    if (err != nil) {
    	app.serverError(w, err)
	}

	var videos []*RenderableVideo
	for _, video := range response.Videos {
		app.infoLog.Println(video.Title)
		renderableVideo := videoToRenderableVideo(*video)
		videos = append(videos, &renderableVideo)
	}

	app.render(w, r, "home.page.tmpl", &templateData{Videos: videos})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	videoId := r.URL.Query().Get(":v")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	app.infoLog.Println("Calling VideoAPI.GetVideo ...")
	response, err := app.videoApi.GetVideo(ctx, &videoApi.GetVideoRequest{VideoId: videoId})
	if err != nil {
		app.errorLog.Println(err)
	}
	app.infoLog.Println("Call completed.")

	createTime, _ := ptypes.Timestamp(response.Video.CreateTimestamp)
	updateTime, _ := ptypes.Timestamp(response.Video.UpdateTimestamp)

	// Convert from video proto (for backend) to renderablevideo struct (for frontend template)
	renderableVideo := RenderableVideo{AccountId: response.Video.AccountId, Title: response.Video.Title, VideoId: response.Video.VideoId,
		Filepath: response.Video.Filepath, CreateTimestamp: createTime,  UpdateTimestamp: updateTime}

	app.render(w, r, "show.page.tmpl", &templateData{Video: renderableVideo})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	inputUrl := r.PostForm.Get("input_url")
	title := r.PostForm.Get("title")

	errors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank."
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(inputUrl) == "" {
		errors["inputUrl"] = "This field cannot be blank."
	}

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{ FormErrors: errors,
			FormData: r.PostForm,
		})
		return

	}

	accountId := int64(1)

	// Create video in storage with state = PROCESSING.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	app.infoLog.Println("Calling VideoAPI.CreateVideo ...")
	createResponse, err := app.videoApi.CreateVideo(ctx, &videoApi.CreateVideoRequest{Title: title, AccountId: accountId})
	if err != nil {
		app.errorLog.Println(err)
	}
	app.infoLog.Println("Call completed.")

	videoId := createResponse.VideoId;

	// Upload the video to filesystem.
	videoPath, err := app.uploadFromUrl(inputUrl)
	if err != nil {
		app.errorLog.Println(err)
	}

	// Create playback.
	app.infoLog.Println("Calling VideoAPI.CreatePlayback ...")
	playbackResponse, err := app.videoApi.CreatePlayback(ctx, &videoApi.CreatePlaybackRequest{VideoId: videoId,
		VideoPath: videoPath, Policy: videoApi.PlaybackPolicy_PUBLIC})
	if err != nil {
		app.errorLog.Println(err)
	}
	app.infoLog.Println("Call completed.")
	app.infoLog.Printf("Created playback with id: %s", playbackResponse.PlaybackId)

	app.session.Put(r, "flash", "Video asset successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/%s", videoId), http.StatusSeeOther)
}


func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{})
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user signup form...")
}
func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
