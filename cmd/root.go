package cmd

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/ayn2op/discordo/internal/config"
	"github.com/ayn2op/discordo/internal/logger"
	"github.com/ayn2op/discordo/internal/ui/root"
	"github.com/ayn2op/discordo/internal/version"
	"github.com/ayn2op/tview"
	"github.com/diamondburned/arikawa/v3/utils/ws"
	"github.com/gdamore/tcell/v3"
)

func Run() error {
	return RunWithArgs(os.Args[1:], os.Stdout)
}

func RunWithArgs(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("discordo", flag.ContinueOnError)
	flags.SetOutput(stdout)

	configPath := flags.String("config-path", config.DefaultPath(), "path of the configuration file")
	logPath := flags.String("log-path", logger.DefaultPath(), "path of the log file")
	logLevel := flags.String("log-level", "info", "log level")
	showVersion := flags.Bool("version", false, "print version information")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *showVersion {
		_, err := fmt.Fprintln(stdout, version.String())
		return err
	}

	var level slog.Level
	switch *logLevel {
	case "debug":
		ws.EnableRawEvents = true
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	logFile, err := logger.Load(*logPath, level)
	if err != nil {
		return fmt.Errorf("failed to load logger: %w", err)
	}
	defer logFile.Close()

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		return fmt.Errorf("failed to create screen: %w", err)
	}

	if err := screen.Init(); err != nil {
		return fmt.Errorf("failed to init screen: %w", err)
	}

	if cfg.Mouse {
		screen.EnableMouse()
	}
	screen.EnablePaste()

	tview.Styles = tview.Theme{}
	app := tview.NewApplication(tview.WithScreen(screen))
	app.SetRoot(root.NewModel(cfg, app))
	return app.Run()
}
