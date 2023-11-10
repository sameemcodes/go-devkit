// errorhandler/errorhandler_test.go
package errorhandler

import (
    "errors"
    "os"
    "testing"

    "go-devkit/csv/csvmodule"
)

func TestHandleError(t *testing.T) {
    // Create a test error
    err := errors.New("test error")

    // Handle the error
    HandleError(err, "additional", "data")

    // Read the error file
    data, readErr := csvmodule.ReadCSV(errorFile)
    if readErr != nil {
        t.Fatal(readErr)
    }

    // Check that the error data was written correctly
    if len(data) != 1 || data[0][1] != err.Error() || data[0][2] != "additional" || data[0][3] != "data" {
        t.Errorf("HandleError() = %v, want %v", data[0], []string{"", err.Error(), "additional", "data"})
    }

    // Clean up the error file
    os.Remove(errorFile)
}