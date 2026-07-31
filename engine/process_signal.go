package engine

import "fmt"

func (e *EngineHandler) EngineProcessSignal(signal MachineSignal) {
	fmt.Println("-------------------->signal: ", signal)
}
