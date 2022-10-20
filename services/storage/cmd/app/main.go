package main

import (
	"context"
	"fmt"
	"os"
	"storage/internal/config"
	"storage/pkg/cloudStorage"
)

func main() {

	cfg, err := config.InitConfig("")
	if err != nil {
		panic(fmt.Errorf("error reading config: %v", err))
	}

	fmt.Println(cfg)

	storage := cloudStorage.NewStorage(cfg.AWS.AccessKey, cfg.Storage.Bucket, cfg.Storage.Host, cfg.AWS.Region)

	file, _ := os.Open("cmd/app/main.go")
	defer file.Close()
	storage.PutObjectToStorage(context.Background(), "cmd/app", "main.go", file)
	files, err := storage.ListObjects(context.Background(), "")
	for _, file := range files {
		fmt.Println(*file)
	}

}
