// Package symbol provide colored log icons
package symbol // import "github.com/shamsher31/gosymbol"

import (
	"fmt"
	"github.com/ttacon/chalk"
)

func Info() {
	fmt.Println(chalk.Blue, "ℹ", chalk.Reset)
}

func Success() {
	fmt.Println(chalk.Green, "✔", chalk.Reset)
}

func Warning() {
	fmt.Println(chalk.Yellow, "⚠", chalk.Reset)
}

func Error() {
	fmt.Println(chalk.Red, "✖", chalk.Reset)
}
