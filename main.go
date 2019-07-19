package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"

	olric "github.com/buraksezer/olric"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	f, err := os.Open("derek-clegg-youre-the-dummy.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()
	done := make(chan bool)
	shutdownDB := make(chan bool)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	argsLength := len(os.Args)

	if argsLength != 1 && argsLength != 2 {
		fmt.Printf("pass exactly one argument!")
		return
	}

	if argsLength == 2 {
		// get position from olric

		m, _ := olric.NewMemberlistConfig("local")
		m.BindAddr = "127.0.0.1"
		m.BindPort = 5656
		c := &olric.Config{
			Name:             "127.0.0.1:3636",
			Peers:            []string{"127.0.0.1:5555"},
			MemberlistConfig: m,
		}

		db, err := olric.New(c)
		if err != nil {
			log.Fatalf("could not create db: %v\n", err)
		}

		go func() {
			err := db.Start()
			if err != nil {
				panic(fmt.Sprintf("could not start db: %v\n", err))
			}
		}()

		go func() {
			<-shutdownDB
			err := db.Shutdown(context.Background())
			if err != nil {
				panic(fmt.Sprintf("could not shutdown gracefully %v\n", err))
			}
			done <- true
		}()

		dm := db.NewDMap("beep-demo")
		position, err := dm.Get("position")
		if err != nil {
			log.Fatalf("could not get position from db: %v", err)
		}

		seekPosition, ok := position.(int)
		if !ok {
			panic(fmt.Sprintf("got position value %v in a wrong data type", position))
		}

		if err = streamer.Seek(seekPosition); err != nil {
			log.Fatalf("could not seek to position in stream: %v", err)
		}

		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			shutdownDB <- true
		})))

		ticker := time.NewTicker(100 * time.Millisecond)
		go func() {
			for range ticker.C {
				position, err := dm.Get("position")
				if err != nil {
					log.Fatalf("could not get position from db: %v", err)
				}

				seekPosition, ok := position.(int)
				if !ok {
					panic(fmt.Sprintf("got position value %v in a wrong data type", position))
				}

				speaker.Lock()
				if err = streamer.Seek(seekPosition); err != nil {
					log.Fatalf("could not seek to position in stream: %v", err)
				}
				speaker.Unlock()
			}
		}()

		<-done
		ticker.Stop()
	} else if argsLength == 1 {
		m, _ := olric.NewMemberlistConfig("local")
		m.BindAddr = "127.0.0.1"
		m.BindPort = 5555
		c := &olric.Config{
			Name:             "127.0.0.1:3535", // Unique in the cluster and used by TCP server.
			Peers:            []string{"127.0.0.1:5656"},
			MemberlistConfig: m,
		}
		db, err := olric.New(c)
		if err != nil {
			log.Fatalf("could not create db: %v\n", err)
		}

		go func() {
			err := db.Start()
			if err != nil {
				panic(fmt.Sprintf("could not start db: %v\n", err))
			}
		}()

		dm := db.NewDMap("beep-demo")
		err = dm.Put("position", 0)
		if err != nil {
			log.Fatalf("could not store position into db: %v", err)
		}

		go func() {
			<-shutdownDB
			err := db.Shutdown(context.Background())
			if err != nil {
				panic(fmt.Sprintf("could not shutdown gracefully %v\n", err))
			}
			done <- true
		}()

		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			shutdownDB <- true
		})))

		ticker := time.NewTicker(100 * time.Millisecond)
		go func() {
			for range ticker.C {
				speaker.Lock()
				position := streamer.Position()
				speaker.Unlock()

				if err = dm.Put("position", position); err != nil {
					log.Fatalf("could not store position into db: %v", err)
				}
			}
		}()

		<-done
	}
}
