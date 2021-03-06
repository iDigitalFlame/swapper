// Copyright (C) 2021 PurpleSec Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package main

import (
	"flag"
	"os"

	"github.com/PurpleSec/swapper"
)

const version = "v1.0.1"

const usage = `Sticker Swapper Telegram Bot ` + version + `
Purple Security (losynth.com/purple) 2021

Usage:
  -h              Print this help menu.
  -f <file>       Configuration file path.
  -d              Dump the default configuration and exit.
  -clear-all      Clear the database of ALL DATA before starting up.
`

func main() {
	var (
		args        = flag.NewFlagSet("Sticker Swapper Telegram Bot "+version, flag.ExitOnError)
		file        string
		dump, empty bool
	)
	args.Usage = func() {
		os.Stderr.WriteString(usage)
		os.Exit(2)
	}
	args.StringVar(&file, "f", "", "Configuration file path.")
	args.BoolVar(&dump, "d", false, "Dump the default configuration and exit.")
	args.BoolVar(&empty, "clear-all", false, "Clear the database of ALL DATA before starting up.")

	if err := args.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(usage)
		os.Exit(2)
	}

	if len(file) == 0 && !dump {
		os.Stderr.WriteString(usage)
		os.Exit(2)
	}

	if dump {
		os.Stdout.WriteString(swapper.Defaults)
		os.Exit(0)
	}

	s, err := swapper.New(file, empty)
	if err != nil {
		os.Stdout.WriteString("Error: " + err.Error() + "!\n")
		os.Exit(1)
	}

	if err := s.Run(); err != nil {
		os.Stdout.WriteString("Error: " + err.Error() + "!\n")
		os.Exit(1)
	}
}
