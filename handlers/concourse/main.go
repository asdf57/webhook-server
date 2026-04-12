package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func getEnvOrPanic(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is required but not set", name))
	}
	return value
}

func triggerConcourseResourceCheck(ciURL, pipelineName, pipelineResource, webhookToken, webhookData string) error {
	url := fmt.Sprintf("https://%s/api/v1/teams/main/pipelines/%s/resources/%s/check/webhook?webhook_token=%s", ciURL, pipelineName, pipelineResource, webhookToken)
	fmt.Printf("Triggering Concourse pipeline with URL: %s\n", url)

	// Make a POST request to the Concourse API with the webhook data
	// You can use the "net/http" package to make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(webhookData)))
	if err != nil {
		return fmt.Errorf("error triggering Concourse pipeline: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to trigger Concourse pipeline, status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("Usage: concourse <config-file>")
		os.Exit(1)
	}

	fmt.Printf("Concourse received config file: %v\n", argsWithoutProg)

	// Grab the json webhook data from os.Args[1]
	pipelineName := os.Args[1]
	pipelineResource := os.Args[2]
	webhookData := os.Args[3]

	fmt.Printf("Pipeline Name: %s\n", pipelineName)
	fmt.Printf("Pipeline Resource: %s\n", pipelineResource)
	fmt.Printf("Webhook Data: %s\n", webhookData)

	// https://ci.ryuugu.dev/api/v1/teams/main/pipelines/render-standups/resources/inventory-trigger/check/webhook?webhook_token=some-secret-token

	// Need CI url, team name, pipeline name, resource name, and webhook token to trigger the pipeline
	ciURL := getEnvOrPanic("CI_URL")
	webhookToken := getEnvOrPanic("WEBHOOK_TOKEN")

	url := fmt.Sprintf("https://%s/api/v1/teams/main/pipelines/%s/resources/%s/check/webhook?webhook_token=%s", ciURL, pipelineName, pipelineResource, webhookToken)
	fmt.Printf("Triggering Concourse pipeline with URL: %s\n", url)

	// Make a POST request to the Concourse API with the webhook data
	if err := triggerConcourseResourceCheck(ciURL, pipelineName, pipelineResource, webhookToken, webhookData); err != nil {
		fmt.Printf("Error triggering Concourse pipeline: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully triggered Concourse pipeline")
}
