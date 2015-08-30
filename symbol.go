// Package symbol provide colored log icons
package symbol // import "github.com/shamsher31/gosymbol"

import "github.com/ttacon/chalk"

func Info() string {
	return chalk.Blue.Color("ℹ ")
}

func Success() string {
	return chalk.Green.Color("✔ ")
}

func Warning() string {
	return chalk.Yellow.Color("⚠ ")
}

func Error() string {
	return chalk.Red.Color("✖ ")
}
