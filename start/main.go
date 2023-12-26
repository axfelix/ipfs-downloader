package main

import (
	"context"
	"flag"
	"fmt"
	"ipfs-downloader/app"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {

	urlPtr := flag.String("url", "", "URL to download")
	dirPtr := flag.String("dir", "", "Destination directory")
	flag.Parse()

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "download-workflow",
		TaskQueue: app.DownloadTaskQueue,
	}

	// Start the Workflow
	url := *urlPtr
	if url == "" {
		log.Fatalln("no URL provided, specify one with --url", err)
	}
	dir := *dirPtr
	if dir == "" {
		log.Fatalln("no destination provided, specify one with --dir", err)
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, app.DownloadWorkflow, url, dir)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Get the results
	var size string
	err = we.Get(context.Background(), &size)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(size, we.GetID(), we.GetRunID())
}

func printResults(size string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", size)
}
