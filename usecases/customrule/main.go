package customRuleUseCase

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Add(buf *bufio.Writer) {
	buf.WriteString("!###############Customlist Start################\n")
	defer buf.WriteString("!################Customlist End#################\n!---------------------EOF-----------------------\n")

	matches, err := filepath.Glob("./rules/*.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, match := range matches {
		cursor, err := os.Open(match)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(cursor)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) > 0 {
				buf.WriteString(line + "\n")
			}
		}
	}
}
