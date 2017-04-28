package main

import (
	"fmt"
	"os/exec"

	"github.com/go-ini/ini"
	"log"
)

type Config struct {
	SSH_Key string
	Git_url string
	Deploy_dir string
	Container_name string
	Secret string
}

func main() {
	//Init Config
	cfg, _ := ini.Load("config.ini")
	var Conf Config
	Conf.SSH_Key = cfg.Section("").Key("ssh_key").String()
	Conf.Deploy_dir = cfg.Section("").Key("deploy_dir").String()
	Conf.Git_url = cfg.Section("").Key("git_url").String()
	Conf.Container_name = cfg.Section("").Key("container_name").String()
	Conf.Secret = cfg.Section("").Key("secret").String()

	if Conf.Git_url != "" && Conf.Deploy_dir != "" && Conf.SSH_Key != "" {
		cmd, err := exec.Command("/bin/bash", "deploy.sh", "-k", Conf.SSH_Key, "-g", Conf.Git_url, "-d", Conf.Deploy_dir, "-c", Conf.Container_name).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", cmd)
	} else {
		fmt.Println("You must provide at least a Git-Url, Deploy-Dir and SSH-Key!")
	}
}
