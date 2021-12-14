package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanCmd := "/usr/local/bin/grype -s AllLayers"
	inputFile := fmt.Sprintf("%s/Downloads/ecr_images.txt", os.Getenv("HOME"))
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		image := scanner.Text()
		imgSlice := strings.Split(image, ".")
		account := imgSlice[0]
		region := imgSlice[3]
		logName := strings.ReplaceAll(strings.ReplaceAll(image, "/", "-"), ":", "-") + ".log"
		fmt.Printf("saml2aws login --role arn:aws:iam::%s:role/Capability --skip-prompt --force\n", account)
		fmt.Printf("aws ecr get-login-password --region %s | docker login --username AWS --password-stdin %s.dkr.ecr.%s.amazonaws.com\n", region, account, region)
		fmt.Println(scanCmd, image, "--file", logName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
