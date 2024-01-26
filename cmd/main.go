package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd := cobra.Command{
		Use:   "app",
		Short: "CA API",
		Run: func(*cobra.Command, []string) {
			serv()
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "run-server",
		Short: "Run CA API Server",
		Run: func(*cobra.Command, []string) {
			serv()
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "run-nsq",
		Short: "Run CA API NSQ Consumer",
		Run: func(*cobra.Command, []string) {
			nsq()
		},
	})

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
