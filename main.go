package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type audioPanel struct {
	name       string
	sampleRate beep.SampleRate
	streamer   beep.StreamSeeker
	ctrl       *beep.Ctrl
	resampler  *beep.Resampler
	volume     *effects.Volume
}

type collection struct {
	songs []*audioPanel
}

const SR beep.SampleRate = 48000

func newAudioPanel(name string, sampleRate beep.SampleRate, streamer beep.StreamSeeker) *audioPanel {
	fmt.Println("new Audio")
	strmr, _ := beep.Loop2(streamer)
	ctrl := &beep.Ctrl{Streamer: strmr}
	resampler := beep.Resample(4, sampleRate, SR, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2, Volume: -5.5}
	return &audioPanel{name, sampleRate, streamer, ctrl, resampler, volume}
}

func initSpeaker() {
	fmt.Println("Speaker init Start")
	speaker.Init(SR, 4800)
	fmt.Println("Speaker init Done")
}

func songCollector(temp *collection) {
	fmt.Println("Start song Collecting")
	songs := []string{}
	originPath := "./songs"
	items, _ := os.ReadDir("./songs")
	for _, item := range items {
		if !item.IsDir() && filepath.Ext(item.Name()) == ".mp3" {
			tempPath := filepath.Join(originPath, item.Name())
			songs = append(songs, tempPath)
		} else {
			continue
		}
	}
	for _, e := range songs {
		f, err := os.Open(e)
		if err != nil {
			report(err)
		}
		streamer, format, err := mp3.Decode(f)
		if err != nil {
			report(err)
		}
		temp.songs = append(temp.songs, newAudioPanel(f.Name(), format.SampleRate, streamer))
	}
	fmt.Println("Done song Collecting")
}

func (c *collection) playCollection() {
	fmt.Println("Start play collection")
	for _, e := range c.songs {
		speaker.Play(e.volume)
	}
	fmt.Println("Done play collection")
}

func (c *collection) changeVolumePerSongInCollection() {
	fmt.Println("Start changing volume per song in collection")
	var l = 0
	for i, e := range c.songs {
		fmt.Println(i, e.name)
		l = i
	}
	fmt.Println("What song to change?")
	var i int

	fmt.Print("Enter [number] to change the volume of a song.\n")
	fmt.Scanln(&i)
	if i <= -1 {
	} else if i > l {
	} else {
		var temp float64
		fmt.Print("Enter the volume you want:\n")
		fmt.Print("To stop any song enter number (float64) >= 2\n")
		fmt.Scanln(&temp)
		if temp >= 2 {
			c.songs[i].ctrl.Paused = true
		} else {
			c.songs[i].ctrl.Paused = false
		}
		c.songs[i].volume.Volume = temp
	}

	fmt.Println("Done changing volume per song in collection")
}

func main() {
	initSpeaker()
	fmt.Println("Volume by default -5.5")
	tmp := new(collection)
	songCollector(tmp)
	tmp.playCollection()
	for {
		fmt.Print("Press [ENTER] to pause/resume. ")
		fmt.Scanln()
		speaker.Lock()
		tmp.changeVolumePerSongInCollection()
		speaker.Unlock()
	}
}

func report(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
