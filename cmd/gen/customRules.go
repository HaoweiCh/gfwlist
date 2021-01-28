package gen

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
				if !IsRuleExists(line) {
					buf.WriteString(line + "\n")

					// 补充到列表防止重复
					if strings.HasPrefix(line, "||") {
						includeRules = append(includeRules, line)
					} else if strings.HasPrefix(line, "@@") {
						excludeRules = append(excludeRules, line)
					} else {
						otherRules = append(otherRules, line)
					}
				}
			}
		}
	}
}
