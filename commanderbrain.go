package main

import "fmt"
import "os"

//import "gitlab.com/gomidi/midi/reader"

var version = "1.0"


func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Commanderbrain v" + version + " online...")

	//midiFile, err := os.Open("/home/patch/go/bin/midi_experiment.mid")
	//checkErr(err)


}

