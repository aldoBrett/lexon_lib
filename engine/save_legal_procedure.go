package engine

/**
 * This function will look for the ID of the legal procedure, if it doesn't
 * exists, it will create a new one. If it exists, it will update the name and
 * description of the legal procedure.
 */
func (e *EngineHandler) EngineSaveLegalProcedure(legalProcedure LegalProcedure) {
	query := `SELECT id FROM lexon.legal_procedures WHERE id = $1`

	var id string
	err := e.pool.QueryRow(e.ctx, query, legalProcedure.ID).Scan(&id)
	if err != nil {
		// If the legal procedure doesn't exist, create a new one
		insertQuery := `INSERT INTO lexon.legal_procedures (id, name, description) VALUES ($1, $2, $3)`
		_, err = e.pool.Exec(e.ctx, insertQuery, legalProcedure.ID, legalProcedure.Name, legalProcedure.Description)
		if err != nil {
			panic(err)
		}

		machineInstanceError := e.EngineCreate(legalProcedure.ID)
		if machineInstanceError != nil {
			panic(machineInstanceError)
		}
	} else {
		// If the legal procedure exists, update its name and description
		updateQuery := `UPDATE lexon.legal_procedures SET name = $1, description = $2 WHERE id = $3`
		_, err = e.pool.Exec(e.ctx, updateQuery, legalProcedure.Name, legalProcedure.Description, legalProcedure.ID)
		if err != nil {
			panic(err)
		}
	}
}
