package ui

import (
	"fmt"
	"lyra/globals"
	"lyra/internal"
	"lyra/types"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

func GetModulesFromUi(project types.Project) []string {
	// Create a lookup map for available modules
	moduleMap := make(map[string]types.Module)
	for _, mod := range globals.AvailableModules {
		moduleMap[mod.Name] = mod
	}

	var modules []string
	for _, moduleName := range project.Modules {
		switch moduleName {
		case "Iot Agent - Json":
			if _, exists := moduleMap["fiware-iot-agent"]; exists {
				modules = append(modules, "fiware-iot-agent")
			}
		case "Quantum Leap":
			if _, exists := moduleMap["fiware-quantumleap"]; exists {
				modules = append(modules, "fiware-quantumleap")
			}
		}
	}
	return modules
}

func RunInteractive() {
	project := types.Project{}

	// Should we run in accessible mode?
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Value(&project.Name),
			huh.NewMultiSelect[string]().
				Title("Fiware Modules").
				Description("What Fiware modules do you want to include").
				Options(huh.NewOptions("IoT Agent - Json", "Quantum Leap")...).
				Value(&project.Modules),
			huh.NewSelect[string]().
				Title("Language").
				Description("What language are you going to use").
				Options(huh.NewOptions("C", "CPP")...).
				Value(&project.Language),
			huh.NewSelect[string]().
				Title("License").
				Description("What license are you going to use").
				Options(huh.NewOptions("MIT", "Apache", "GPLv3", "None")...).
				Value(&project.License),
			huh.NewConfirm().
				Title("Git version control").
				Description("Do you want to create a git repository").
				Value(&project.Git),
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
		keyword := func(s any) string {
			var str string
			switch v := s.(type) {
			case string:
				str = v
			case []string:
				str = strings.Join(v, ", ")
			default:
				str = fmt.Sprintf("%v", s)
			}
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(str)
		}
		fmt.Fprintf(&sb,
			"%s\nModules:%s\nName:%s\nWritten in: %s\nLicense:%s\nGit:%s",
			lipgloss.NewStyle().Bold(true).Render("Project"),
			keyword(project.Modules),
			keyword(project.Name),
			keyword(project.Language),
			keyword(project.License),
			keyword(project.Git),
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
	// fmt.Println("From Ui:", GetModulesFromUi(project))
	project.Modules = GetModulesFromUi(project)
	project.License = strings.ToLower(project.License)
	internal.ScaffoldProject(project)
}
