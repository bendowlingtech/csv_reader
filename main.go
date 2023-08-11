package csv_reader

import (
	"fmt"
	"encoding/csv"
	"flag"
	"os"
)

func main() {

	file_path := os.Args[1]
	line_amount := os.Args[2]
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}


}
