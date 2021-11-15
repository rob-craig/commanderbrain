package main

import "fmt"
import "strconv"
import "log"
import "os"

import "gitlab.com/gomidi/midi/reader"

var version = "1.0"
var trackOneNotes map[uint64]string

var maxPos uint64 = 0;

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func setupLogFile() {

	f, err := os.OpenFile("testlogfile.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    	log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

}

func loadTrackNotes() {

	trackOneNotes = make(map[uint64]string)

	midirdr := reader.New(reader.NoLogger(), 
						  reader.NoteOn(loadNoteOn),
						  reader.NoteOff(loadNoteOff),)

	err := reader.ReadSMFFile(midirdr, "/home/patch/go/bin/midi_experiment.mid")
	checkErr(err)
	


}

func loadNoteOn(pos *reader.Position, channel, key, velocity uint8) {
	trackOneNotes[pos.AbsoluteTicks] = "NOTE ON"
	maxPos = pos.AbsoluteTicks;
}

func loadNoteOff(pos *reader.Position, channel, key, velocity uint8) {
	trackOneNotes[pos.AbsoluteTicks] = "NOTE OFF"
	maxPos = pos.AbsoluteTicks;
}


func main() {
	setupLogFile()
	
	fmt.Println("Commanderbrain v" + version + " online...")
	log.Println("Commanderbrain v" + version + " online...")
	

	fmt.Println("loading midi file into memory")
	log.Println("loading midi file into memory")
	loadTrackNotes()

	fmt.Println("ok check this test out:")
	log.Println("ok check this test out:")

	var i uint64
	for i=0;i<maxPos;i++ {

		var msg = strconv.FormatUint(i, 10) + ": "
		if val, ok := trackOneNotes[i]; ok {
			msg += val;
			
		} else {
			msg += "..."
		}	
		
		log.Println(msg)

	}







}

