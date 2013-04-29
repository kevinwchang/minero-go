// Package cmd provides an easily extendable interface to work with chat commands.
package cmd

import (
	"github.com/toqueteos/minero/server/player"
)

type Cmder interface {
	Tab(args []string) []string
	Do(from *player.Player, args []string) bool
}