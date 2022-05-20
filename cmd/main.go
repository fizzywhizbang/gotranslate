package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	trans "github.com/fizzywhizbang/gotranslate"
)

func main() {

	//get key from file
	key, err := os.ReadFile("key.txt")
	if err != nil {
		log.Panic(err)
	}
	//set variables
	var (
		target = flag.String("t", "fr", "Target language (two-letter code)")
		q      = flag.String("p", "cheese", "Word or phrase to translate")
		source = flag.String("s", "en", "Source language")
		model  = "nmt"
		format = "text"
	)
	flag.Parse()

	urlVals := make(url.Values)
	urlVals.Set("key", string(key))
	urlVals.Set("target", *target)
	urlVals.Set("format", format)
	urlVals.Set("q", *q)
	urlVals.Set("source", *source)
	urlVals.Set("model", model)

	var response trans.Response
	body := trans.GetBody("https://translation.googleapis.com/language/translate/v2?" + urlVals.Encode())

	json.Unmarshal(body, &response)
	fmt.Println("Word to translate:", *q)
	fmt.Println("Detected Language:", trans.ReturnLang(*source))
	fmt.Println("In ", trans.ReturnLang(*target), ":", response.Data.Translations[0].TranslatedText)

}
