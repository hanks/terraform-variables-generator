package main

import (
	"log"
	"os"
	"sort"

	"github.com/hanks/terraform-variables-generator/cmd"
	c "github.com/hanks/terraform-variables-generator/configs"
	"github.com/hanks/terraform-variables-generator/version"

	"github.com/urfave/cli"
)

func createApp() *cli.App {
	app := cli.NewApp()

	app.Name = "terraform-variables-genrator"
	app.UsageText = "tfvargen [--config|-c vars.toml] [targetDir]"
	app.Usage = "Simple Tool to Generate Variables file from Terraform Configuration."
	app.Version = version.Version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: c.VarConfName,
			Usage: "`CONF_NAME` to customize generated variables.tf, will search in working directory",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(c *cli.Context) error {
		targetDir := c.Args().Get(0)
		conf := c.String("config")

		cmd.Generate(targetDir, conf)

		return nil
	}

	return app
}

func main() {
	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
