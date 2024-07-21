package server

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

func Main() int {

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	
	hostFlag := flags.String("host", "localhost", "specific host for the server to listen on (default is localhost)")
	portFlag := flags.String("port", "8080", "specific port to listen on (default is 8080)")
	addrFlag := flags.String("addr", "localhost:8080", "address [host:port] to listen; don't use this flag if any of the host or port flag is used")

	flags.Parse(os.Args[1:])

	rootDir := "."
	if len(flags.Args())==1 {
		rootDir = flags.Arg(0)
	}
	
	flagSets := flagsSet(flags)

	var addr string

	if (flagSets["addr"]){
		addr = *addrFlag
	}else {
		addr = *hostFlag + ":" + *portFlag
	}

	srv := &http.Server{
		Addr: addr,
	}

	mux := http.NewServeMux()

	fileHandler := http.FileServer(http.Dir(rootDir))
	mux.Handle("/", fileHandler)

	srv.Handler = mux

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listening on the port: %v", err)
		return 1
	}

	fmt.Fprintf(os.Stdout, "Server is running on %s", addr)

	if err := srv.Serve(listener); err != nil {
		return 1
	}

	return 0
}

func flagsSet(flags *flag.FlagSet) map[string]bool {
	s := make(map[string]bool)
	flags.Visit(func(f *flag.Flag){
		s[f.Name] = true
	})
	return s
}