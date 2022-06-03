package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	splitted := strings.Split(text, ".")
	minutes := strings.TrimSpace(splitted[0])[:2]
	seconds := strings.TrimSpace(splitted[1])[:2]
	converted := fmt.Sprintf("%vm%vs", minutes, seconds)
	// fmt.Printf("%v\n minutes: %v, seconds: %v, converted: %v\n", splitted, minutes, seconds, converted)
	duration, _ := time.ParseDuration(converted)

	defaultTime := time.Unix(now, 0)
	fmt.Printf("%v\n", defaultTime.UTC().Add(duration).Format(time.UnixDate))
}
