package container

import (
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/docker/docker/client"
)

type ResourceParser struct {
	ReadTime string `json:"read"`
	CpuStats struct {
		CpuUsage struct {
			TotalUsage int64 `json:"total_usage"`
		} `json:"cpu_usage"`
	} `json:"cpu_stats"`

	MemoryStats struct {
		Usage int64 `json:"usage"`
		Stats struct {
			InactiveFile int64 `json:"inactive_file"`
		} `json:"stats"`
		Limit int64 `json:"limit"`
	} `json:"memory_stats"`

	NetworkStats map[string]struct {
		RxBytes int64 `json:"rx_bytes"`
		TxBytes int64 `json:"tx_bytes"`
	} `json:"networks"`
}

type CPUResources struct {
	Usage int64
}

type MemoryResources struct {
	UsageBytes int64
	LimitBytes int64
}

type NetworkResources struct {
	ReceiveBytes  int64
	TransmitBytes int64
}

type Resources struct {
	UpdatedAt time.Time
	CPU       CPUResources
	Memory    MemoryResources
	Network   map[string]NetworkResources
}

type Container struct {
	ContainerID string
	Project     string
	Service     string
	Resources   Resources
}

func ParseResourceUsage(data []byte) (*Resources, error) {
	var parser ResourceParser
	err := json.Unmarshal(data, &parser)
	if err != nil {
		return nil, err
	}

	var networkResources = make(map[string]NetworkResources)
	for key, stats := range parser.NetworkStats {
		networkResources[key] = NetworkResources{
			ReceiveBytes:  stats.RxBytes,
			TransmitBytes: stats.TxBytes,
		}
	}

	readTime, err := time.Parse(time.RFC3339Nano, parser.ReadTime)
	if err != nil {
		return nil, err
	}

	return &Resources{
		UpdatedAt: readTime,
		CPU:       CPUResources{Usage: parser.CpuStats.CpuUsage.TotalUsage},
		Memory: MemoryResources{
			UsageBytes: parser.MemoryStats.Usage - parser.MemoryStats.Stats.InactiveFile,
			LimitBytes: parser.MemoryStats.Limit,
		},
		Network: networkResources,
	}, nil
}

func (c *Container) ReadResource(ctx context.Context, cli *client.Client) {
	stats, err := cli.ContainerStats(ctx, c.ContainerID, false)
	if err != nil {
		return
	}
	defer stats.Body.Close()

	statsBytes, err := io.ReadAll(stats.Body)
	if err != nil {
		return
	}

	resources, err := ParseResourceUsage(statsBytes)
	if err != nil {
		return
	}

	c.Resources = *resources
}
