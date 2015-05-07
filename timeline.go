package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jeffbmartinez/cleanexit"
	"github.com/jeffbmartinez/config"
	"github.com/jeffbmartinez/log"

	"github.com/jeffbmartinez/timeline/handler"
	"github.com/jeffbmartinez/timeline/storage/influxdb"
)

const (
	PROJECT_NAME = "timeline"

	EXIT_SUCCESS       = 0
	EXIT_FAILURE       = 1
	EXIT_USAGE_FAILURE = 2 // Same as golang's flag module uses, hardcoded at https://github.com/golang/go/blob/release-branch.go1.4/src/flag/flag.go#L812

	INFLUXDB_HOST = "localhost"
	INFLUXDB_PORT = 8086
	INFLUXDB_NAME = "test_timeline"

	CONFIG_FILENAME = "config.json"
)

type Config struct {
	InfluxDb InfluxDbConfig
}

type InfluxDbConfig struct {
	Username string
	Password string
}

func main() {
	cleanexit.SetUpExitOnCtrlC(getPrintPrettyExitMessageFunc(PROJECT_NAME))

	allowAnyHostToConnect, listenPort := getCommandLineArgs()

	http.HandleFunc("/api/event/single", handler.Single)
	http.HandleFunc("/api/event/start", handler.Start)
	http.HandleFunc("/api/event/stop", handler.Stop)

	listenHost := "localhost"
	if allowAnyHostToConnect {
		listenHost = ""
	}

	var configuration Config
	err := config.ReadSpecific(CONFIG_FILENAME, &configuration)

	if err != nil {
		log.Fatal("Trouble reading config file")
	}

	verifyConfigOrDie(configuration)

	initializeErr := influxdb.Initialize(
		INFLUXDB_HOST,
		INFLUXDB_PORT,
		INFLUXDB_NAME,
		configuration.InfluxDb.Username,
		configuration.InfluxDb.Password,
	)

	if initializeErr != nil {
		fmt.Println("Count not connect to influxDB")
		log.Fatalf("Could not initialize connection to influxDB: %v", initializeErr)
	}

	displayServerInfo(listenHost, listenPort)

	listenAddress := fmt.Sprintf("%v:%v", listenHost, listenPort)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func verifyConfigOrDie(configuration Config) {
	const NUMBER_OF_CHECKS = 10

	problems := make([]string, 0, NUMBER_OF_CHECKS)

	if configuration.InfluxDb.Username == "" {
		problems = append(problems, "Config is missing influxDb.username")
	}

	if configuration.InfluxDb.Password == "" {
		problems = append(problems, "Config is missing influxDb.password")
	}

	if len(problems) > 0 {
		log.Fatal(problems)
	}
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
	const DEFAULT_PORT = 8000

	flag.BoolVar(&allowAnyHostToConnect, "a", false, "Use to allow any ip address (any host) to connect. Default allows ony localhost.")
	flag.IntVar(&port, "port", DEFAULT_PORT, "Port on which to listen for connections.")

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
		os.Exit(EXIT_USAGE_FAILURE)
	}

	return
}

func displayServerInfo(listenHost string, listenPort int) {
	visibleTo := listenHost
	if visibleTo == "" {
		visibleTo = "All ip addresses"
	}

	fmt.Printf("%v is running.\n\n", PROJECT_NAME)
	fmt.Printf("Visible to: %v\n", visibleTo)
	fmt.Printf("Port: %v\n\n", listenPort)
	fmt.Printf("Hit [ctrl-c] to quit\n")
}
