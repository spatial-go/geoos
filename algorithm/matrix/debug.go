package matrix

import (
	"fmt"
	"log"
	"os"

	"github.com/spatial-go/geoos"
)

const dir = "/debugtools/debug_data/"

// WriteMatrix Write the matrix object to a file in matrix format.
func WriteMatrix(filename string, m Steric) {
	file, err := os.Create(geoos.EnvPath() + dir + filename)
	defer file.Close()

	if err != nil {
		log.Println(err)
	} else {
		if _, err := fmt.Fprintf(file, "%v", m); err != nil {
			log.Println(err)
		}
	}
	//fmt.Printf("%v", m)
}
