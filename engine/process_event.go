package engine

import "fmt"

func (e *EngineHandler) EngineProcessEvent(event MachineEvent) {
	fmt.Println("EngineProcessEvent-------->event: ", event)
}
