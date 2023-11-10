// go-devkit/errorhandler/errorhandler.go
package errorhandler

import (
    "fmt"
    "time"

    "go-devkit/csv/csvmodule"
)

var (
    errorFile = "errors.csv"
)

func HandleError(err error, data ...string) {
    // Prepare the error data
    row := []string{time.Now().Format(time.RFC3339), err.Error()}
    row = append(row, data...)

    // Write the error data to the CSV file
    if err := csvmodule.WriteCSV([][]string{row}, errorFile, 1); err != nil {
        // If an error occurs while writing to the CSV file, log it to the console
        fmt.Println("Failed to write error to CSV file:", err)
    }
}