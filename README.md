# Game of Life

This is a game of life simulation written in Go. I had first written one in Python, but could not remember where I saved it, so I built another in Go.
It takes in initial state in a CSV form. The CSV file can contain fields of 1 or 0, where 1 means alive and 0 means dead. There are some sample files in `/patterns`

## Usage

1. Make sure you have Go installed and set up. You can check [here](https://go.dev/doc/install) to do that.
2. Run the following:

```sh
go run main.go <path-to-csv>
# For example:
go run main.go patterns/gosper_glider.csv
```
