package inspect

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetGPUInfo() GPUInfo {
	out, err := exec.Command("nvidia-smi", "--query-gpu=name,memory.total", "--format=csv,noheader,nounits").Output()
	if err == nil {
		lines := strings.Split(string(out), "\n")
		if len(lines) > 0 && lines[0] != "" {
			parts := strings.Split(lines[0], ",")
			if len(parts) >= 2 {
				name := strings.TrimSpace(parts[0])
				vram, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

				return GPUInfo{
					Name:   name,
					VRAMGB: vram / 1024,
					Found:  true,
				}
			}
		}
	}

	out, err = exec.Command("sh", "-c", "lspci | grep -i vga").Output()
	if err == nil {
		return GPUInfo{
			Name:  strings.TrimSpace(string(out)),
			Found: true,
		}
	}

	return GPUInfo{Found: false}
}