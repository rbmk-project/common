// SPDX-License-Identifier: GPL-3.0-or-later

package cliutils_test

import (
	"context"
	"strings"
	"testing"

	"github.com/rbmk-project/common/cliutils"
)

type fakecmd struct {
	err error
}

var _ cliutils.Command = fakecmd{}

// Help implements [cliutils.Command].
func (f fakecmd) Help(argv ...string) error {
	return nil
}

// Main implements [cliutils.Command].
func (f fakecmd) Main(ctx context.Context, argv ...string) error {
	return f.err
}

func TestCommandWithSubCommands(t *testing.T) {
	type testcase struct {
		argv    []string
		failure string
	}

	cases := []testcase{{
		argv:    []string{"rbmk"},
		failure: "",
	}, {
		argv:    []string{"rbmk", "-h"},
		failure: "",
	}, {
		argv:    []string{"rbmk", "--help"},
		failure: "",
	}, {
		argv:    []string{"rbmk", "help"},
		failure: "",
	}, {
		argv:    []string{"rbmk", "help", "env"},
		failure: "",
	}, {
		argv:    []string{"rbmk", "help", "__nonexistent__"},
		failure: "no such help topic",
	}, {
		argv:    []string{"rbmk", "env"},
		failure: "",
	}, {
		argv:    []string{"rbmk", "__nonexistent__"},
		failure: "no such command",
	}}

	for _, tc := range cases {
		t.Run(strings.Join(tc.argv, " "), func(t *testing.T) {
			cmd := cliutils.NewCommandWithSubCommands(
				"rbmk", "", map[string]cliutils.Command{
					"env": fakecmd{},
				},
			)
			err := cmd.Main(context.Background(), tc.argv...)
			switch {
			case tc.failure == "" && err == nil:
				// all good
			case tc.failure != "" && err == nil:
				t.Fatal("expected", tc.failure, "got", err)
			case tc.failure == "" && err != nil:
				t.Fatal("expected", tc.failure, "got", err.Error())
			case tc.failure != "" && err != nil:
				if err.Error() != tc.failure {
					t.Fatal("expected", tc.failure, "got", err.Error())
				}
			}
		})
	}
}

func TestHelpRequested(t *testing.T) {
	type testcase struct {
		argv   []string
		expect bool
	}

	cases := []testcase{
		// Empty argv and absence of any argument
		{
			argv:   []string{},
			expect: false,
		},
		{
			argv:   []string{"dig"},
			expect: false,
		},

		// Standalone help, help followed by arguments, help not as
		// the second position in the argv
		{
			argv:   []string{"dig", "help"},
			expect: true,
		},
		{
			argv:   []string{"dig", "help", "log.jsonl"},
			expect: true,
		},
		{
			argv:   []string{"dig", "--log-file", "log.jsonl", "help"},
			expect: false,
		},

		// With -h or --help as the last argument
		{
			argv:   []string{"dig", "--log-file", "log.jsonl", "-h"},
			expect: true,
		},
		{
			argv:   []string{"dig", "--log-file", "log.jsonl", "--help"},
			expect: true,
		},

		// Without any of -h, --help, or help
		{
			argv:   []string{"dig", "--log-file", "log.jsonl"},
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(strings.Join(tc.argv, " "), func(t *testing.T) {
			got := cliutils.HelpRequested(tc.argv...)
			if got != tc.expect {
				t.Fatal("expected", tc.expect, "got", got)
			}
		})
	}
}
