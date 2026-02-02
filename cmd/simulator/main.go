package main

import (
	"context"
	"fmt"
	"os"
	"parking-management-system-v1/internal/app"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/logger"

	"github.com/AlecAivazis/survey/v2"
)

var ctx = context.Background()

func main() {
	traceID := logger.GenerateTraceID()
	log := logger.WithTraceID(traceID)
	log.Info("Application Starting", "main")

	clearScreen()

	cfg, _ := config.LoadConfig()
	db := app.InitDatabase(cfg)
	defer db.Close()

	container := app.NewContainer(db.DB, cfg)

	clearScreen()

	for {
		var result string
		prompt := &survey.Select{
			Message: "Main Menu Group:",
			Options: []string{"Master Data", "Transaction (Coming Soon)", "System Info", "Exit"},
		}

		err := survey.AskOne(prompt, &result)
		if err != nil {
			return
		}

		switch result {
		case "Master Data":
			masterMenu(container)
		case "Transaction (Coming Soon)":
			fmt.Println("Sabar, logic-nya belum lo bikin!")
		case "System Info":
			fmt.Printf("DB: %s | App: %s\n", cfg.Postgres.Database, "Parking v1-Simulator")
			pause()
		case "Exit":
			os.Exit(0)
		}
		clearScreen()
	}
}

func masterMenu(c *app.Container) {
	for {
		clearScreen()
		var result string
		prompt := &survey.Select{
			Message: "Select Master Service:",
			Options: []string{"Vehicle Type", "Customer Regist Source", "Back to Main"},
		}

		err := survey.AskOne(prompt, &result)
		if err != nil {
			return
		}

		switch result {
		case "Vehicle Type":
			vehicleTypeLogic(c)
		case "Customer Regist Source":
			fmt.Println("Tinggal lo contek pola Vehicle Type ya!")
			pause()
		case "Back to Main":
			return
		}
	}
}
