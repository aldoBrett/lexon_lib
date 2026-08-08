package engine

import (
	"amox/lex_engine_lib/domain"
	"fmt"
)

type ProcessSignalParams struct {
	signal           MachineSignal
	legalProcedureID *string
}

// applyTransition moves a MachineInstance to the target state of the given transition and
// records the move in the transitions history table. Every state change should go through
// this so the history table stays a complete audit trail.
func (e *EngineHandler) applyTransition(machineInstance *domain.MachineInstance, transition *domain.MachineStateTransition) (*domain.MachineInstance, error) {
	updatedMachineInstance, err := e.repos.MachineInstances.UpdateMachineInstanceCurrentState(
		machineInstance.ID,
		transition.TargetStateID,
	)
	if err != nil {
		fmt.Println("Error updating machine instance:", err)
		return nil, err
	}

	if _, err := e.repos.MachineTransitionsHistory.CreateMachineStateTransitionHistory(machineInstance.ID, transition.ID); err != nil {
		fmt.Println("Error recording machine state transition history:", err)
		return nil, err
	}

	return updatedMachineInstance, nil
}

func (e *EngineHandler) ProcessSignal(params ProcessSignalParams) (*domain.MachineInstance, error) {

	if params.legalProcedureID != nil {
		//? Should this go?
		machineInstance, err := e.repos.MachineInstances.GetMachineInstanceByLegalProcedureID(*params.legalProcedureID)
		if err != nil {
			fmt.Println("Error retrieving machine instance:", err)
			return nil, err
		}

		fmt.Println("Machine Instance ID:", machineInstance.ID)
		fmt.Println("Current State ID:", machineInstance.CurrentStateID)
		fmt.Println("Legal Procedure ID:", machineInstance.LegalProcedureID)

	}

	if params.signal.EventID != nil {
		machineInstance, err := e.repos.MachineInstances.GetMachineInstanceByLegalProcedureID(e.legalProcedure.ID)
		if err != nil {
			fmt.Println("Error retrieving machine instance:", err)
			return nil, err
		}

		currentStateID := "CIV.ORD.S00"
		if machineInstance.CurrentStateID != nil {
			currentStateID = *machineInstance.CurrentStateID
		}

		transition, err := e.repos.MachineTransitions.GetMachineTransitionBySourceAndEvent(currentStateID, *params.signal.EventID)
		if err != nil {
			fmt.Println("Error retrieving transition:", err)
			return nil, err
		}

		return e.applyTransition(machineInstance, transition)
	}

	return nil, nil
}
