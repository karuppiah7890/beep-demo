package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/faiface/beep"
	"google.golang.org/grpc"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/karuppiah7890/beep-demo/stream"
)

type streamServer struct {
	streamer beep.StreamSeekCloser
}

func (s streamServer) Position(_ *stream.PositionRequest, respStream stream.Stream_PositionServer) error {
	ticker := time.NewTicker(500 * time.Millisecond)

	for range ticker.C {
		speaker.Lock()
		position := s.streamer.Position()
		err := respStream.Send(&stream.PositionResponse{Position: int64(position)})

		if err != nil {
			log.Fatalf("error occurred while sending response %v", err)
			return err
		}

		speaker.Unlock()

		len := s.streamer.Len()
		if position == len {
			ticker.Stop()
			break
		}
	}

	return nil
}

func newStreamServer(streamer beep.StreamSeekCloser) *streamServer {
	return &streamServer{streamer: streamer}
}

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
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	argsLength := len(os.Args)

	if argsLength != 1 && argsLength != 2 {
		fmt.Printf("pass exactly one argument!")
		return
	}

	if argsLength == 2 {
		// Set up a connection to the server.
		conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		streamClient := stream.NewStreamClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		respStream, err := streamClient.Position(ctx, &stream.PositionRequest{})
		if err != nil {
			log.Fatalf("could not get stream position: %v", err)
		}

		positionResponse, err := respStream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("%v.Position(_) = _, %v", streamClient, err)
		}
		if err = streamer.Seek(int(positionResponse.Position)); err != nil {
			log.Fatalf("could not seek to position in stream: %v", err)
		}

		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		for {
			positionResponse, err := respStream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("%v.Position(_) = _, %v", streamClient, err)
			}
			if err = streamer.Seek(int(positionResponse.Position)); err != nil {
				log.Fatalf("could not seek to position in stream: %v", err)
			}
		}

		<-done
	} else if argsLength == 1 {
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		lis, err := net.Listen("tcp", "0.0.0.0:8080")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		stream.RegisterStreamServer(grpcServer, newStreamServer(streamer))

		go func() {
			<-done
			grpcServer.GracefulStop()
		}()

		err = grpcServer.Serve(lis)

		if err != nil {
			fmt.Printf("error occurred while starting grpc server : %v", err)
		}
	}
}
