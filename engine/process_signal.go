package engine

import (
	"amox/lex_engine_lib/domain"
	"fmt"
)

type ProcessSignalParams struct {
	signal           MachineSignal
	legalProcedureID *string
}

func (e *EngineHandler) ProcessSignal(params ProcessSignalParams) (*domain.MachineInstance, error) {
	fmt.Println("|>*<|*****------------->params: ", params)

	if params.legalProcedureID != nil {
		machineInstance, err := e.repos.MachineInstances.GetMachineInstanceByLegalProcedureID(*params.legalProcedureID)
		if err != nil {
			fmt.Println("Error retrieving machine instance:", err)
			return nil, err
		}

		fmt.Println("Machine Instance ID:", machineInstance.ID)
		fmt.Println("Current State ID:", machineInstance.CurrentStateID)
		fmt.Println("Legal Procedure ID:", machineInstance.LegalProcedureID)

	}

	fmt.Println("SSsSsSSsSssSsSsignal:  ", params.signal)

	if params.signal.TransitionID != nil {
		// Bring the transition from the database using the transition ID
		fmt.Println("TTTTTTTTTTT: ", params.signal.TransitionID)
		transition, err := e.repos.MachineTransitions.GetMachineTransitionByID(*params.signal.TransitionID)
		if err != nil {
			fmt.Println("Error retrieving transition:", err)
			return nil, err
		}

		fmt.Println("||||| Transition:", transition)
		fmt.Println("legalProcedure:  ", e.legalProcedure)
		machineInstance, err := e.repos.MachineInstances.GetMachineInstanceByLegalProcedureID(e.legalProcedure.ID)
		if err != nil {
			fmt.Println("Error retrieving machine instance:", err)
			return nil, err
		}

		fmt.Println("> Machine Instance ID:", machineInstance.ID)
		fmt.Println("> Current State ID:", machineInstance.CurrentStateID)
		fmt.Println("> Legal Procedure ID:", machineInstance.LegalProcedureID)
		// This is the case for when we're starting the machine
		if machineInstance.CurrentStateID == nil && transition.ID == "T001" {
			// TODO: add the history stuff

			updatedMachineInstance, err := e.repos.MachineInstances.UpdateMachineInstanceCurrentState(
				machineInstance.ID,
				transition.TargetStateID,
			)
			if err != nil {
				fmt.Println("Error updating machine instance:", err)
				return nil, err
			}

			return updatedMachineInstance, nil
		}

	}

	return nil, nil
}
