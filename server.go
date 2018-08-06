package main

import (
	stdContext "context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gmmshops.go/app"
	"gmmshops.go/app/configs"
	"gmmshops.go/app/providers/logs"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logs.Logger().Fatal(".env file not found")
	}

}

func main() {

	// helper.ShowLogo("logo.png", -1, 35)

	app := app.NewApplication("API-App", "Choosak", configs.Configure)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX or Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill  is equivalent with the syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)

		select {
		case <-ch:
			logs.Logger().Info("Server: Shutdown in 5 sec.")

			timeout := 5 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}()

	app.Listen(os.Getenv("APP_ADDR"))

}
