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

func Copyright() string {
	return chalk.White.Color("© ")
}

func Registered() string {
	return chalk.White.Color("® ")
}

func Pi() string {
	return chalk.White.Color("π ")
}

func Omega() string {
	return chalk.White.Color("Ω ")
}

func Theta() string {
	return chalk.White.Color("Θ ")
}

func Beta() string {
	return chalk.White.Color("β ")
}

func Delta() string {
	return chalk.White.Color("δ ")
}
