package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type PenguinConfig struct {
	Name                string
	BodyFont            string
	TitleFont           string
	Border              bool
	BorderRadius        string
	Primary             string
	PrimaryDark         string
	Secondary           string
	SecondaryDark       string
	Surface             string
	SurfaceDark         string
	SurfaceAlt          string
	SurfaceAltDark      string
	OnSurface           string
	OnSurfaceStrong     string
	OnSurfaceDark       string
	OnSurfaceStrongDark string
	OnPrimary           string
	OnPrimaryDark       string
	OnSecondary         string
	OnSecondaryDark     string
	Outline             string
	OutlineDark         string
	OutlineStrong       string
	OutlineStrongDark   string
	Danger              string
	OnDanger            string
	Info                string
	OnInfo              string
	Warning             string
	OnWarning           string
	Success             string
	OnSuccess           string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Penguin UI Share URL: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	shareURL, err := url.Parse(input)
	if err != nil {
		fmt.Printf("The url you have entered is incorrect: %v\n", err.Error())
		return
	}
	queryParams := shareURL.Query()
	configParams := queryParams["config"]
	if len(configParams) == 0 {
		fmt.Println("The url you have entered is incorrect: no config param found")
		return
	}
	configSplit := strings.Split(configParams[0], "_")

	if len(configSplit) != 33 {
		fmt.Println("The url you have entered is incorrect: not enough params")
		return
	}
	border, err := strconv.ParseBool(configSplit[3])
	if err != nil {
		fmt.Println("The border you have entered is incorrect: not true or false")
		return
	}
	penguinUiConfig := PenguinConfig{
		Name:                configSplit[0],
		BodyFont:            configSplit[1],
		TitleFont:           configSplit[2],
		Border:              border,
		BorderRadius:        configSplit[4],
		Primary:             configSplit[5],
		PrimaryDark:         configSplit[6],
		Secondary:           configSplit[7],
		SecondaryDark:       configSplit[8],
		Surface:             configSplit[9],
		SurfaceDark:         configSplit[10],
		SurfaceAlt:          configSplit[11],
		SurfaceAltDark:      configSplit[12],
		OnSurface:           configSplit[13],
		OnSurfaceStrong:     configSplit[14],
		OnSurfaceDark:       configSplit[15],
		OnSurfaceStrongDark: configSplit[16],
		OnPrimary:           configSplit[17],
		OnPrimaryDark:       configSplit[18],
		OnSecondary:         configSplit[19],
		OnSecondaryDark:     configSplit[20],
		Outline:             configSplit[21],
		OutlineDark:         configSplit[22],
		OutlineStrong:       configSplit[23],
		OutlineStrongDark:   configSplit[24],
		Danger:              configSplit[25],
		OnDanger:            configSplit[26],
		Info:                configSplit[27],
		OnInfo:              configSplit[28],
		Warning:             configSplit[29],
		OnWarning:           configSplit[30],
		Success:             configSplit[31],
		OnSuccess:           configSplit[32],
	}

	file, err := os.Create("penguin_config.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "package main")
	fmt.Fprintln(file, "const (")
	fmt.Fprintf(file, "\tName string = \"%s\"\n", penguinUiConfig.Name)
	fmt.Fprintf(file, "\tBodyFont string = \"%s\"\n", penguinUiConfig.BodyFont)
	fmt.Fprintf(file, "\tTitleFont string = \"%s\"\n", penguinUiConfig.TitleFont)
	fmt.Fprintf(file, "\tBorder bool = %t\n", penguinUiConfig.Border)
	fmt.Fprintf(file, "\tBorderRadius string = \"%s\"\n", penguinUiConfig.BorderRadius)
	fmt.Fprintf(file, "\tPrimary string = \"%s\"\n", penguinUiConfig.Primary)
	fmt.Fprintf(file, "\tPrimaryDark string = \"%s\"\n", penguinUiConfig.PrimaryDark)
	fmt.Fprintf(file, "\tSecondary string = \"%s\"\n", penguinUiConfig.Secondary)
	fmt.Fprintf(file, "\tSecondaryDark string = \"%s\"\n", penguinUiConfig.SecondaryDark)
	fmt.Fprintf(file, "\tSurface string = \"%s\"\n", penguinUiConfig.Surface)
	fmt.Fprintf(file, "\tSurfaceDark string = \"%s\"\n", penguinUiConfig.SurfaceDark)
	fmt.Fprintf(file, "\tSurfaceAlt string = \"%s\"\n", penguinUiConfig.SurfaceAlt)
	fmt.Fprintf(file, "\tSurfaceAltDark string = \"%s\"\n", penguinUiConfig.SurfaceAltDark)
	fmt.Fprintf(file, "\tOnSurface string = \"%s\"\n", penguinUiConfig.OnSurface)
	fmt.Fprintf(file, "\tOnSurfaceStrong string = \"%s\"\n", penguinUiConfig.OnSurfaceStrong)
	fmt.Fprintf(file, "\tOnSurfaceDark string = \"%s\"\n", penguinUiConfig.OnSurfaceDark)
	fmt.Fprintf(file, "\tOnSurfaceStrongDark string = \"%s\"\n", penguinUiConfig.OnSurfaceStrongDark)
	fmt.Fprintf(file, "\tOnPrimary string = \"%s\"\n", penguinUiConfig.OnPrimary)
	fmt.Fprintf(file, "\tOnPrimaryDark string = \"%s\"\n", penguinUiConfig.OnPrimaryDark)
	fmt.Fprintf(file, "\tOnSecondary string = \"%s\"\n", penguinUiConfig.OnSecondary)
	fmt.Fprintf(file, "\tOnSecondaryDark string = \"%s\"\n", penguinUiConfig.OnSecondaryDark)
	fmt.Fprintf(file, "\tOutline string = \"%s\"\n", penguinUiConfig.Outline)
	fmt.Fprintf(file, "\tOutlineDark string = \"%s\"\n", penguinUiConfig.OutlineDark)
	fmt.Fprintf(file, "\tOutlineStrong string = \"%s\"\n", penguinUiConfig.OutlineStrong)
	fmt.Fprintf(file, "\tOutlineStrongDark string = \"%s\"\n", penguinUiConfig.OutlineStrongDark)
	fmt.Fprintf(file, "\tDanger string = \"%s\"\n", penguinUiConfig.Danger)
	fmt.Fprintf(file, "\tOnDanger string = \"%s\"\n", penguinUiConfig.OnDanger)
	fmt.Fprintf(file, "\tInfo string = \"%s\"\n", penguinUiConfig.Info)
	fmt.Fprintf(file, "\tOnInfo string = \"%s\"\n", penguinUiConfig.OnInfo)
	fmt.Fprintf(file, "\tWarning string = \"%s\"\n", penguinUiConfig.Warning)
	fmt.Fprintf(file, "\tOnWarning string = \"%s\"\n", penguinUiConfig.OnWarning)
	fmt.Fprintf(file, "\tSuccess string = \"%s\"\n", penguinUiConfig.Success)
	fmt.Fprintf(file, "\tOnSuccess string = \"%s\"\n", penguinUiConfig.OnSuccess)
	fmt.Fprintln(file, ")")

	cmd := exec.Command("go", "fmt", "penguin_config.go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running go fmt:", err)
		return
	}

	fmt.Println("Configuration written to penguin_config.go")
}
