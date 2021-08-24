package main // This is an executable package

import ( // Importing functionality from the os library package
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() { // Executable entry point, command line parameters parsed later
	path := flag.String("path", "myapp.log", "The path to the log that should be parsed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, CRITIAL")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil { // Error types are pointer values
		log.Fatal(err)
	}
	defer f.Close() // Close the file when scope is closed

	r := bufio.NewReader(f) // Reads some type of input stream, file in this case
	for {
		s, err := r.ReadString('\n') // Read until a newline character
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
