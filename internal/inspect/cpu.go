package inspect

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetCPUInfo() CPUInfo {
	out, err := exec.Command("lscpu").Output()
	if err != nil {
		return CPUInfo{}
	}

	text := string(out)

	cores := extractInt(text, `^CPU\(s\):\s+(\d+)`)
	threads := extractInt(text, `^Thread\(s\) per core:\s+(\d+)`)
	model := extractString(text, `^Model name:\s+(.+)`)

	return CPUInfo{
		Cores:   cores,
		Threads: threads,
		Model:   model,
	}
}

func extractInt(text, pattern string) int {
	re := regexp.MustCompile(pattern)
	for _, line := range strings.Split(text, "\n") {
		if re.MatchString(line) {
			m := re.FindStringSubmatch(line)
			val, _ := strconv.Atoi(strings.TrimSpace(m[1]))
			return val
		}
	}
	return 0
}

func extractString(text, pattern string) string {
	re := regexp.MustCompile(pattern)
	for _, line := range strings.Split(text, "\n") {
		if re.MatchString(line) {
			m := re.FindStringSubmatch(line)
			return strings.TrimSpace(m[1])
		}
	}
	return ""
}