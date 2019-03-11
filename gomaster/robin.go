package main

import (
	"log"
	"time"

	"github.com/jiansoft/robin"
)

func main() {
	//The method is going to execute only once after 2000 ms.
	robin.Delay(2000).Do(runCron, "a Delay 2000 ms")

	minute := 11
	second := 50

	//Every Friday is going to execute once at 14:11:50 (HH:mm:ss).
	robin.EveryFriday().At(14, minute, second).Do(runCron, "Friday")

	//Every N day  is going to execute once at 14:11:50(HH:mm:ss)
	robin.Every(1).Days().At(14, minute, second).Do(runCron, "Days")

	//Every N hours is going to execute once at 11:50:00(HH:mm:ss).
	robin.Every(1).Hours().At(0, minute, second).Do(runCron, "Every 1 Hours")

	//Every N minutes is going to execute once at 50(ss).
	robin.Every(1).Minutes().At(0, 0, second).Do(runCron, "Every 1 Minutes")

	//Every N seconds is going to execute once
	robin.Every(10).Seconds().Do(runCron, "Every 10 Seconds")

	p1 := player{Nickname: "Player 1"}
	p2 := player{Nickname: "Player 2"}
	p3 := player{Nickname: "Player 3"}
	p4 := player{Nickname: "Player 4"}

	//Create a channel
	channel := robin.NewChannel()

	//Four player subscribe the channel
	channel.Subscribe(p1.eventFinalBossResurge)
	channel.Subscribe(p2.eventFinalBossResurge)
	p3Subscribe := channel.Subscribe(p3.eventFinalBossResurge)
	p4Subscribe := channel.Subscribe(p4.eventFinalBossResurge)

	//Publish a message to the channel and then the four subscribers of the channel will
	//receives the message each that "The boss resurge first." .
	channel.Publish("The boss resurge first.")

	//Unsubscribe p3 and p4 from the channel.
	channel.Unsubscribe(p3Subscribe)
	p4Subscribe.Unsubscribe()

	//This time just p1 and p2 receives the message that "The boss resurge second.".
	channel.Publish("The boss resurge second.")

	//Unsubscribe all subscribers from the channel
	channel.Clear()

	//The channel is empty so no one can receive the message
	channel.Publish("The boss resurge third.")
}

func runCron(s string) {
	log.Printf("I am %s CronTest %v\n", s, time.Now())
}

type player struct {
	Nickname string
}

func (p player) eventFinalBossResurge(someBossInfo string) {
	log.Printf("%s receive a message : %s", p.Nickname, someBossInfo)
}
