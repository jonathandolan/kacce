package main

import (
	"bufio"
	//"chess"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	in = strings.TrimSuffix(in, "\n")
	if strings.EqualFold(in, "uci") {
		fmt.Println("id name kacce")
		fmt.Println("id author Jonathan Dolan")
		fmt.Println("uciok")
	}
	in, _ = reader.ReadString('\n')
	in = strings.TrimSuffix(in, "\n")
	if strings.EqualFold(in, "isready") {
		fmt.Println("readyok")
	}
}
