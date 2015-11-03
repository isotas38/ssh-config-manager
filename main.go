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
	addParams      = add.Flag("params", "").PlaceHolder("KEY:VALUE").Short('P').StringMap()
	update         = app.Command("update", "")
	updateHost     = update.Arg("host", "").Required().String()
	updateHostName = update.Flag("hostname", "HostName of the specified host").Short('h').String()
	updateUser     = update.Flag("user", "").Short('u').String()
	updatePort     = update.Flag("port", "").Short('p').String()
	updateIdentify = update.Flag("identify", "").Short('i').String()
	updateParams   = update.Flag("params", "").PlaceHolder("KEY:VALUE").Short('P').StringMap()
	move           = app.Command("mv", "Rename Host")
	moveOldHost    = move.Arg("old_host", "").Required().String()
	moveNewHost    = move.Arg("new_host", "").Required().String()
	cp             = app.Command("cp", "")
	cpOldHost      = cp.Arg("old_host", "").Required().String()
	cpNewHost      = cp.Arg("new_host", "").Required().String()
	remove         = app.Command("rm", "")
	removeHost     = remove.Arg("host", "").Required().String()
)

var (
	hosts           Hosts
	ssh_config_file = os.ExpandEnv("$HOME/.ssh/config")
)

func addCommand(name, ip, user, port, identify string, params map[string]string) {
	if hosts = hosts.addHost(name, ip, user, port, identify, params); hosts != nil {
		hosts.saveConfig(ssh_config_file)
	} else {
		log.Fatalf("host %s is already exist\n", name)
	}
}

func updateCommand(name, ip, user, port, identify string, params map[string]string) {
	if hosts = hosts.updateHost(name, ip, user, port, identify, params); hosts != nil {
		hosts.saveConfig(ssh_config_file)
	} else {
		log.Fatalf("host %s is not found.\n", name)
	}
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

func moveCommand(old_host, new_host string) {
	if hosts = hosts.moveHost(old_host, new_host); hosts != nil {
		hosts.saveConfig(ssh_config_file)
	} else {
		log.Fatalf("host %s is not found.\n", old_host)
	}
}

func cpCommand(old_host, new_host string) {
	if hosts = hosts.copyHost(old_host, new_host); hosts != nil {
		hosts.saveConfig(ssh_config_file)
	} else {
		log.Fatalf("host %s is not found.\n", old_host)
	}
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
		updateCommand(*updateHost, *updateHostName, *updateUser, *updatePort, *updateIdentify, *updateParams)
	case move.FullCommand():
		moveCommand(*moveOldHost, *moveNewHost)
	case cp.FullCommand():
		cpCommand(*cpOldHost, *cpNewHost)
	case remove.FullCommand():
		removeCommand(*removeHost)
	}
}
