package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/switchupcb/copygen/cli/generator"
	"github.com/switchupcb/copygen/cli/loader"
)

// Environment represents the copygen environment.
type Environment struct {
	YMLPath string // The .yml file path used as a configuration file.
	Output  bool   // Whether to print the generated code to stdout.
}

// CLI runs the copygen command and returns its exit status.
func CLI(args []string) int {
	var env Environment
	err := env.parseArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 2
	}

	if err = env.run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}
	return 0
}

// parseArgs parses the provided command line arguments.
func (e *Environment) parseArgs(args []string) error {
	// define the command line arguments.
	ymlPtr := flag.String("yml", "", "The path to the .yml flag used for code generation (from the current working directory).")
	output := flag.Bool("o", false, "Use -o to print generated code to the screen.")

	// parse the command line arguments.
	flag.Parse()

	// yml
	ymlLen := len(*ymlPtr)
	if ymlLen == 0 {
		return fmt.Errorf("You must specify a .yml configuration file using -yml.")
	} else if ymlLen < 4 || ".yml" != (*ymlPtr)[ymlLen-4:] {
		return fmt.Errorf("The specified file (-yml) is not a .yml file.")
	}
	e.YMLPath = *ymlPtr

	// output
	e.Output = *output
	return nil
}

func (e *Environment) run() error {
	gen, err := loader.LoadYML(e.YMLPath)
	if err != nil {
		return err
	}

	if err = generator.Generate(gen, e.Output); err != nil {
		return err
	}
	return nil
}
