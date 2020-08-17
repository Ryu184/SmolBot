package roll

import (
	"time"
	"math/rand"

	"github.com/jonas747/dcmd"
	"github.com/jonas747/yagpdb/commands"
)

var Command = &commands.YAGCommand{
	Cooldown:    2,
	CmdCategory: commands.CategoryFun,
	Name:        "8Ball",
	Description: "Wisdom",
	Arguments: []*dcmd.ArgDef{
		&dcmd.ArgDef{Name: "What to ask", Type: dcmd.String},
	},
	RequiredArgs: 1,
	RunFunc: func(cmd *dcmd.Data) (interface{}, error) {
		rand.Seed(time.Now().UnixNano())
		answers := []string{
			"It is certain",
			"It is decidedly so",
			"Without a doubt",
			"Yes definitely",
			"You may rely on it",
			"As I see it yes",
			"Most likely",
			"Outlook good",
			"Yes",
			"Signs point to yes",
			"Reply hazy try again",
			"Ask again later",
			"Better not tell you now",
			"Cannot predict now",
			"Concentrate and ask again",
			"Don't count on it",
			"My reply is no",
			"My sources say no",
			"Outlook not so good",
			"Very doubtful",
		}
		out := "Magic 8-Ball says: "
		out += answers[rand.Intn(len(answers))]
		return out, nil
	},
}