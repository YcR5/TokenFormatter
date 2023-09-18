package main

import (
	"DiscordVanity/utilities"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
)

var (
	Art = `
	 ██  ██  ███████  ██████  ██████  ███    ███  █████  ████████ ████████ ███████ ██████  
	████████ ██      ██    ██ ██   ██ ████  ████ ██   ██    ██       ██    ██      ██   ██ 
	 ██  ██  █████   ██    ██ ██████  ██ ████ ██ ███████    ██       ██    █████   ██████  
	████████ ██      ██    ██ ██   ██ ██  ██  ██ ██   ██    ██       ██    ██      ██   ██ 
	 ██  ██  ██       ██████  ██   ██ ██      ██ ██   ██    ██       ██    ███████ ██   ██ 
																																													
 	Discord : @26mm
	Instagram : @ldle
	Made With Love By : #RedEye..

`
	ListName string

	Threads string

	Postition int
)

func main() {

	prompThq := &survey.Input{
		Message: "Text File - Example (List.txt): ",
	}
	_ = survey.AskOne(prompThq, &ListName, survey.WithValidator(survey.Required))

	prompTh := &survey.Input{
		Message: "Spliter Example (: or , or ; or / depends on what u want): ",
	}
	_ = survey.AskOne(prompTh, &Threads, survey.WithValidator(survey.Required))

	prompThs := &survey.Input{
		Message: "Postition Example \n(0 = Test1:test2:Test3 IF U WANT Test1) \n(1 = Test1:test2:Test3 IF U WANT test2) \n(2 = Test1:test2:Test3 IF U WANT test3) \nSo Just Write 1 Or 2 Or .... : ",
	}
	_ = survey.AskOne(prompThs, &Postition, survey.WithValidator(survey.Required))

	file, err := os.Open(ListName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), Threads)
		fmt.Println(split[Postition])
		file, _ := os.OpenFile("Splited.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		datawriter := bufio.NewWriter(file)
		_, _ = datawriter.WriteString(split[Postition] + "\n")
		_ = datawriter.Flush()
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", count)

	utilities.Exit()
}

func init() {
	color.Red.Printf(Art)
}
