package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var token string

func main() {
	flag.StringVar(&token, "token", "token", "The token to operate the influxdb")
	flag.Parse()

	bucket := "stats-cpu"
	org := "mapper-cpu"
	// token := "7UTkMfW4uuFdIJHDBapqggz8mSS_UWbqm6pmDuk3HGpT3zV1Gaf2dWnFIzveOLDO2Ec86jbFmSpoZVPNtXvMzA=="
	// Store the URL of your InfluxDB instance
	url := "http://127.0.0.1:8086"
	// Create new client with default option for server url authenticate by token
	client := influxdb2.NewClient(url, token)

	// User blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)

	for {
		// Create point using full params constructor
		cpustats := float32(60 + rand.Intn(20))
		p := influxdb2.NewPoint("cpu",
			map[string]string{"cpu-stats": "cpu-stats"},
			map[string]interface{}{"cpu-stats": cpustats},
			time.Now())
		// Write point immediately
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("successfully sent %v to cloud", cpustats)
		}
		time.Sleep(5 * time.Second)
	}
	// Ensures background processes finishes
	client.Close()
}
