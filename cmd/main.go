package main

import (
	"AquaBot/packages/comPort"
	"AquaBot/packages/computer"
	"AquaBot/packages/myBot"
	"AquaBot/packages/structs"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main(){
	bot := myBot.New(getToken(),&structs.Vars)

	bot.Start()

	input := comPort.New("COM20", 9600)
	fmt.Println("Hello world")
	computer.New(input, &structs.Vars)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	select{}
}


func getToken()string{
	pwd, err := os.Getwd()
	tokenB, err := os.ReadFile(pwd + "/files/token.txt")
	if err != nil{
		log.Fatal(err)
	}
	return string(tokenB)
}