package main

import (
    "context"
    "fmt"
    _ "github.com/influxdata/influxdb-client-go/v2" 
)

func main() {
	// Create client
	client := influxdb2.NewClient("http://localhost:8086", "hIRBzHqdlh1D6VRd-CV7hH07YLDXdwwvM2GPb3Cd-h5OLgitJZbFOKjC5DaKERmDOXI3VUiJJDbUs_957vwh9Q==")
	// Get query client
	queryAPI := client.QueryAPI("demo")
	// Get QueryTableResult
	result, err := queryAPI.Query(context.Background(), `from(bucket:"golang")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "cpu")`)
	if err != nil {
		panic(err)
	}
	// Iterate over query response
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// Access data
		fmt.Printf("value: %v\n", result.Record().Value())
	}
	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %s\n", result.Err().Error())
	}
	// Ensures background processes finishes
	client.Close()
}
