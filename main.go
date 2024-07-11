package main

import (
	"os"
	"os/signal"
	"syscall"
)

// var logger = log.New(flag.CommandLine.Output(), "", log.LstdFlags)

// func panicGateError(err error) {
// 	if e, ok := err.(gateiows.GateAPIError); ok {
// 		log.Fatal(fmt.Sprintf("Gate API error, label: %s, message: %s", e.Label, e.Message))
// 	}
// 	log.Fatal(err)
// }

func main() {

	ConnectGateioWs()

	ch := make(chan os.Signal)
	signal.Ignore(syscall.SIGPIPE, syscall.SIGALRM)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT, syscall.SIGKILL)
	<-ch

}
