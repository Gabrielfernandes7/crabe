package inspect

import "github.com/Gabrielfernandes7/crabe/internal/ui"

func Run() {
	ui.Init()
	ui.Title("Crabe Inspect - Análise de Hardware")

	profile := HardwareProfile{
		CPU: GetCPUInfo(),
		RAM: GetRAMInfo(),
		GPU: GetGPUInfo(),
	}

	analysis := Analyze(profile)

	printProfile(profile)
	printAnalysis(analysis)
}

func printProfile(h HardwareProfile) {
	ui.Section("Sistema")

	ui.Info("CPU: %s (%d cores)", h.CPU.Model, h.CPU.Cores)
	ui.Info("RAM: %d GB", h.RAM.TotalGB)

	if h.GPU.Found {
		ui.Info("GPU: %s (%d GB VRAM)", h.GPU.Name, h.GPU.VRAMGB)
	} else {
		ui.Warning("GPU não detectada")
	}
}

func printAnalysis(a Analysis) {
	ui.Section("Análise")

	ui.Success("Modelo recomendado: %s", a.ModelRecommendation)
	ui.Success("Quantização: %s", a.Quantization)

	for _, n := range a.Notes {
		ui.Warning(n)
	}

	ui.Section("Score Crabe")
	ui.Info("%.1f / 10", a.Score)
}