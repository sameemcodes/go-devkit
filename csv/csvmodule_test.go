// csvmodule/csvmodule_test.go
package csvmodule

import (
	"reflect"
	"testing"
)

func TestWriteAndReadCSV(t *testing.T) {
	// Generate some test data
	data := [][]string{{"Name", "Age"}, {"Alice", "30"}, {"Bob", "25"}}

	// Test with different numbers of goroutines
	for _, maxGoroutines := range []int{1, 2, 5, 10} {
		filename := "test.csv"

		// Write the data to a CSV file
		if err := WriteCSV(data, filename, maxGoroutines); err != nil {
			t.Fatalf("Failed to write CSV: %v", err)
		}

		// Read the data back from the file
		readData, err := ReadCSV(filename, maxGoroutines)
		if err != nil {
			t.Fatalf("Failed to read CSV: %v", err)
		}

		// Verify the data
		if !reflect.DeepEqual(readData, data) {
			t.Errorf("Unexpected data: got %v, want %v", readData, data)
		}
	}
}