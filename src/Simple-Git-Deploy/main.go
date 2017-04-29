package main

import (
	"fmt"
	//"os/exec"

	"github.com/go-ini/ini"
	"log"
	"net/http"
	//"encoding/json"
	//"os/exec"
	//"github.com/go-playground/form"
	//"encoding/json"
	"github.com/go-playground/form"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"os/exec"
)

type Config struct {
	SSH_Key        string
	Git_url        string
	Deploy_dir     string
	Container_name string
	Secret         string
	Interface      string
}

func config(key string) string {
	cfg, _ := ini.Load("config.ini")
	return cfg.Section("").Key(key).String()
}

type Payload struct {
	Secret string `json:"secret"`
	Ref    string `json:"ref"`
}

func main() {
	if config("git_url") != "" && config("deploy_dir") != "" && config("ssh_key") != "" {

		// Server
		http.HandleFunc("/", handleHook)
		err := http.ListenAndServe(config("interface"), nil) // setting listening port
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		fmt.Println("You must provide at least a Git-Url, Deploy-Dir and SSH-Key!")
	}
}

var decoder *form.Decoder

func handleHook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading response: %v", err)
		}

		encoded, _ := url.QueryUnescape(string(body[8:]))
		//log.Println(encoded)

		var hook_data Payload
		err = json.Unmarshal([]byte(encoded), &hook_data)
		if err != nil {
			log.Printf("error decoding response: %v", err)
		}
		log.Println("Recived payload, Secret:", hook_data.Secret, ", Ref:", hook_data.Ref)

		if hook_data.Secret == config("secret") {
			if hook_data.Ref == config("branch_name") || hook_data.Ref == "refs/heads/" + config("secret") {
				log.Println("Recived corresponding secret: ", hook_data.Secret)
				log.Println("Starting update...")
				cmd, err := exec.Command("/bin/bash", "deploy.sh", "-k", config("ssh_key"), "-g", config("git_url"), "-d", config("deploy_dir"), "-c", config("container_name")).Output()
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s\n", cmd)
				log.Println("Finished update.")
			}
		}
	}
}
