package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Vstural/mailer/pkg/config"
	"github.com/Vstural/mailer/pkg/mailer"
	"github.com/Vstural/mailer/pkg/utils"
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
	err = mailer.SendOutlookMail(
		conf,
		to,
		title,
		content,
	)
	if err != nil {
		panic(err)
	}
}
