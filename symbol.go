// Package symbol provide colored log icons
package symbol // import "github.com/shamsher31/gosymbol"

import "github.com/ttacon/chalk"

var symbols = map[string]string{
	"Info":        "ℹ ",
	"Success":     "✔ ",
	"Warning":     "⚠ ",
	"Error":       "✖ ",
	"Copyright":   "© ",
	"Registered":  "® ",
	"Pi":          "π ",
	"Omega":       "Ω ",
	"Theta":       "Θ ",
	"Beta":        "β ",
	"Delta":       "δ ",
	"Microsecond": "µ ",
}

func Info() string {
	return chalk.Blue.Color(symbols["Info"])
}

func Success() string {
	return chalk.Green.Color(symbols["Success"])
}

func Warning() string {
	return chalk.Yellow.Color(symbols["Warning"])
}

func Error() string {
	return chalk.Red.Color(symbols["Error"])
}

func Copyright() string {
	return chalk.White.Color(symbols["Copyright"])
}

func Registered() string {
	return chalk.White.Color(symbols["Registered"])
}

func Pi() string {
	return chalk.White.Color(symbols["Pi"])
}

func Omega() string {
	return chalk.White.Color(symbols["Omega"])
}

func Theta() string {
	return chalk.White.Color(symbols["Theta"])
}

func Beta() string {
	return chalk.White.Color(symbols["Beta"])
}

func Delta() string {
	return chalk.White.Color(symbols["Delta"])
}

func Microsecond() string {
	return chalk.White.Color(symbols["Microsecond"])
}
