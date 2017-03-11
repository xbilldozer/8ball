package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var magicAnswers = [...]string{
	// Positive outcomes
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",

	// Neutral outcomes
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",

	// Negative outcomes
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

var ShakeFailureMessage = "You can't just shake me bro, ask a question!"

func main() {
	fmt.Println(Shake())
}

func Shake() string {
	if len(os.Args) < 2 {
		return ShakeFailureMessage
	}

	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("Magic 8-Ball says: %s", magicAnswers[rand.Intn(len(magicAnswers))])
}
