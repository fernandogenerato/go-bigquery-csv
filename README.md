# BigQuery CSV Insertion API

This repository contains a simple API built with Go and Fiber to insert data into Google BigQuery from uploaded files.

## Requirements

- Go (>=1.14)
- Google Cloud Platform Account with BigQuery API enabled

## Setup

Clone the repository:

```bash
git clone https://github.com/your_username/your_repository.git
```
Install dependencies:
```bash
go mod tidy
```
Set up Google Cloud credentials:
```
Place your Google Cloud service account key file (credentials_file.json) in the root directory of the project.

Set your Google Cloud project ID:

Replace project_id in main.go with your Google Cloud project ID.
```

Run the application:
```bash
go run main.go
```
## Usage
API provides a single endpoint for inserting data from uploaded files:
```bash
go run main.go
```
### Route
`POST /api`

### Request Parameters
`dataset`: The BigQuery dataset where the data will be inserted.

`table`: The name of the table in the specified dataset.

`file`: File location.
### Example Request
```bash
curl -X POST \
  -F "dataset=my_dataset" \
  -F "table=my_table" \
  -F "file=@/path/to/your/file.csv" \
  http://localhost:7171/api
```

Make sure to replace placeholders like `your_username`, `your_repository`, `project_id`, `my_dataset`, `my_table`, and adjust the file paths as needed.
