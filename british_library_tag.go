package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	. "github.com/abhiyerra/workmachine/crowdflow"
	writers "github.com/abhiyerra/workmachine/crowdflow/writers"
	"log"
	"os"
)

type ImageTagging struct {
	ImageUrl             InputField  `work_desc:"Use this image to fill the information below." work_id:"image_url" work_type:"image"`
	Tags                 OutputField `work_desc:"List all the relevent tags separated by a comma for the image. Ex. trees, castle, person" work_id:"tags"`
	TextInImage          OutputField `work_desc:"Put any caption that appears of the image here. Put one item per line if there are multiple." work_id:"text_in_image" work_type:"long_text"`
	IsCorrectOrientation OutputField `work_desc:"Is the image in the correct orientation?" work_id:"is_correct_orientation" work_type:"checkbox"`
	IsLandscape          OutputField `work_desc:"Is the image of a landscape (a non urban setting)?" work_id:"is_landscape" work_type:"checkbox"`
	IsPattern            OutputField `work_desc:"Is the image of a pattern?" work_id:"is_pattern" work_type:"checkbox"`
	IsPerson             OutputField `work_desc:"Does the image contain people?" work_id:"is_person" work_type:"checkbox"`
	TraditionalClothing  OutputField `work_desc:"If the image has people are they wearing traditional clothes?" work_id:"traditional_clothing" work_type:"checkbox"`
	IsMap                OutputField `work_desc:"Is the image a map?" work_id:"is_map" work_type:"checkbox"`
	IsDiagram            OutputField `work_desc:"Is the image a diagram?" work_id:"is_diagram" work_type:"checkbox"`
}

func imageUrls(in_file string) (images []ImageTagging) {
	file, err := os.Open(in_file)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	for _, i := range records {
		//		fmt.Printf("%s\n", i[1])
		images = append(images, ImageTagging{ImageUrl: InputField(i[1])})
	}

	if err != nil {
		panic(err)
	}

	return
}

func main() {
	var in_file string
	flag.StringVar(&in_file, "in_file", "", "input file")
	flag.Parse()

	if in_file == "" {
		fmt.Println("No in file")
		os.Exit(1)
	}

	results_filename := fmt.Sprintf("%s_out.csv", in_file)
	results_file, err := os.OpenFile(results_filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer results_file.Close()

	image_urls := imageUrls(in_file)

	description := `
Look at the image and fill out the appropriate fields. We want to be able to tag all the images correctly. Fill out any appropriate tag that you see.
 <a href="https://github.com/abhiyerra/britishlibrary/wiki/Instructions-&-FAQ">Here are further Instructions and FAQ</a>`

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
