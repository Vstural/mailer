package main

import (
	"fmt"
	"mailer/pkg/config"
	"mailer/pkg/mailer"
	"mailer/pkg/utils"
	"os"
	"strings"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		utils.PrintHelp()
		return
	}

	targetStr := os.Args[1]
	title := os.Args[2]
	content := os.Args[3]

	to := strings.Split(targetStr, ";")
	err = mailer.SendMail(
		conf,
		to,
		title,
		content,
	)
	if err != nil {
		panic(err)
	}
}
