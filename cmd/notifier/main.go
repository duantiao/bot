package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/duantiao/bot/pkg/lark"
)

var (
	buildBranch string
	buildCommit string
	buildTime   string

	Version = fmt.Sprintf("Branch [%s] Commit [%s] Build Time [%s]", buildBranch, buildCommit, buildTime)
)

// App
type App struct {
	ctx        context.Context
	cancel     context.CancelFunc
	exitChan   chan os.Signal
	closeQueue []io.Closer
}

func NewApp(exitChan chan os.Signal) *App {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-ctx.Done():
			exitChan <- os.Interrupt
		case <-exitChan:
			cancel()
		}
		close(exitChan)
	}()

	return &App{
		ctx:      ctx,
		cancel:   cancel,
		exitChan: exitChan,
	}
}

func (app *App) Context() context.Context {
	return app.ctx
}

func (app *App) Run(args []string) {
	var cliApp = cli.NewApp()
	cliApp.Name = "Notifier"
	cliApp.Usage = "notice message"
	cliApp.Version = Version
	cliApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "lark-robot-access-token",
			Usage: "lark custom robot access token",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "lark-robot-secret",
			Usage: "lark custom robot secret",
			Value: "",
		},
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "Enable debug",
			EnvVars: []string{"DEBUG"},
		},
		&cli.BoolFlag{
			Name:   "dry-run",
			Hidden: true,
			Usage:  "Notifier dry run",
		},
	}
	cliApp.Action = func(c *cli.Context) error {
		logger:=log.Default()
		if c.Bool("debug") {
			logger =  log.New(os.Stdout, "Notifier", log.LstdFlags|log.Lshortfile)
			os.Setenv("DEBUG", "1")
		}

		// new db manager
		var (
			larkRobotClient        *lark.Client
		)

		// new lark client
		if accessToken, secret := c.String("lark-robot-access-token"), c.String("lark-robot-secret"); accessToken != "" && secret != "" {
			larkRobotClient = lark.NewClient(accessToken, secret)
		} else {
			return cli.Exit("invalid lark robot client args", 1)
		}

		// new notifier
		notifier, err := NewNotifier(logger, larkRobotClient)
		if err != nil {
			return cli.Exit(err, 1)
		}

		// dry run
		if c.Bool("dry-run") {
			notifier.NoticeWorkflow(app.ctx)
		}

		go func() {
			now := time.Now()
			todayAtZero, _ := time.ParseInLocation("2006-01-02", now.Format("2006-01-02"), time.FixedZone("Asia/Shanghai", 8*3600))
			interval :=todayAtZero.Add(time.Hour * 16).Sub(now)
			if interval <= 0 {
				interval =todayAtZero.Add(time.Hour * (16+24)).Sub(now)
			}
			timerForProblematicOnlineCrawler := time.NewTimer(interval)

			for {
				select {
				case <-app.ctx.Done():
					return
				case <-timerForProblematicOnlineCrawler.C:
					timerForProblematicOnlineCrawler.Reset(time.Hour * 24)
					func() {
						ctx, cancel := context.WithTimeout(app.ctx, time.Minute*3)
						defer cancel()
						notifier.NoticeWorkflow(ctx)
					}()
				}
			}
		}()

		<-app.exitChan
		app.cancel()
		return nil
	}
	cliApp.Run(args)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	NewApp(c).Run(os.Args)
}
