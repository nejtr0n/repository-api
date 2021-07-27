package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	application := &cli.App{
		Name:      "gitlab api service",
		Compiled:  time.Now(),
		Copyright: "nejtr0n",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "http_port",
				Value:   ":8000",
				Usage:   "http port to bind",
				EnvVars: []string{"APP_HTTP_PORT"},
			},
			&cli.StringFlag{
				Name:     "gitlab_token",
				Required: true,
				Usage:    "gitlab access token",
				EnvVars:  []string{"APP_GITLAB_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "gitlab_base_url",
				Value:   "https://localhost",
				Usage:   "gitlab base url",
				EnvVars: []string{"APP_GITLAB_BASE_URL"},
			},
		},
		Action: func(config *cli.Context) error {
			router, err := InitRouter(config)
			if err != nil {
				return err
			}

			srv := &http.Server{
				Addr:    config.String("http_port"),
				Handler: router,
			}

			go func() {
				log.Info("Starting http server")
				serverErr := srv.ListenAndServe()
				if serverErr != nil && serverErr != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", serverErr)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Info("Shutting down server...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err = srv.Shutdown(ctx)
			if err != nil {
				return err
			}

			log.Info("Server exiting")

			return nil
		},
	}

	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
