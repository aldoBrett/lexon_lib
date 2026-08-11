package engine

import (
	"github.com/google/uuid"
)

type CreateParams struct {
	LegalProcedure LegalProcedure
	LegalRecord    LegalRecord
	LegalClaims    []LegalClaim
	Initialize     *bool
}

func (e *EngineHandler) Create(params CreateParams) error {
	// Default initialize to true if not provided
	initialize := true
	if params.Initialize != nil {
		initialize = *params.Initialize
	}

	savedLegalProcedure, err := e.repos.LegalProcedure.SaveLegalProcedure(params.LegalProcedure)
	if err != nil {
		return err
	}
	e.legalProcedure = savedLegalProcedure

	if _, err := e.repos.MachineInstances.CreateMachineInstanceForLegalProcedure(params.LegalProcedure.ID); err != nil {
		return err
	}

	// e.repos.LegalProcedure.SaveLegalProcedure(params.legalProcedure)
	legalRecord := params.LegalRecord
	legalRecord.LegalProcedureID = &params.LegalProcedure.ID

	savedLegalRecord, err := e.repos.LegalRecord.SaveLegalRecord(legalRecord)
	if err != nil {
		return err
	}

	for _, claim := range params.LegalClaims {
		claim.LegalRecordID = savedLegalRecord.ID

		e.repos.LegalClaims.SaveLegalClaim(claim)
	}

	if initialize {
		// if _, err := e.repos.MachineInstances.CreateMachineInstanceForLegalProcedure(params.legalProcedure.ID); err != nil {
		// 	fmt.Println("eeeeeeeeeeeeeee: ", err)
		// 	return err
		// }

		// We're going to initalize the machine because the user has set the legal claims
		if params.LegalClaims != nil {
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

			if params.LegalRecord.LegalActionID != nil {
				legalActionEventID := "CIV.ORD.E002"

				_, err := e.ProcessSignal(ProcessSignalParams{
					signal: MachineSignal{
						ID:      uuid.New().String(),
						Code:    "",
						Origin:  "system",
						EventID: &legalActionEventID,
					},
				})

				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}
