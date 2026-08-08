package engine

import (
	"github.com/google/uuid"
)

type CreateParams struct {
	legalProcedure LegalProcedure
	legalRecord    LegalRecord
	legalClaims    []LegalClaim
	initialize     *bool
}

func (e *EngineHandler) Create(params CreateParams) error {
	// Default initialize to true if not provided
	initialize := true
	if params.initialize != nil {
		initialize = *params.initialize
	}

	savedLegalProcedure, err := e.repos.LegalProcedure.SaveLegalProcedure(params.legalProcedure)
	if err != nil {
		return err
	}
	e.legalProcedure = savedLegalProcedure

	if _, err := e.repos.MachineInstances.CreateMachineInstanceForLegalProcedure(params.legalProcedure.ID); err != nil {
		return err
	}

	// e.repos.LegalProcedure.SaveLegalProcedure(params.legalProcedure)
	legalRecord := params.legalRecord
	legalRecord.LegalProcedureID = &params.legalProcedure.ID

	savedLegalRecord, err := e.repos.LegalRecord.SaveLegalRecord(legalRecord)
	if err != nil {
		return err
	}

	for _, claim := range params.legalClaims {
		claim.LegalRecordID = savedLegalRecord.ID

		e.repos.LegalClaims.SaveLegalClaim(claim)
	}

	if initialize {
		// if _, err := e.repos.MachineInstances.CreateMachineInstanceForLegalProcedure(params.legalProcedure.ID); err != nil {
		// 	fmt.Println("eeeeeeeeeeeeeee: ", err)
		// 	return err
		// }

		// We're going to initalize the machine because the user has set the legal claims
		if params.legalClaims != nil {
			eventID := "CIV.ORD.E001"

			_, err := e.ProcessSignal(ProcessSignalParams{
				signal: MachineSignal{
					ID:      uuid.New().String(),
					Code:    "",
					Origin:  "system",
					EventID: &eventID,
				},
			})

			if err != nil {
				return err
			}
		}

	}

	return nil
}
