package gen

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	customRuleUseCase "gfwlist/usecases/customrule"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Gen(cmd *cobra.Command, args []string) {
	fmt.Println("gen called")

	cursor, err := os.Open("./gfwlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	cursorWrite, err := os.Create("./gfwlist.haowei.txt")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(cursorWrite)

	all, err := ioutil.ReadAll(cursor)
	if err != nil {
		log.Fatal(err)
	}

	out := make([]byte, base64.StdEncoding.EncodedLen(len(all)))
	log.Println(base64.StdEncoding.Decode(out, all))
	scanner := bufio.NewScanner(bytes.NewReader(out))
	var (
		includeRules []string // 强制走代理规则条数
		excludeRules []string // 强制不走代理规则条数
		otherRules   []string
	)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "---EOF---") {
			customRuleUseCase.Add(writer)
		} else {
			_, _ = writer.WriteString(line + "\n")

		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "!") {
			//log.Println(line)
		} else {
			if strings.Contains(line, "AutoProxy") {
				log.Println(line)
			} else if strings.HasPrefix(line, "||") {
				includeRules = append(includeRules, line)
			} else if strings.HasPrefix(line, "@@") {
				excludeRules = append(excludeRules, line)
			} else {
				otherRules = append(otherRules, line)
			}
		}

	}

	writer.Flush()

	j, _ := json.MarshalIndent(map[string]interface{}{
		"includeCount": len(includeRules),
		"excludeCount": len(excludeRules),
		"otherCount":   len(otherRules),
	}, "", " ")
	log.Printf("%s\n", j)
}
