package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	diskName   string
	numRows    int
	numColumns int
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "[Error] Wrong number of arguments\n")
		os.Exit(1)
	}

	numRows, err := strconv.Atoi(os.Args[2])
	if err != nil || numRows < 2 || numRows > 32 {
		fmt.Fprintf(os.Stderr, "[Error] Rows must be a number between 2 and 32\n")
		os.Exit(1)
	}

	numColumns, err := strconv.Atoi(os.Args[3])
	if err != nil || numColumns < 8 || numColumns > 64 || numColumns%2 != 0 {
		fmt.Fprintf(os.Stderr, "[Error] Columns must be an even number between 8 and 64\n")
		os.Exit(1)
	}

	diskName := os.Args[1]
	if len(diskName)*2 > numColumns-8 {
		fmt.Fprintf(os.Stderr, "[Error] Disk name too long for row length\n")
		os.Exit(1)
	}

	config := &Config{diskName, numRows, numColumns}

	run(config)
}

func run(config *Config) {
	// fmt.Printf("Name: %s, Rows: %d, Columns: %d\n", config.diskName, config.numRows, config.numColumns)

	fmt.Print("XX:")
	for i := range config.numColumns / 16 {
		if i == 1 {
			fmt.Print(" ")
		}

		if i != 0 {
			fmt.Printf("               %d", i)
		}
	}
	fmt.Println()

	fmt.Print("XX:")
	for i := range config.numColumns {
		fmt.Printf("%X", i%16)
	}
	fmt.Println()

	for i := range config.numRows {
		if i == 0 {
			diskCode := strings.ToUpper(hex.EncodeToString([]byte(config.diskName)))
			fmt.Printf("00:0010000%s", diskCode)
			for range config.numColumns - 7 - len(diskCode) {
				fmt.Print("0")
			}
			fmt.Println()
		} else if i == config.numRows-1 {
			fmt.Printf("%02X:100", i)
			for range config.numColumns - 3 {
				fmt.Print("0")
			}
			fmt.Println()
		} else {
			fmt.Printf("%02X:1%02X", i, i+1)
			for range config.numColumns - 3 {
				fmt.Print("0")
			}
			fmt.Println()
		}
	}
}
