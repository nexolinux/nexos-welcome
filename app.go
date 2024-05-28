package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var desktopEnv string

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func init() {
	desktopEnv = getDesktopEnvironment()
}

func getDesktopEnvironment() string {
	return os.Getenv("XDG_CURRENT_DESKTOP")
}

func (a *App) CurrentTheme() string {
	var currThemeName string
	switch desktopEnv {
	case "KDE":
		cmd := exec.Command("kreadconfig5", "--file", "kdeglobals", "--group", "General", "--key", "widgetStyle")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		}

		currThemeName = strings.TrimSpace(string(output))
	case "XFCE":
		cmd := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName", "-v")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		}

		currThemeName = strings.TrimSpace(string(output))
	case "GNOME":
		cmd := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "gtk-theme")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		}

		currThemeName = strings.Trim(strings.TrimSpace(string(output)), "'")
	}

	return currThemeName
}

func (a *App) ToggleTheme(dark bool) {
	var style string
	switch desktopEnv {
	case "KDE":
		if dark {
			style = "kde dark theme"
		} else {
			style = "kde light theme"
		}
		cmd := exec.Command("kwriteconfig5", "--file", "kdeglobals", "--group", "KDE", "--key", "LookAndFeelPackage", style)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("failed to change KDE theme:", err)
		}
	case "GNOME":
		if dark {
			style = "prefer-dark"
		} else {
			style = "prefer-light"
		}
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "color-scheme", style)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("failed to change GNOME theme:", err)
		}
	case "XFCE":
		if dark {
			style = "Adwaita-dark"
		} else {
			style = "Adwaita"
		}
		cmd := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName", "-s", style)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("failed to change XFCE theme:", err)
		}
	default:
		fmt.Println("unsupported desktop environment:", desktopEnv)
	}
}

func (a *App) MirrorList(command string) {
	pkexecCmd := exec.Command("pkexec", "sh", "-c", command)

	pkexecCmd.Stdin = os.Stdin
	pkexecCmd.Stdout = os.Stdout
	pkexecCmd.Stderr = os.Stderr

	err := pkexecCmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
