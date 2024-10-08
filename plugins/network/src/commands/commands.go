package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dokku/dokku/plugins/common"
)

const (
	helpHeader = `Usage: dokku network[:COMMAND]

Manage network settings for an app

Additional commands:`

	helpContent = `
    network:create <network>, Creates an attachable docker network
    network:destroy <network>, Destroys a docker network
    network:exists <network>, Checks if a docker network exists
    network:info <network> [--format text|json], Outputs information about a docker network
    network:list [--format text|json], Lists all docker networks
    network:rebuild <app>, Rebuilds network settings for an app
    network:rebuildall, Rebuild network settings for all apps
    network:report [<app>] [<flag>], Displays a network report for one or more apps
    network:set <app> <property> (<value>...), Set or clear a network property for an app`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "network", "network:help":
		usage()
	case "help":
		result, err := common.CallExecCommand(common.ExecCommandInput{
			Command: "ps",
			Args:    []string{"-o", "command=", strconv.Itoa(os.Getppid())},
		})
		if err == nil && strings.Contains(result.StdoutContents(), "--all") {
			fmt.Println(helpContent)
		} else {
			fmt.Print("\n    network, Manage network settings for an app\n")
		}
	default:
		dokkuNotImplementExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("failed to retrieve DOKKU_NOT_IMPLEMENTED_EXIT environment variable")
			dokkuNotImplementExitCode = 10
		}
		os.Exit(dokkuNotImplementExitCode)
	}
}

func usage() {
	common.CommandUsage(helpHeader, helpContent)
}
