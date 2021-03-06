package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/determined-ai/determined/agent/version"
)

type options struct {
	logLevel string
	noColor  bool
}

func newRootCmd() *cobra.Command {
	opts := options{}

	cmd := &cobra.Command{
		Use:     "determined-agent",
		Version: version.Version,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := bindEnv("DET_", cmd); err != nil {
				return err
			}
			level, err := log.ParseLevel(opts.logLevel)
			if err != nil {
				return err
			}
			log.SetLevel(level)
			log.SetFormatter(&log.TextFormatter{
				FullTimestamp: true,
				ForceColors:   true,
				DisableColors: opts.noColor,
			})
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.logLevel, "log-level", "l", "info",
		"set the logging level (can be one of: debug, info, warn, error, or fatal)")
	cmd.PersistentFlags().BoolVar(&opts.noColor, "no-color", false, "disable colored output")

	cmd.AddCommand(newCompletionCmd())
	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newRunCmd())

	return cmd
}
