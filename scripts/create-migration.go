package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func create_migration(tableName string) {
	var sb strings.Builder
	timestamp := time.Now().Format("20060102150405.003059_")
	regexString := regexp.MustCompile(`^(.*?)\.(.*)$`)
	replaceString := "${1}$2"
	sb.WriteString("migrations/scripts/")
	sb.WriteString(regexString.ReplaceAllString(timestamp, replaceString))
	sb.WriteString(tableName)
	sb.WriteString(".sql")
	fileName := sb.String()
	emptyFile, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		log.Println(err)
	}
	log.Println("Created SQL File:", fileName)
	emptyFile.Close()
}

func main() {
	if len(os.Args) <= 1 {
		log.Panic("Please specify a migration name as first argument")
	}
	fileName := strings.ToLower(strings.TrimSpace(os.Args[1]))
	create_migration(fileName)
}
