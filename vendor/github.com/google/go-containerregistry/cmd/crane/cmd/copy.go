// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/spf13/cobra"
)

func init() { Root.AddCommand(NewCmdCopy()) }

// NewCmdCopy creates a new cobra.Command for the copy subcommand.
func NewCmdCopy() *cobra.Command {
	return &cobra.Command{
		Use:     "copy SRC DST",
		Aliases: []string{"cp"},
		Short:   "Efficiently copy a remote image from src to dst",
		Args:    cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, args []string) {
			src, dst := args[0], args[1]
			if err := crane.Copy(src, dst, options...); err != nil {
				log.Fatal(err)
			}
		},
	}
}
