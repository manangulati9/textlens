package lib

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Args struct {
	Verbosity        string
	ClipboardHandler string
	Help             bool
	Version          bool
	Reset            bool
	CliMode          bool
	BackgroundMode   bool
}

func ParseArgs() *Args {
	args := Args{}

	flag.StringVar(&args.Verbosity, "verbosity", "warnings", "Set level of detail for console output (available: 'warnings', 'debug', 'errors')")
	flag.BoolVar(&args.Help, "help", false, "Get info about command and flags.")
	flag.BoolVar(&args.Version, "version", false, "Print Textlens version and exit.")
	flag.BoolVar(&args.Reset, "reset", false, "Reset all settings to default values.")
	flag.BoolVar(&args.CliMode, "cli-mode", false, "Print text after detection to stdout and exits immediately.")
	flag.BoolVar(&args.BackgroundMode, "background-mode", false, "Start minimized to tray, without capturing.")
	flag.StringVar(&args.ClipboardHandler, "clipboard-handler", "", "Force using specific clipboard handler instead of auto-selecting.")

	flag.Parse()
	return &args
}

func (args *Args) ValidateArgs() {
	if args.Version {
		fmt.Printf("Textlens %s\n", Version)
		os.Exit(0)
	}

	if args.Help {
		flag.Usage()
		os.Exit(0)
	}

	switch args.Verbosity {
	case "debug":
	case "warnings":
	case "errors":
	default:
		fmt.Println(`Error: Invalid value for -verbosity flag. Allowed values: "debug" | "warnings" | "errors"`)
		os.Exit(1)
	}

	if args.Reset {
		// Do something
	}

	if args.CliMode {
		// Do something
	}

	if args.BackgroundMode {
		// Do something
	}
}

func SetWaylandEnvs() {
	envs := []string{"XCURSOR_SIZE=24", "QT_QPA_PLATFORM=wayland"}
	for _, env := range envs {
		parts := strings.Split(env, "=")
		if _, didFind := os.LookupEnv(parts[0]); !didFind {
			os.Setenv(parts[0], parts[1])
		}
	}
}

// TODO:
func SetFlatpakEnvs() {
}

// TODO:
func SetAppImageEnvs() {
}

/*
Wrapper function that takes an unsafe function as an argument.
It captures any panic and returns a tuple of the result and an error.
*/
func RunUnsafeFunc[T any](fn func() T) (result T, err error) {
	defer func() {
		if r := recover(); r != nil {
			LogError.Println("Provided function errored out; recovered")
			err = errors.New("Function panicked")
			var ok bool
			result, ok = r.(T) // Attempt to assert type to T; otherwise, keep result as zero value
			if !ok {
				result = *new(T) // Assign zero value of T if panic value cannot be asserted
			}
		}
	}()
	result = fn()
	return result, nil
}
