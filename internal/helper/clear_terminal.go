package helper

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

//https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearTerminal() error {
	if value, ok := clear[runtime.GOOS]; ok {
		value()
		return nil
	}

	return errors.New("Your platform is unsupported! I can't clear terminal screen :(")
}
