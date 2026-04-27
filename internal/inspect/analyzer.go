package inspect

func Analyze(h HardwareProfile) Analysis {
	var notes []string

	score := 0.0

	if h.RAM.TotalGB >= 32 {
		score += 4
	} else if h.RAM.TotalGB >= 16 {
		score += 3
		notes = append(notes, "RAM suficiente para modelos médios")
	} else if h.RAM.TotalGB >= 8 {
		score += 2
		notes = append(notes, "RAM limitada")
	} else {
		notes = append(notes, "RAM insuficiente para LLMs modernos")
	}

	if h.GPU.Found && h.GPU.VRAMGB >= 12 {
		score += 4
	} else if h.GPU.Found && h.GPU.VRAMGB >= 6 {
		score += 3
		notes = append(notes, "GPU razoável")
	} else if h.GPU.Found {
		score += 1
		notes = append(notes, "GPU fraca")
	} else {
		notes = append(notes, "Sem GPU dedicada")
	}

	if h.CPU.Cores >= 8 {
		score += 2
	} else {
		score += 1
	}

	var model string
	var quant string

	if h.GPU.VRAMGB >= 16 {
		model = "qwen2.5-coder:14b"
		quant = "Q5_K_M"
	} else if h.GPU.VRAMGB >= 10 {
		model = "qwen2.5-coder:14b"
		quant = "Q4_K_M"
	} else if h.RAM.TotalGB >= 16 {
		model = "qwen2.5-coder:7b"
		quant = "Q4_K_M (CPU)"
	} else {
		model = "3B models"
		quant = "Q4 (limitado)"
	}

	return Analysis{
		ModelRecommendation: model,
		Quantization:        quant,
		Notes:               notes,
		Score:               score,
	}
}