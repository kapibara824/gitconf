package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name       string `yaml:name`
	Email      string `yaml:email`
	Signingkey string `yaml:signingkey`
	GPGSign    string `yaml:gpgsign`
	Program    string `yaml:program`
}

func changeConfig(command, value string) error {
	cmd := exec.Command("git", "config", "--local", command, value)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	app := cli.NewApp()
	app.Name = "gitchanger"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.HelpFlag,
	}
	app.Action = func(c *cli.Context) {
		var (
			config     Config
			configFile string
		)

		if len(c.Args().Get(0)) != 0 {
			configFile = c.Args().Get(0)
		} else {
			configFile = "./config.yaml"
		}
		log.Println(configFile)
		buf, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(buf, &config)
		if err != nil {
			log.Fatal(err)
		}

		changeConfig("user.name", config.Name)
		changeConfig("user.email", config.Email)
		changeConfig("user.signingkey", config.Signingkey)
		changeConfig("gpg.program", config.Program)
		changeConfig("commit.gpgsign", config.GPGSign)
	}

	app.Run(os.Args)
	cmd := exec.Command("git", "config", "--list", "--local")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
