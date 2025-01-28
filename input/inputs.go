package input

import (
	"bufio"
	"fmt"
	"os"
)

func Read() {
	in := bufio.NewScanner(os.Stdin)
	var data string
	for in.Scan() {
		temp := in.Text()
		data += temp
		if temp == "break" {
			break
		}
	}
	fmt.Println(data)
}
