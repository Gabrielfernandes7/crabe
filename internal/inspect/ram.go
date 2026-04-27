package inspect

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetRAMInfo() RAMInfo {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return RAMInfo{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)

			if len(fields) >= 2 {
				kb, err := strconv.Atoi(fields[1])
				if err != nil {
					return RAMInfo{}
				}

				gb := kb / 1024 / 1024

				return RAMInfo{
					TotalGB: gb,
				}
			}
		}
	}

	return RAMInfo{}
}