package csvmodule_test

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"

	"github.com/USERNAME/go-devkit/csv/csvmodule"
)

func TestReadCSV(t *testing.T) {
	type args struct {
		filename      string
		maxGoroutines int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Read CSV file successfully",
			args: args{
				filename:      "testdata/test.csv",
				maxGoroutines: 2,
			},
			want: [][]string{
				{"Name", "Age", "City"},
				{"John", "25", "New York"},
				{"Jane", "30", "San Francisco"},
			},
			wantErr: false,
		},
		{
			name: "Read non-existent CSV file",
			args: args{
				filename:      "testdata/non-existent.csv",
				maxGoroutines: 2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := csvmodule.ReadCSV(tt.args.filename, tt.args.maxGoroutines)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSV() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteCSV(t *testing.T) {
	type args struct {
		data          [][]string
		filename      string
		maxGoroutines int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Write CSV file successfully",
			args: args{
				data: [][]string{
					{"Name", "Age", "City"},
					{"John", "25", "New York"},
					{"Jane", "30", "San Francisco"},
				},
				filename:      "testdata/test.csv",
				maxGoroutines: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := csvmodule.WriteCSV(tt.args.data, tt.args.filename, tt.args.maxGoroutines); (err != nil) != tt.wantErr {
				t.Errorf("WriteCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Verify that the file was written correctly
			got, err := ioutil.ReadFile(tt.args.filename)
			if err != nil {
				t.Errorf("WriteCSV() error reading file = %v", err)
			}
			want := "Name,Age,City\nJohn,25,New York\nJane,30,San Francisco\n"
			if string(got) != want {
				t.Errorf("WriteCSV() got = %v, want %v", string(got), want)
			}
			// Clean up the test file
			if err := os.Remove(tt.args.filename); err != nil {
				t.Errorf("WriteCSV() error removing file = %v", err)
			}
		})
	}
}

func TestAppendCSVSequentially(t *testing.T) {
	type args struct {
		data          [][]string
		filename      string
		maxGoroutines int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Append to CSV file successfully",
			args: args{
				data: [][]string{
					{"Name", "Age", "City"},
					{"John", "25", "New York"},
					{"Jane", "30", "San Francisco"},
				},
				filename:      "testdata/test.csv",
				maxGoroutines: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the test file
			if err := ioutil.WriteFile(tt.args.filename, []byte("Name,Age,City\n"), 0644); err != nil {
				t.Errorf("AppendCSVSequentially() error creating file = %v", err)
			}
			if err := csvmodule.AppendCSVSequentially(tt.args.data, tt.args.filename, tt.args.maxGoroutines); (err != nil) != tt.wantErr {
				t.Errorf("AppendCSVSequentially() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Verify that the file was written correctly
			got, err := ioutil.ReadFile(tt.args.filename)
			if err != nil {
				t.Errorf("AppendCSVSequentially() error reading file = %v", err)
			}
			want := "Name,Age,City\nJohn,25,New York\nJane,30,San Francisco\n"
			if string(got) != want {
				t.Errorf("AppendCSVSequentially() got = %v, want %v", string(got), want)
			}
			// Clean up the test file
			if err := os.Remove(tt.args.filename); err != nil {
				t.Errorf("AppendCSVSequentially() error removing file = %v", err)
			}
		})
	}
}

func TestAppendCSVConcurrently(t *testing.T) {
	type args struct {
		data          [][]string
		filename      string
		maxGoroutines int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Append to CSV file successfully",
			args: args{
				data: [][]string{
					{"Name", "Age", "City"},
					{"John", "25", "New York"},
					{"Jane", "30", "San Francisco"},
				},
				filename:      "testdata/test.csv",
				maxGoroutines: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the test file
			if err := ioutil.WriteFile(tt.args.filename, []byte("Name,Age,City\n"), 0644); err != nil {
				t.Errorf("AppendCSVConcurrently() error creating file = %v", err)
			}
			if err := csvmodule.AppendCSVConcurrently(tt.args.data, tt.args.filename, tt.args.maxGoroutines); (err != nil) != tt.wantErr {
				t.Errorf("AppendCSVConcurrently() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Verify that the file was written correctly
			got, err := ioutil.ReadFile(tt.args.filename)
			if err != nil {
				t.Errorf("AppendCSVConcurrently() error reading file = %v", err)
			}
			want := "Name,Age,City\nJohn,25,New York\nJane,30,San Francisco\n"
			if string(got) != want {
				t.Errorf("AppendCSVConcurrently() got = %v, want %v", string(got), want)
			}
			// Clean up the test file
			if err := os.Remove(tt.args.filename); err != nil {
				t.Errorf("AppendCSVConcurrently() error removing file = %v", err)
			}
		})
	}
}

