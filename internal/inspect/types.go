package inspect

type CPUInfo struct {
	Cores   int
	Threads int
	Model   string
}

type RAMInfo struct {
	TotalGB int
}

type GPUInfo struct {
	Name   string
	VRAMGB int
	Found  bool
}

type HardwareProfile struct {
	CPU CPUInfo
	RAM RAMInfo
	GPU GPUInfo
}

type Analysis struct {
	ModelRecommendation string
	Quantization        string
	Notes               []string
	Score               float64
}