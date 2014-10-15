func main() {
	results_filename := fmt.Sprintf("%s_out.csv", in_file)
	results_file, err := os.OpenFile(results_filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer results_file.Close()

	image_urls := imageUrls(in_file)

	brit_batch := TaskConfig{
		Title:       "Tag the appropriate images",
		Description: description,
		Write:       writers.Csv(results_file),
		Tasks:       image_urls,
	}

	log.Printf("Loaded %d images and starting\n", len(image_urls))

	go EnableHtmlServer(IndexConfig{
		Name: brit_batch.Title,
	Url: "
	})
	NewBatch(brit_batch).Run()
}
