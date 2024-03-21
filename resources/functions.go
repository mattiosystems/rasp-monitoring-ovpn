package resources

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
	"strconv"
	"strings"
	"time"
)

func bytesToMB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func bytesToGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

func GetCPUUsedPercentage() string {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error while reading CPU current usage:", err)
		return "X"
	}
	roundedCPUPercent := strconv.FormatFloat(cpuPercent[0], 'f', 0, 64)
	return roundedCPUPercent
}

func GetRAMUsed() string {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Failed to get RAM usage:", err)
		return "X"
	}
	return strconv.FormatFloat(bytesToGB(memInfo.Used), 'f', 2, 64)
}

func GetRAMUsedPercentage() string {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Failed to get RAM usage:", err)
		return "X"
	}
	return strconv.FormatFloat(memInfo.UsedPercent, 'f', 0, 64)
}

func GetRAMTotal() string {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Failed to get RAM usage:", err)
		return "X"
	}
	return strconv.FormatFloat(bytesToGB(memInfo.Total), 'f', 2, 64)
}

func GetStorageUsed() string {
	storage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Failed to get storage usage:", err)
		return "X"
	}
	return strconv.FormatFloat(bytesToGB(storage.Used), 'f', 2, 64)
}

func GetStorageUsedPercentage() string {
	storage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Failed to get storage usage:", err)
		return "X"
	}
	return strconv.FormatFloat(storage.UsedPercent, 'f', 0, 64)
}

func GetStorageTotal() string {
	storage, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Failed to get storage usage:", err)
		return "X"
	}
	return strconv.FormatFloat(bytesToGB(storage.Total), 'f', 2, 64)
}

func GetCPUTemperature() string {
	cpuTempFile := "/sys/class/thermal/thermal_zone0/temp"
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	tempStr, err := os.ReadFile(cpuTempFile)
	if err != nil {
		fmt.Println("Failed to read CPU temperature:", err)
		return "X"
	}
	tempInt, err := strconv.Atoi(strings.TrimSpace(string(tempStr)))
	if err != nil {
		fmt.Println("Failed to convert CPU temperature:", err)
		return "X"
	}
	tempCelsius := float64(tempInt) / 1000.0
	return strconv.FormatFloat(tempCelsius, 'f', 0, 64)
}

func GetCPUTemperaturePercentage() string {
	temp, err := strconv.ParseInt(GetCPUTemperature(), 10, 32)
	maxTemp := 85.0
	if err != nil {
		fmt.Println("Failed to convert temperature:", err)
		return "X"
	}
	percent := float64(temp) / maxTemp * 100
	return strconv.FormatFloat(percent, 'f', 0, 64)
}

func GetHostname() string {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Failed to get system data:", err)
		return "Loading hostname..."
	}
	return info.Hostname
}

func GetOS() string {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Failed to get system data:", err)
		return "Loading OS..."
	}
	return info.OS
}

func GetKernel() string {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Failed to get system data:", err)
		return "Loading Kernel Version..."
	}
	return info.KernelVersion
}

func GetPlatform() string {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Failed to get system data:", err)
		return "Loading Platform..."
	}
	return info.Platform
}

func GetPlatformFamily() string {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Failed to get system data:", err)
		return "Loading Platform Family..."
	}
	return info.PlatformFamily
}

func GetPlatformVersion() string {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Failed to get system data:", err)
		return "Loading Platform Family..."
	}
	return info.PlatformVersion
}

func GetCPUVendorID() string {
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Failed to get CPU data:", err)
		return "Loading CPU Vendor ID..."
	}
	return info[0].VendorID
}

func GetCPUModel() string {
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Failed to get CPU data:", err)
		return "Loading CPU model..."
	}
	return info[0].Model
}

func GetCPUMhz() string {
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Failed to get CPU data:", err)
		return "Loading CPU Mhz..."
	}
	return strconv.FormatFloat(info[0].Mhz, 'f', 0, 64)
}
