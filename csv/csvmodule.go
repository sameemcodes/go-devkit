// csvmodule/csvmodule.go
package csvmodule

import (
	"encoding/csv"
	"io"
	"os"
	"sync"
)

// ReadCSV reads a CSV file concurrently
func ReadCSV(filename string, maxGoroutines int) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data := [][]string{}

	var wg sync.WaitGroup
	goroutineSem := make(chan struct{}, maxGoroutines)

	errors := make(chan error, 1)
	done := make(chan bool, 1)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		goroutineSem <- struct{}{}
		wg.Add(1)

		go func(record []string) {
			defer wg.Done()
			data = append(data, record)
			<-goroutineSem
		}(record)
	}

	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		return data, nil
	case err := <-errors:
		return nil, err
	}
}

// WriteCSV writes data to a CSV file concurrently
func WriteCSV(data [][]string, filename string, maxGoroutines int) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var wg sync.WaitGroup
	goroutineSem := make(chan struct{}, maxGoroutines)

	errors := make(chan error, 1)
	done := make(chan bool, 1)

	go func() {
		for _, row := range data {
			goroutineSem <- struct{}{}
			wg.Add(1)

			go func(row []string) {
				defer wg.Done()
				if err := writer.Write(row); err != nil {
					select {
					case errors <- err:
					default:
					}
				}
				<-goroutineSem
			}(row)
		}
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		return nil
	case err := <-errors:
		return err
	}
}

// AppendCSVSequentially appends data to a CSV file sequentially using multiple goroutines
func AppendCSVSequentially(data [][]string, filename string, maxGoroutines int) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var wg sync.WaitGroup

	for _, row := range data {
		wg.Add(1)

		go func(row []string) {
			defer wg.Done()
			if err := writer.Write(row); err != nil {
				panic(err)
			}
		}(row)
	}

	wg.Wait()

	return nil
}



	