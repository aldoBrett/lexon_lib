package engine

import "fmt"

type ProcessSignalParams struct {
	signal           MachineSignal
	legalProcedureID *string
}

func (e *EngineHandler) EngineProcessSignal(params ProcessSignalParams) {
	fmt.Println("------------->params: ", params)

	if params.legalProcedureID != nil {
		machineInstanceQuery := `SELECT id, current_state_id, legal_procedure_id FROM lexon.machine_instances WHERE legal_procedure_id = $1`
		var machineInstance MachineInstance
		err := e.pool.QueryRow(
			e.ctx,
			machineInstanceQuery,
			params.legalProcedureID,
		).Scan(
			&machineInstance.ID,
			&machineInstance.CurrentStateID,
			&machineInstance.LegalProcedureID,
		)
		if err != nil {
			fmt.Println("Error retrieving machine instance:", err)
			return
		}

		fmt.Println("Machine Instance ID:", machineInstance.ID)
		fmt.Println("Current State ID:", machineInstance.CurrentStateID)
		fmt.Println("Legal Procedure ID:", machineInstance.LegalProcedureID)
	}
}
