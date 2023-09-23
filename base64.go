package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	b64 "encoding/base64"
)

var (
	col = flag.Int("w", 64, "Wrap lines after N columns")
	dec = flag.Bool("d", false, "Decode instead of Encode")
	pad = flag.Bool("n", false, "No padding")
)

func main() {
	flag.Parse()

	if *col == 0 && len(flag.Args()) > 0 {
		inputFile := flag.Arg(0)

		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Println("Error reading the file:", err)
			os.Exit(1)
		}

		inputData := string(data)
		inputData = strings.TrimSuffix(inputData, "\r\n")
		inputData = strings.TrimSuffix(inputData, "\n")

		if *dec == false && *pad == false {
			sEnc := b64.StdEncoding.EncodeToString([]byte(inputData))
			fmt.Println(sEnc)
		} else if *dec && *pad == false {
			sDec, _ := b64.StdEncoding.DecodeString(inputData)
			os.Stdout.Write(sDec)
		}

		if *dec == false && *pad == true {
			sEnc := b64.StdEncoding.WithPadding(-1).EncodeToString([]byte(inputData))
			fmt.Println(sEnc)
		} else if *dec && *pad == true {
			sDec, _ := b64.StdEncoding.WithPadding(-1).DecodeString(inputData)
			os.Stdout.Write(sDec)
		}
	} else {
		var inputData string

		if len(flag.Args()) == 0 {
			data, _ := ioutil.ReadAll(os.Stdin)
			inputData = string(data)
		} else {
			inputFile := flag.Arg(0)

			data, err := ioutil.ReadFile(inputFile)
			if err != nil {
				fmt.Println("Error reading the file:", err)
				os.Exit(1)
			}
			inputData = string(data)
		}

		inputData = strings.TrimSuffix(inputData, "\r\n")
		inputData = strings.TrimSuffix(inputData, "\n")

		if *col != 0 {
			if *dec == false && *pad == false {
				sEnc := b64.StdEncoding.EncodeToString([]byte(inputData))
				for _, chunk := range split(sEnc, *col) {
					fmt.Println(chunk)
				}
			} else if *dec && *pad == false {
				sDec, _ := b64.StdEncoding.DecodeString(inputData)
				os.Stdout.Write(sDec)
			}

			if *dec == false && *pad == true {
				sEnc := b64.StdEncoding.WithPadding(-1).EncodeToString([]byte(inputData))
				for _, chunk := range split(sEnc, *col) {
					fmt.Println(chunk)
				}
			} else if *dec && *pad == true {
				sDec, _ := b64.StdEncoding.WithPadding(-1).DecodeString(inputData)
				os.Stdout.Write(sDec)
			}
		} else {
			if *dec == false && *pad == false {
				sEnc := b64.StdEncoding.EncodeToString([]byte(inputData))
				fmt.Println(sEnc)
			} else if *dec && *pad == false {
				sDec, _ := b64.StdEncoding.DecodeString(inputData)
				os.Stdout.Write(sDec)
			}

			if *dec == false && *pad == true {
				sEnc := b64.StdEncoding.WithPadding(-1).EncodeToString([]byte(inputData))
				fmt.Println(sEnc)
			} else if *dec && *pad == true {
				sDec, _ := b64.StdEncoding.WithPadding(-1).DecodeString(inputData)
				os.Stdout.Write(sDec)
			}
		}
	}
}

func split(s string, size int) []string {
	ss := make([]string, 0, len(s)/size+1)
	for len(s) > 0 {
		if len(s) < size {
			size = len(s)
		}
		ss, s = append(ss, s[:size]), s[size:]
	}
	return ss
}
