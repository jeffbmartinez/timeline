package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jeffbmartinez/cleanexit"
	"github.com/jeffbmartinez/log"

	"github.com/jeffbmartinez/timeline/handler"
	"github.com/jeffbmartinez/timeline/storage/influxdb"
)

const (
	projectName = "timeline"

	exitSuccess      = 0
	exitFailure      = 1
	exitUsageFailure = 2 // Same as golang's flag module uses, hardcoded at https://github.com/golang/go/blob/release-branch.go1.4/src/flag/flag.go#L812
)

func main() {
	cleanexit.SetUpExitOnCtrlC(getPrintPrettyExitMessageFunc(projectName))

	allowAnyHostToConnect, listenPort := getCommandLineArgs()

	http.HandleFunc("/api/event", handler.Event)
	http.HandleFunc("/api/measurement", handler.Measurement)

	listenHost := "localhost"
	if allowAnyHostToConnect {
		listenHost = ""
	}

	connection, err := influxdb.GetClient()
	if err != nil {
		log.Fatalf("Unable to get an influxdb client: %v", err)
	}

	err = influxdb.TestConnection(connection)
	if err != nil {
		fmt.Println("Could not connect to influxDB")
		log.Fatalf("InfluxDB connection test failed: %v", err)
	}

	displayServerInfo(listenHost, listenPort)

	listenAddress := fmt.Sprintf("%v:%v", listenHost, listenPort)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func getPrintPrettyExitMessageFunc(projectName string) func() {
	return func() {
		/* \b is the equivalent of hitting the back arrow. With the two following
		   space characters they serve to hide the "^C" that is printed when
		   ctrl-c is typed.
		*/
		fmt.Printf("\b\b  \n[ctrl-c] %v is shutting down\n", projectName)
	}
}

func getCommandLineArgs() (allowAnyHostToConnect bool, port int) {
	const defaultPort = 8000

	flag.BoolVar(&allowAnyHostToConnect, "a", false, "Use to allow any ip address (any host) to connect. Default allows ony localhost.")
	flag.IntVar(&port, "port", defaultPort, "Port on which to listen for connections.")

	flag.Parse()

	/* Don't accept any positional command line arguments. flag.NArgs()
	counts only non-flag arguments. */
	if flag.NArg() != 0 {
		/* flag.Usage() isn't in the golang.org documentation,
		but it's right there in the code. It's the same one used when an
		error occurs parsing the flags so it makes sense to use it here as
		well. Hopefully the lack of documentation doesn't mean it's gonna be
		changed it soon. Worst case can always copy that code into a local
		function if it goes away :p
		Currently using go 1.4.1
		https://github.com/golang/go/blob/release-branch.go1.4/src/flag/flag.go#L411
		*/
		flag.Usage()
		os.Exit(exitUsageFailure)
	}

	return
}

func displayServerInfo(listenHost string, listenPort int) {
	visibleTo := listenHost
	if visibleTo == "" {
		visibleTo = "All ip addresses"
	}

	fmt.Printf("%v is running\n\n", projectName)
	fmt.Printf("Visible to: %v\n", visibleTo)
	fmt.Printf("Port: %v\n\n", listenPort)
	fmt.Printf("Hit [ctrl-c] to quit\n")
}
