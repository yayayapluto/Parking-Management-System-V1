package main

import (
	"fmt"
	"github.com/pterm/pterm"
	"parking-management-system-v1/internal/app"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"

	"github.com/AlecAivazis/survey/v2"
)

func vehicleTypeLogic(c *app.Container) {
	for {
		clearScreen()

		var result string
		prompt := &survey.Select{
			Message: "Select Action:",
			Options: []string{"Get All (Paginated)", "Create New", "Get By ID (Hashed)", "Back"},
		}

		err := survey.AskOne(prompt, &result)
		if err != nil {
			return
		}

		switch result {
		case "Get All (Paginated)":
			res, err := c.VehicleTypeService.GetAll(ctx, requests.PaginationRequest{Page: 1, Limit: 10})
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				renderVehicleTable(res.Data.([]responses.VehicleTypeResponse))
			}
			pause()

		case "Create New":
			answers := struct {
				Code        string
				Name        string
				Description string
			}{}

			var qs = []*survey.Question{
				{Name: "code", Prompt: &survey.Input{Message: "Code:"}, Validate: survey.Required},
				{Name: "name", Prompt: &survey.Input{Message: "Name:"}, Validate: survey.Required},
				{Name: "description", Prompt: &survey.Input{Message: "Description:"}},
			}

			err := survey.Ask(qs, &answers)
			if err != nil {
				break
			}

			res, err := c.VehicleTypeService.Create(ctx, requests.CreateVehicleTypeRequest{
				Code:        answers.Code,
				Name:        answers.Name,
				Description: answers.Description,
			})

			if err != nil {
				fmt.Printf("\nGagal: %v\n", err)
			} else {
				fmt.Printf("\nSukses Simpan! ID: %s\n", res.ID)
			}
			pause()

		case "Get By ID (Hashed)":
			var id string
			prompt := &survey.Input{Message: "Masukkan HashID:"}
			err := survey.AskOne(prompt, &id)
			if err != nil {
				return
			}

			res, err := c.VehicleTypeService.GetByID(ctx, id)
			if err != nil {
				fmt.Printf("\nData tidak ditemukan: %v\n", err)
			} else {
				renderVehicleTable([]responses.VehicleTypeResponse{res})
			}
			pause()

		case "Back":
			return
		}
	}
}

func renderVehicleTable(data []responses.VehicleTypeResponse) {
	clearScreen()

	header := []string{"Hash ID", "Code", "Name", "Description", "Active", "Created At"}

	var tableData [][]string
	tableData = append(tableData, header)

	for _, v := range data {
		active := "Yes"
		if !v.IsActive {
			active = "No"
		}

		tableData = append(tableData, []string{
			v.ID,
			v.Code,
			v.Name,
			v.Description,
			active,
			v.CreatedAt,
		})
	}

	err := pterm.DefaultTable.
		WithHasHeader().
		WithBoxed().
		WithData(tableData).
		Render()
	if err != nil {
		return
	}
}

func pause() {
	var unused string
	prompt := &survey.Input{
		Message: "Tekan [ENTER] untuk kembali...",
	}
	err := survey.AskOne(prompt, &unused)
	if err != nil {
		return
	}
}
