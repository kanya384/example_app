package useCase

/*
cfg, err := config.InitConfig("")
	if err != nil {
		panic(fmt.Errorf("error reading config: %v", err))
	}

	fmt.Println(cfg)

	file, _ := ioutil.ReadFile("tmp/tmp.jpg")

	storage := cloudStorage.NewStorage(cfg.AWS.AccessKey, cfg.Storage.Bucket, cfg.Storage.Host, cfg.AWS.Region)
	imageResizer := imageResizer.NewImageResizer()

	useCase := uc.New(storage, imageResizer, uc.Options{})

	imageParams := uc.PutImageFileParams{
		Path:      "test",
		Name:      "testName",
		Format:    "jpg",
		Bucket:    "file-storage",
		MaxWidth:  1920,
		MaxHeight: 1080,
		Content:   file,
		Versions: []uc.ImageVersions{
			{
				MaxWidth:  600,
				MaxHeight: 600,
				Suffix:    "-middle",
			},
			{
				MaxWidth:  250,
				MaxHeight: 250,
				Suffix:    "-thumb1",
			},
			{
				MaxWidth:  150,
				MaxHeight: 150,
				Suffix:    "-thumb2",
			},
		},
	}
	err = useCase.PutImageFile(context.Background(), imageParams)
	fmt.Println(err)

	files, err := useCase.ListFilesInPath(context.Background(), uc.ListFilesInPathParams{
		Path:   "test",
		Bucket: "file-storage",
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(*file)
	}
*/
