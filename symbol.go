// Package symbol provide colored log icons
package symbol // import "github.com/shamsher31/gosymbol"

import (
	"fmt"
	"github.com/ttacon/chalk"
)

func Info() {
	fmt.Print(chalk.Blue, "ℹ", " ", chalk.Reset)
}

func Success() {
	fmt.Print(chalk.Green, "✔", " ", chalk.Reset)
}

func Warning() {
	fmt.Print(chalk.Yellow, "⚠", " ", chalk.Reset)
}

func Error() {
	fmt.Print(chalk.Red, "✖", " ", chalk.Reset)
}
