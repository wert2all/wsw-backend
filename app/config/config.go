package config

import (
	"flag"
	"strconv"

	"github.com/getsentry/sentry-go"
)

const (
	defaultHost     = "localhost"
	defaultHTTPHost = "http://" + defaultHost
)

type Postgres struct {
	Port     int
	Host     string
	DB       string
	User     string
	Password string
}

type Gowitness struct {
	ScreenshotPath    string
	ScreenshotBaseUrl string
}

type (
	Rollbar struct {
		Token string
	}
	Config struct {
		App       AppConfig
		Postgres  Postgres
		Gowitness Gowitness
		Sentry    sentry.ClientOptions
		Rollbar   Rollbar
	}
)

type ListenHost struct {
	Host string
	Port int
}

type AppConfig struct {
	Listen        ListenHost
	AssetsBaseURL string
}

func NewConfig() Config {
	var (
		listenHostFlag string
		listenPortFlag int

		assetsBaseURL string

		dbHostFlag         string
		dbPortFlag         int
		dbNameFlag         string
		dbUserNameFlag     string
		dbUserPasswordFlag string

		screenShotPath    string
		screenshotBaseURL string
	)

	flag.StringVar(&listenHostFlag, "listen-host", "", "Listen host")
	flag.IntVar(&listenPortFlag, "listen-port", 8000, "Listen port")

	flag.StringVar(&assetsBaseURL, "assets-base-url", defaultHTTPHost+":"+strconv.Itoa(listenPortFlag)+"/assets/", "Assets base url")

	flag.StringVar(&dbHostFlag, "db-host", defaultHost, "Database host")
	flag.IntVar(&dbPortFlag, "db-port", 5432, "Database port")
	flag.StringVar(&dbNameFlag, "db-name", "wsw", "Database name")
	flag.StringVar(&dbUserNameFlag, "db-user-name", "wsw", "Database user name")
	flag.StringVar(&dbUserPasswordFlag, "db-user-password", "wsw", "Database user password")

	flag.StringVar(&screenShotPath, "screenshot-path", "data/screenshots", "Screenshot path")
	flag.StringVar(&screenshotBaseURL, "screenshot-base-url", "http://localhost:8000/screenshot/", "Base url for screenshot")

	flag.Parse()

	config := Config{
		App: AppConfig{
			Listen: ListenHost{
				Host: listenHostFlag,
				Port: listenPortFlag,
			},
			AssetsBaseURL: assetsBaseURL,
		},
		Postgres: Postgres{
			Host:     dbHostFlag,
			Port:     dbPortFlag,
			DB:       dbNameFlag,
			User:     dbUserNameFlag,
			Password: dbUserPasswordFlag,
		},
		Gowitness: Gowitness{
			ScreenshotPath:    screenShotPath,
			ScreenshotBaseUrl: screenshotBaseURL,
		},
		Sentry: sentry.ClientOptions{
			Dsn:           "https://563bfbafd64427d650b376395d83765c@o390093.ingest.us.sentry.io/4508046587002880",
			EnableTracing: false,
		},
		Rollbar: Rollbar{
			Token: "b4935a32816e485ca41d70d2ae2884dc",
		},
	}
	// utils.D(config)
	return config
}
