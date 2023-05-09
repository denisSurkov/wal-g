package pg

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
	"github.com/wal-g/wal-g/internal"
	"github.com/wal-g/wal-g/internal/databases/postgres"
)

const WalPrefetchShortDescription = `Used for prefetching process forking
and should not be called by user.`

// WalPrefetchCmd represents the walPrefetch command
var WalPrefetchCmd = &cobra.Command{
	Use:    "wal-prefetch wal_name prefetch_location",
	Short:  WalPrefetchShortDescription,
	Args:   cobra.ExactArgs(2),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		reconfigureLoggers()

		folder, err := internal.ConfigureFolder()
		tracelog.ErrorLogger.FatalOnError(err)

		postgres.HandleWALPrefetch(internal.NewFolderReader(folder), args[0], args[1])
	},
}

func init() {
	Cmd.AddCommand(WalPrefetchCmd)
}

// wal-prefetch (WalPrefetchCmd) is internal tool, so to avoid confusion about errors in restoration process
// we reconfigure loggers specially for internal use. All logs having PREFETCH prefix can be safely ignored
func reconfigureLoggers() {
	tracelog.ErrorLogger.SetPrefix(fmt.Sprintf("PREFETCH %s", tracelog.ErrorLogger.Prefix()))

	tracelog.DebugLogger.SetPrefix(fmt.Sprintf("PREFETCH %s", tracelog.DebugLogger.Prefix()))

	tracelog.WarningLogger.SetPrefix(fmt.Sprintf("PREFETCH %s", tracelog.WarningLogger.Prefix()))

	tracelog.InfoLogger.SetPrefix(fmt.Sprintf("PREFETCH %s", tracelog.InfoLogger.Prefix()))
}
