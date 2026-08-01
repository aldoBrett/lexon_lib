package engine

import "fmt"

type ProcessSignalParams struct {
	signal           MachineSignal
	legalProcedureID *string
}

func (e *EngineHandler) EngineProcessSignal(params ProcessSignalParams) {
	fmt.Println("------------->params: ", params)

	if params.legalProcedureID != nil {
		machineInstance, err := e.repos.MachineInstance.GetMachineInstanceByLegalProcedureID(*params.legalProcedureID)
		if err != nil {
			fmt.Println("Error retrieving machine instance:", err)
			return
		}

		fmt.Println("Machine Instance ID:", machineInstance.ID)
		fmt.Println("Current State ID:", machineInstance.CurrentStateID)
		fmt.Println("Legal Procedure ID:", machineInstance.LegalProcedureID)
	}
}
