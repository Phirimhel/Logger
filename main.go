package main

import (
	"fmt"
	"logger/logger"
)

func main() {

	logger, logFileCloseFunc, err := logger.NewLogger("DEBUG")

	if err != nil {
		fmt.Println("ERROR", err)
		panic(err)
	}

	defer logFileCloseFunc()

	logger.Debug("Debug log test")
	logger.Info("Info log test")
	logger.Error("Erorr log test")

}
