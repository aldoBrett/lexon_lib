package engine

/**
 * This function is used to get a LegalProcedure from the database by its
 * id.
 */
func (e *EngineHandler) EngineGetLegalProcedure(id string) LegalProcedure {
	query := `SELECT id, name, description FROM lexon.legal_procedures WHERE id = $1`
	var legalProcedure LegalProcedure
	err := e.pool.QueryRow(e.ctx, query, id).Scan(&legalProcedure.ID, &legalProcedure.Name, &legalProcedure.Description)
	if err != nil {
		panic(err)
	}

	return legalProcedure
}
