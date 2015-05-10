// Copyright 2015 Jeff Martinez. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE.txt file
// or at http://opensource.org/licenses/MIT

package cleanexit

import (
	"os"
	"os/signal"
)

// Standard command line tool code for 'successful termination'.
const EXIT_SUCCESS = 0

/*
	Sets up a separate goroutine to listen for ctrl-c.
	As soon as a ctrl-c is entered in the terminal, the program
	is terminated. An program exit code of 0 is returned.

	Use SetUpExitOnCtrlC to allow for a custom function to run
	right before program termination.
*/
func SetUpSimpleExitOnCtrlC() {
	SetUpExitOnCtrlC(func() {})
}

/*
	Similar to SetUpSimpleExitOnCtrlC except you can pass in a
	function to run right before the program terminates.

	Can be used for clean up, or to print a nice exit message. I've
	used this to hide the "^C" that is printed to a terminal when a
	user hits ctrl-c. You can do this by printing a couple \b
	characters followed by spaces to stdout. Like this:

	fmt.Printf("\b\b  \n")

	The \b characters are the equivalent of hitting the back arrow
	key. Wrap that in a function, pass it in to SetUpExitOnCtrlC and
	you're set.

	You can use your own os.Exit(...) within the cleanup function
	to signify non-successful exit if necessary.
*/
func SetUpExitOnCtrlC(cleanup func()) {
	const NUM_PARALLEL_SIGNALS_TO_PROCESS = 1

	killChannel := make(chan os.Signal, NUM_PARALLEL_SIGNALS_TO_PROCESS)
	signal.Notify(killChannel, os.Interrupt, os.Kill)

	go func() {
		<-killChannel
		cleanup()
		os.Exit(EXIT_SUCCESS)
	}()
}
