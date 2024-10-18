package main

import (
	"log/slog"
	"os"
)

var jsonHandler = slog.NewJSONHandler(os.Stdout, nil)
var l = slog.New(jsonHandler)
