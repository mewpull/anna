package main

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitControlLogResetVerbosityCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call InitControlLogResetVerbosityCmd")

	newCmd := &cobra.Command{
		Use:   "verbosity",
		Short: "Make Anna reset log verbosity.",
		Long:  "Make Anna reset log verbosity.",
		Run:   a.ExecControlLogResetVerbosityCmd,
	}

	return newCmd
}

func (a *annactl) ExecControlLogResetVerbosityCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call ExecControlLogResetVerbosityCmd")

	if len(args) > 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	ctx := context.Background()

	err := a.LogControl.ResetVerbosity(ctx)
	if err != nil {
		a.Log.WithTags(spec.Tags{L: "F", O: a, T: nil, V: 1}, "%#v", maskAny(err))
	}
}
