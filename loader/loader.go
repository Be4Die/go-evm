// loader/loader.go
package loader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Be4Die/go-evm/vm"
)

func LoadProgram(filename string, baseAddr uint16, memory *vm.Memory) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	addr := baseAddr

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		bytes, err := parseLine(line)
		if err != nil {
			return err
		}

		for _, b := range bytes {
			if err := memory.WriteByteAt(addr, b); err != nil {
				return err
			}
			addr++
		}
	}

	return scanner.Err()
}

func parseLine(line string) ([]byte, error) {
	if len(line) != 6 {
		return nil, fmt.Errorf("invalid line length")
	}

	bytes := make([]byte, 3)
	for i := 0; i < 3; i++ {
		hexStr := line[i*2 : i*2+2]
		b, err := strconv.ParseUint(hexStr, 16, 8)
		if err != nil {
			return nil, err
		}
		bytes[i] = byte(b)
	}
	return bytes, nil
}