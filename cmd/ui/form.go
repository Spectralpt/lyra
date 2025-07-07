package ui

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

type Project struct {
	Name    string
	Type    string
	Lang    string
	License string
}

func RunInteractive() {
	test := Project{}

	// Should we run in accessible mode?
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Type").
				Description("What type of project do you want to create").
				Options(huh.NewOptions("Fiware Context Broker", "Fiware IoT Agent")...).
				Value(&test.Type),
			huh.NewInput().
				Title("Project Name").
				Value(&test.Name),
			huh.NewSelect[string]().
				Title("Language").
				Description("What language are you going to use").
				Options(huh.NewOptions("C", "CPP")...).
				Value(&test.Lang),
			huh.NewSelect[string]().
				Title("License").
				Description("What license are you going to use").
				Options(huh.NewOptions("MIT", "Apache", "GPLv3")...).
				Value(&test.License),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	prepareBurger := func() {
		time.Sleep(2 * time.Second)
	}

	_ = spinner.New().Title("Creating your project...").Accessible(accessible).Action(prepareBurger).Run()

	// Print Project summary.
	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"%s\nType:%s\nName:%s\nWritten in: %s\nLicense:%s",
			lipgloss.NewStyle().Bold(true).Render("Project"),
			keyword(test.Type),
			keyword(test.Name),
			keyword(test.Lang),
			keyword(test.License),
		)

		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)
	}
}
