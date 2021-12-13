package models

type Server struct {
	Os   *OS   `json:"os"`
	Cpu  *Cpu  `json:"cpu"`
	Ram  *Ram  `json:"ram"`
	Disk *Disk `json:"disk"`
}

type OS struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	UsedCpus []float64 `json:"usedCpus"`
	Cores    int       `json:"cores"`
}

type Ram struct {
	UsedMiB     uint64  `json:"usedMiB"`
	TotalMiB    uint64  `json:"totalMiB"`
	UsedPercent float64 `json:"usedPercent"`
}

type Disk struct {
	UsedMiB     uint64  `json:"usedMiB"`
	UsedGiB     uint64  `json:"usedGiB"`
	TotalMiB    uint64  `json:"totalMiB"`
	TotalGiB    uint64  `json:"totalGiB"`
	UsedPercent float64 `json:"usedPercent"`
}
