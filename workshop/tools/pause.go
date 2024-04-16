package tools

import (
	"bufio"
	"fmt"
	"os"
)

func Pause() {
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
