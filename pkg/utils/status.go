package utils

import "os"

// Costanti per staus code di uscita

const (
    // WRONG_SYNTAX = 
    // INVALID_CONFIGURATION =
    // GENERIC_ERROR DURING EXECUTION
)

type ErrorWithStatus struct {

}

// Add functions to exit with status and custom message to
func ExitWithStatus() {
    os.Exit(0)
}
