package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var flagFName, flagAction string
var flagTotalActions, flagUser, flagItem int

func parseCommandLine() {
	flag.StringVar(&flagFName, "fname", "default", "Location + filename where the data to be saved. Ex: --fname=/home/herry/default")
	flag.IntVar(&flagTotalActions, "totalDatas", 10000, "Total rows to be generated to file")
	flag.IntVar(&flagUser, "user", 50, "Total user id your platform have")
	flag.IntVar(&flagItem, "item", 1000, "Total item id your platform have")
	flag.StringVar(&flagAction, "action", "atw,watch,comment", "List of Action user did")
	flag.Parse()
}

func customRandom(totalRandom int, totalObject int, allowZero bool, recurringCount int) (result []int) {
	rand.Seed(time.Now().UnixNano() << 1)
	result = rand.Perm(totalRandom)
	rand.Seed(time.Now().UnixNano() >> 1)
	fRandom2 := rand.Perm(totalRandom)

	for j := 0; j < recurringCount; j++ {
		for i := 0; i < totalRandom; i++ {
			result[i] = ((result[i] * fRandom2[i]) % totalObject)
			if !allowZero {
				result[i]++
			}
		}
	}
	return
}

func main() {
	parseCommandLine()

	fName := flagFName + "_" + time.Now().Format("20060102_150405")
	fWrite, err := os.OpenFile(fName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println("[error]", err)
	}
	defer fWrite.Close()
	bufWrite := bufio.NewWriter(fWrite)

	actionTypes := strings.Split(flagAction, ",")
	users := customRandom(flagTotalActions, flagUser, false, 5)
	items := customRandom(flagTotalActions, flagItem, false, 5)
	acts := customRandom(flagTotalActions, len(actionTypes), true, 5)

	for i := 0; i < flagTotalActions; i++ {
		newString := fmt.Sprintln(fmt.Sprintf("%d,%d,%s", users[i], items[i], actionTypes[acts[i]]))
		bufWrite.WriteString(newString)
		bufWrite.Flush()
	}
}
