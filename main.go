package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli"

	//"github.com/redbeardlab/mqtt-stereo/backend"
	"./backend"
)

func main() {
	app := cli.NewApp()
	app.Name = "mqtt-player"
	log.Print("Start application")
	log.SetOutput(os.Stderr)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "record",
			Usage: "What file use to record and playback",
			Value: "mqtt-record.txt",
		},
		cli.StringFlag{
			Name:  "topic",
			Usage: "What topic listen and what topic play back",
			Value: "/#",
		},
		cli.StringFlag{
			Name:  "url",
			Usage: "Where to listen to the MQTT broker",
			Value: "localhost",
		},
		cli.IntFlag{
			Name:  "port",
			Usage: "What port to use to connect to the broker",
			Value: 1883,
		},
		cli.StringFlag{
			Name:  "user",
			Usage: "User name of the client",
			Value: "",
		},
		cli.StringFlag{
			Name:  "password",
			Usage: "User password of the client",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "record",
			Aliases: []string{"r", "rec"},
			Usage:   "Record traffic from a MQTT broker",
			Action: func(c *cli.Context) error {
				log.Print("REQUEST RECORD : ...")
				log.Print("    url=" + c.GlobalString("url"))
				log.Print("    port=" + c.GlobalString("port"))
				log.Print("    user=" + c.GlobalString("user"))
				log.Print("    password=" + c.GlobalString("password"))
				log.Print("    topic=" + c.GlobalString("topic"))
				backend.StartRecording(c)
				log.Print("DONE RECORD : ...")
				return nil
			},
		},
		{
			Name:    "play",
			Aliases: []string{"p"},
			Usage:   "Play back previous registered traffic to a MQTT broker",
			Action: func(c *cli.Context) error {
				fmt.Println("Playing back the recorded traffic from: ", c.GlobalString("record"))
				backend.PlayBack(c)
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "ff",
					Usage: "Fast forward, if true plays all the messages in the file without respecting the times differences"},

				cli.BoolFlag{
					Name:  "loop",
					Usage: "Loop the player and keep playing all the messages, in order, indefinitely",
				},
			},
		},
	}

	app.Run(os.Args)
}
