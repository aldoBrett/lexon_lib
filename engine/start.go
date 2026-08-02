package engine

import "fmt"

type StartParams struct {
	legalProcedure LegalProcedure
	legalRecord    LegalRecord
	legalClaims    []LegalClaim
	// TODOO: implement when the a document was used to create the legal procedure
	// documentData   *DocumentData
}

func (e *EngineHandler) Start(params StartParams) error {
	if err := e.repos.LegalProcedure.SaveLegalProcedure(params.legalProcedure); err != nil {
		return err
	}

	if err := e.repos.MachineInstance.CreateMachineInstanceForLegalProcedure(params.legalProcedure.ID); err != nil {
		return err
	}

	fmt.Println("|||| legalProcedure: ", params.legalProcedure)
	fmt.Println("|||| legalRecord: ", params.legalRecord)
	fmt.Println("|||| legalClaims: ", params.legalClaims)
	e.repos.LegalProcedure.SaveLegalProcedure(params.legalProcedure)
	legalRecord := params.legalRecord
	legalRecord.LegalProcedureID = &params.legalProcedure.ID

	savedLegalRecord, err := e.repos.LegalRecord.SaveLegalRecord(legalRecord)
	fmt.Println("ssssssssssssssssssss: ", savedLegalRecord)
	if err != nil {
		return err
	}

	fmt.Println("======-=-=-=-=-=-=-=-=-=-")

	for _, claim := range params.legalClaims {
		claim.LegalRecordID = savedLegalRecord.ID

		fmt.Println("-------------claim: ", claim)
		e.repos.LegalClaims.SaveLegalClaim(claim)
	}

	return nil
}
