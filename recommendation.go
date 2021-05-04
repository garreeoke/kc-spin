package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	kcBaseUrl = "http://kubecost.work.garreeoke.io/model/savings/requestSizing?"
)

func (e *EfficiencyRec) PodRecommendation() error {

	var resizing Sizing
	// Make api call to KC
	kcUrl := kcBaseUrl + "window=" + e.Window + "&p=" + e.Percentile + "&targetCPUUtilization=" + e.TargetCpuUtil + "targetRAMUtilization=" + e.TargetRamUtil
	log.Println("URL: ", kcUrl)
	resp, err := http.Get(kcUrl)
	if err != nil {
		return err
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//log.Println(string(r))
	err = json.Unmarshal(r, &resizing)
	if err != nil {
		return err
	}

	// Loop through requestSizingData
	for _, controller := range resizing.Controllers {
		switch controller.Type {
		case "deployment", "statefulset":
			if controller.Namespace == e.Namespace && controller.Name == e.Name {
				for containerName, container := range controller.Containers {
					if containerName == e.ContainerName {
						log.Println("Container name: ", containerName)
						e.CPU = fmt.Sprintf("%vm", int(container.Target.Cpu*1000))
						e.Mem = fmt.Sprintf("%vMi", int(container.Target.Ram*1000*.95367431640625))
						e.Efficiency = int(container.Efficiency * 100)
						log.Println("Container.Target.CPU: ", e.CPU)
						log.Println("Container.Target.Ram: ", e.Mem)
						if e.Efficiency < 50 {
							log.Printf("Efficiency %v is below threshold of 50", e.Efficiency)
						}
					}
				}
			}
		}
	}
	return nil
}
