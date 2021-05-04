package main

import (
	"context"
	"net/http"
)

type Payload struct {
	Code       int
	Ctx        context.Context
	W          http.ResponseWriter
	Data       []byte   `json:"data,omitempty"`
	EfficiencyRec EfficiencyRec `json:"eff_rec,omitempty"`
	Error      error
}

// EfficiencyRec ... data for recommendation
type EfficiencyRec struct {
	// Input
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	ControllerType string `json:"controller_type,omitempty"`
	ContainerName string `json:"container_name,omitempty"`
	Window string `json:"window,omitempty"`
	Percentile string `json:"percentile,omitempty"`
	TargetCpuUtil string `json:"target_cpu_util,omitempty"`
	TargetRamUtil string `json:"target_ram_util,omitempty"`
	// Output
	Efficiency int `json:"efficiency,omitempty"`
	CPU string `json:"cpu,omitempty"`
	Mem string `json:"memory,omitempty"`
}

// Sizing ... top level data
type Sizing struct {
	//MonthlySavings    float64      `json:"monthlySavings"`
	//MonthlySavingsCPU float64      `json:"monthlySavingsCPU"`
	//MonthlySavingsRAM float64      `json:"monthlySavingsRAM"`
	Controllers       []Controller `json:"controllers"`
}

// Controller ... different k8s controllers
type Controller struct {
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Namespace   string         `json:"namespace"`
	ClusterName string         `json:"clusterName"`
	//Pods        map[string]Pod `json:"pods"`
	Containers map[string]Container `json:"containers"`
}

// Container ... k8s container info
type Container struct {
	Name     string  `json:"name"`
	Efficiency float64 `json:"efficiency"`
	MonthlySavingsCPU float64 `json:"monthlySavingsCPU"`
 	Requests Compute `json:"requests"`
	Usage    Compute `json:"usage"`
	Target   Compute `json:"target"`
}

// Compute .. compute
type Compute struct {
	Cpu float64 `json:"cpu"`
	Ram float64 `json:"ram"`
}

