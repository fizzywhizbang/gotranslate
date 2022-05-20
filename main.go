package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Response struct {
	Data struct {
		Translations []Translation
	}
}
type Translation struct {
	TranslatedText         string
	DetectedSourceLanguage string
}

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

	var response Response
	body := GetBody("https://translation.googleapis.com/language/translate/v2?" + urlVals.Encode())

	json.Unmarshal(body, &response)
	fmt.Println("Word to translate:", *q)
	fmt.Println("Detected Language:", returnLang(*source))
	fmt.Println("In ", returnLang(*target), ":", response.Data.Translations[0].TranslatedText)

}

func returnLang(lang string) string {
	switch lang {
	case "af":
		return "Afrikaans"
	case "ga":
		return "Irish"
	case "sq":
		return "Albanian"
	case "it":
		return "Italian"
	case "ar":
		return "Arabic"
	case "ja":
		return "Japanese"
	case "az":
		return "Azerbaijani"
	case "kn":
		return "Kannada"
	case "eu":
		return "Basque"
	case "ko":
		return "Korean"
	case "bn":
		return "Bengali"
	case "la":
		return "Latin"
	case "be":
		return "Belarusian"
	case "lv":
		return "Latvian"
	case "bg":
		return "Bulgarian"
	case "lt":
		return "Lithuanian"
	case "ca":
		return "Catalan"
	case "mk":
		return "Macedonian"
	case "zh-CN":
		return "Chinese Simplified"
	case "ms":
		return "Malay"
	case "zh-TW":
		return "Chinese Traditional"
	case "mt":
		return "Maltese"
	case "hr":
		return "Croatian"
	case "no":
		return "Norwegian"
	case "cs":
		return "Czech"
	case "fa":
		return "Persian"
	case "da":
		return "Danish"
	case "pl":
		return "Polish"
	case "nl":
		return "Dutch"
	case "pt":
		return "Portuguese"
	case "en":
		return "English"
	case "ro":
		return "Romanian"
	case "eo":
		return "Esperanto"
	case "ru":
		return "Russian"
	case "et":
		return "Estonian"
	case "sr":
		return "Serbian"
	case "tl":
		return "Filipino"
	case "sk":
		return "Slovak"
	case "fi":
		return "Finnish"
	case "sl":
		return "Slovenian"
	case "fr":
		return "French"
	case "es":
		return "Spanish"
	case "gl":
		return "Galician"
	case "sw":
		return "Swahili"
	case "ka":
		return "Georgian"
	case "sv":
		return "Swedish"
	case "de":
		return "German"
	case "ta":
		return "Tamil"
	case "el":
		return "Greek"
	case "te":
		return "Telugu"
	case "gu":
		return "Gujarati"
	case "th":
		return "Thai"
	case "ht":
		return "Haitian Creole"
	case "tr":
		return "Turkish"
	case "iw":
		return "Hebrew"
	case "uk":
		return "Ukrainian"
	case "hi":
		return "Hindi"
	case "ur":
		return "Urdu"
	case "hu":
		return "Hungarian"
	case "vi":
		return "Vietnamese"
	case "is":
		return "Icelandic"
	case "cy":
		return "Welsh"
	case "id":
		return "Indonesian"
	case "yi":
		return "Yiddish"
	}

	return "French"
}

func GetBody(url string) []byte {
	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"User-Agent":   []string{"Mozilla/5.0"},
		"Content-Type": []string{"application/json"},
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
