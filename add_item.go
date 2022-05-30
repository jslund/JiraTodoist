package function

import (
	"context"
	"os"
	"strconv"

	todoist "github.com/sachaos/todoist/lib"
)

func AddItem(title string, priority int) error {
	config := &todoist.Config{AccessToken: os.Getenv("ACCESS_TOKEN"), DebugMode: false, Color: false}

	client := todoist.NewClient(config)

	item := todoist.Item{}

	item.Content = title
	item.Priority = priority

	if projectId, err := strconv.Atoi(os.Getenv("PROJECT_ID")); err == nil {
		item.ProjectID = projectId
	}
	if err := client.AddItem(context.Background(), item); err != nil {
		return err
	}
	return client.Sync(context.Background())
}
