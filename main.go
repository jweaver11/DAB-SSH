package main

// An example Bubble Tea server.
// "ssh -p 1234 localhost"

import (
	//"DAB-SSH/models"

	"DAB-SSH/models"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
)

// Sets the host server as local on the port 1234
const (
	host = "localhost"
	port = 1234
)

func main() {
	// Creates a new ssh wish server and uses the teaHandler as the middle ware
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Starts the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Closes the server
	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

// You can wire any Bubble Tea model up to the middleware with a function that
// handles the incoming ssh.Session.
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {

	// Starts ssh session
	_, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, "Could not start ssh session")
		return nil, nil
	}

	// First model user sees when connecting
	// Model controls connections from there
	m := models.CreateWelcomePage()

	// Returns our model in full screen mode
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
