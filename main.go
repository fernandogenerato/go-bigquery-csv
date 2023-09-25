package main

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/gofiber/fiber/v2"
)

const (
	projectID       = "project_id"
	credentialsEnv  = "GOOGLE_APPLICATION_CREDENTIALS"
	credentialsFile = "credentials_file.json"
)

func init() {
	// key file is in the same directory as main.go
	os.Setenv(credentialsEnv, credentialsFile)
}

func main() {
	app := fiber.New()
	app.Group("/api", insertFromFile) // /api
	log.Fatal(app.Listen(":7171"))
}

func newBigqueryClient(ctx context.Context) (*bigquery.Client, error) {
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func insertFromFile(c *fiber.Ctx) error {
	ctx := c.Context()
	client, err := newBigqueryClient(ctx)
	if err != nil {
		return c.JSON(err)
	}
	defer client.Close()

	params := new(struct {
		Dataset string `form:"dataset"`
		Table   string `form:"table"`
	})
	if err = c.BodyParser(params); err != nil {
		return c.JSON(err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(err)
	}

	for _, files := range form.File {
		for _, file := range files {
			source, err := newReaderSource(file)
			if err != nil {
				return c.JSON(err)
			}
			source.AutoDetect = true
			source.SkipLeadingRows = 1
			loader := client.Dataset(params.Dataset).Table(params.Table).LoaderFrom(source)
			job, err := loader.Run(ctx)
			if err != nil {
				return c.JSON(err)
			}
			status, err := job.Wait(ctx)
			if err != nil {
				return c.JSON(err)
			}
			if err := status.Err(); err != nil {
				return c.JSON(err)
			}
		}
	}
	return nil
}

func newReaderSource(file *multipart.FileHeader) (*bigquery.ReaderSource, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	return bigquery.NewReaderSource(io.Reader(f)), nil
}
