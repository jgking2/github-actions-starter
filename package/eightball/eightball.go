package eightball

import (
	"math/rand"
	"time"
)

var fortunes = [20]string{
	"As I see it, yes.",
	"Ask again later.",
	"Better not tell you now.",
	"Cannot predict now.",
	"Concentrate and ask again.",
	"Don’t count on it.",
	"It is certain.",
	"It is decidedly so.",
	"Most likely.",
	"My reply is no.",
	"My sources say no.",
	"Outlook not so good.",
	"Outlook good.",
	"Reply hazy, try again.",
	"Signs point to yes.",
	"Very doubtful.",
	"Without a doubt.",
	"Yes.",
	"Yes – definitely.",
	"You may rely on it.",
}

// DetermineMyFortune uses advanced machine learning and the latest in psychic...ology to determine the answer to your deepest concerns
func DetermineMyFortune(question string) string {
	//Question is actually irrelavant...suckers...
	rand.Seed(time.Now().UnixNano())
	fortuneIndex := rand.Intn(len(fortunes))
	return fortunes[fortuneIndex]
}
