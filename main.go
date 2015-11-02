package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app            = kingpin.New("scm", "ssh config manager.")
	dump           = app.Command("dump", "")
	list           = app.Command("list", "List all Host.")
	show           = app.Command("show", "Dumps one section.")
	showHost       = show.Arg("host", "Name for a host").Required().String()
	add            = app.Command("add", "Add host")
	addHost        = add.Arg("host", "Name for a host").Required().String()
	addHostName    = add.Arg("hostname", "HostName of the specified host").Required().String()
	addUser        = add.Flag("user", "").Short('u').String()
	addPort        = add.Flag("port", "").Short('p').String()
	addIdentify    = add.Flag("identify", "").Short('i').String()
	addParams      = add.Flag("params", "").Short('P').StringMap()
	update         = app.Command("update", "")
	updateHost     = update.Arg("host", "").Required().String()
	updateHostName = update.Arg("hostname", "HostName of the specified host").Required().String()
	remove         = app.Command("remove", "")
	removeHost     = remove.Arg("host", "").Required().String()
)

var (
	hosts           Hosts
	ssh_config_file = os.ExpandEnv("$HOME/.ssh/config")
)

func addCommand(name, ip, user, port, identify string, params map[string]string) {
	hosts = hosts.addHost(name, ip, user, port, identify, params)
	hosts.saveConfig(ssh_config_file)
}

func listCommand() {
	for _, v := range hosts.listHost() {
		fmt.Println(v)
	}
}

func showCommand(name string) {
	_, host := hosts.GetHost(name)
	fmt.Println(host)
}

func removeCommand(name string) {
	if hosts = hosts.removeHost(name); hosts != nil {
		hosts.saveConfig(ssh_config_file)
	} else {
		log.Fatalf("host %s is not found.\n", name)
	}
}

func main() {
	file, err := os.Open(ssh_config_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hosts, err = Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case dump.FullCommand():
		fmt.Print(hosts)
	case show.FullCommand():
		showCommand(*showHost)
	case add.FullCommand():
		addCommand(*addHost, *addHostName, *addUser, *addPort, *addIdentify, *addParams)
	case list.FullCommand():
		listCommand()
	case update.FullCommand():
		fmt.Println(hosts.updateHost(*updateHost, *updateHostName))
	case remove.FullCommand():
		removeCommand(*removeHost)
	}
}
