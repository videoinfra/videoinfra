package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"time"
	"github.com/golang/protobuf/ptypes"
	"github.com/golangcollege/sessions"
	videoApi "videoinfra/videos/service"
)
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session *sessions.Session
	videoApi videoApi.VideoAPIClient
	templateCache map[string]*template.Template
}

type RenderableVideo struct {
	VideoId string
	Title string
	AccountId int64
	VideoState string
	Filepath string
	CreateTimestamp time.Time
	UpdateTimestamp time.Time
}

type templateData struct {
	CurrentYear int
	Flash string
	FormData url.Values
	FormErrors map[string]string
	Video RenderableVideo
	Videos []*RenderableVideo
}

var (
	videoApiServerAddr = flag.String("videoapi_server_addr", "localhost:10000", "The server address in the format of host:port")
)

func videoToRenderableVideo(video videoApi.Video) RenderableVideo {
	createTime, _ := ptypes.Timestamp(video.CreateTimestamp)
	updateTime, _ := ptypes.Timestamp(video.UpdateTimestamp)

	renderableVideo := RenderableVideo{AccountId: video.AccountId, Title: video.Title, VideoId: video.VideoId,
		Filepath: video.Filepath, CreateTimestamp: createTime,  UpdateTimestamp: updateTime}
	return renderableVideo
}



// The serverError helper writes an error message and stack trace to the errorLog, // then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description // to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to // the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func main() {
	addr := flag.String("addr", ":4300", "HTTP Network Address")
    secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret Key")
    flag.Parse()

	// Initialize loggers.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize VideoAPI client.
	var optsVideoApiClient []grpc.DialOption
	optsVideoApiClient = append(optsVideoApiClient, grpc.WithInsecure())
	videoConn, _ := grpc.Dial(*videoApiServerAddr, optsVideoApiClient...)
	defer videoConn.Close()
	videoApiClient := videoApi.NewVideoAPIClient(videoConn)

	templateCache, err := newTemplateCache("./web/ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		session: session,
		videoApi: videoApiClient,
		templateCache: templateCache,
	}

	// Initialize a tls.Config struct to hold the non-default TLS settings we want // the server to use.
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}


	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
		TLSConfig: tlsConfig,
		// Add Idle, Read and Write timeouts to the server.
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Println("Starting server on ", *addr)
	err = srv.ListenAndServeTLS("./web/tls/cert.pem", "./web/tls/key.pem")
	errorLog.Fatal(err)
}
