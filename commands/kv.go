package commands

import (
	"github.com/spf13/cobra"
)

type Kv struct {
	*Cmd
}

func (root *Cmd) initKv() {
	k := Kv{Cmd: root}

	kvCmd := &cobra.Command{
		Use:   "kv",
		Short: "Consul K/V endpoint interface",
		Long:  "Consul K/V endpoint interface",
		Run: func(cmd *cobra.Command, args []string) {
			root.Help()
		},
	}

	k.AddDeleteSub(kvCmd)
	k.AddLockSub(kvCmd)
	k.AddReadSub(kvCmd)
	k.AddUnlockSub(kvCmd)
	k.AddWatchSub(kvCmd)
	k.AddWriteSub(kvCmd)

	k.AddCommand(kvCmd)
}
