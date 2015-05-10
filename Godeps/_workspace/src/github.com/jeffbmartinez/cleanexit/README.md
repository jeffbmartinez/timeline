# cleanexit

Package to help cleanly terminate a golang program via ctrl-c.

## Functions

### SetUpSimpleExitOnCtrlC()

Sets up a separate goroutine to listen for ctrl-c.
As soon as a ctrl-c is entered in the terminal, the program
is terminated. An program exit code of 0 is returned.

Use `SetUpExitOnCtrlC` to allow for a custom function to run
right before program termination.

### SetUpExitOnCtrlC(cleanup func())

Similar to `SetUpSimpleExitOnCtrlC` except you can pass in a
function to run right before the program terminates.

Can be used for clean up, or to print a nice exit message. I've
used this to hide the ^C that is printed to a terminal when a
user hits ctrl-c. You can do this by printing a couple \b
characters followed by spaces to stdout:

    fmt.Printf("\b\b  \n")

The \b characters are the equivalent of hitting the back arrow
key. Wrap that in a function, pass it in to SetUpExitOnCtrlC and
you're set.

You can use your own call to os.Exit(...) within the cleanup function
to signify non-successful exit if necessary, otherwise an exit code
of 0 (zero) will be used.

### Example

    import (
        "fmt"
        "os"
        
        "github.com/jeffbmartinez/cleanexit"
    )
    
    func main() {
        cleanexit.SetUpExitOnCtrlC(myExitFunc)
        
        // ...
    }
    
    func myExitFunc() {
        // clean up connections, etc
        
        if (/* something went wrong */) {
            os.Exit(FAILURE_CODE)
        }
        
        fmt.Printf("\b\b  \nGoodbye!\n")
    }
