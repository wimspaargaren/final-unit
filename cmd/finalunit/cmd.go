package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Default constants
const (
	DefaultPopulationSize   = 30
	DefaultTestCasesPerFunc = 18
	// Currently the best way to detect cycles is to count
	// the amount of times some struct is created
	DefaultAmountRecursion = 3
	DefaultNoImprovedGens  = 10
	DefaultTargetFitness   = 0.95
)

func initCmd(globalOpts *Opts) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "finalunit",
		Short: "Final Unit is a tool to generate unit test for Go automatically.",
		Long: `Final Unit is a command line tool to automatically 
	generate unit test cases for your Go source code. 
	It uses evolutionary based machine learning in order to try 
	and create a test suite with the highest coverage it can find. `,
		Version: Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			target, err := cmd.Flags().GetFloat64("target-fitness")
			if err != nil {
				return err
			}

			if target <= 0 || target > 1 {
				return fmt.Errorf("--target-fitness flag must between 0 and 1")
			}
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&globalOpts.Dir, "dir", "d", ".", "Dir for which to execute the generator")
	rootCmd.PersistentFlags().BoolVarP(&globalOpts.Debug, "debug", "D", false, "Run generator in debug mode")
	rootCmd.PersistentFlags().BoolVarP(&globalOpts.Verbose, "verbose", "v", false, "Run generator in verbose mode")
	// gen opts
	rootCmd.Flags().IntVar(&globalOpts.OrganismAmount, "org-amount", DefaultPopulationSize, "Set amount of organisms in the population")
	rootCmd.Flags().IntVar(&globalOpts.TestCasesPerFunc, "test-cases-func", DefaultTestCasesPerFunc, "Set amount of test cases created for every function")
	rootCmd.Flags().IntVar(&globalOpts.MaxRecursion, "max-recursion", DefaultAmountRecursion, "Set the amount of times one struct is created")
	// population opts
	rootCmd.Flags().IntVar(&globalOpts.MaxNoImprovGens, "no-improve-gens", DefaultNoImprovedGens, "Set max amount of generations without improvements before the generator halts ")
	rootCmd.Flags().Float64Var(&globalOpts.Target, "target-fitness", DefaultTargetFitness, "Set number between 0 and 1 indicating the target coverage we try to hit")

	return rootCmd
}
