// Package seeds provides static seed data and loaders for lexon reference tables.
package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MachineState mirrors a row of lexon.machine_states.
type MachineState struct {
	ID          string
	Name        string
	Description string
	StageID     string
	Kind        string
}

// MachineStates is the seed data for the CIV.ORD (civil ordinario) machine states.
var MachineStates = []MachineState{
	{ID: "CIV.ORD.S00", Name: "Diagnóstico de procedencia", StageID: "CIV.M01", Kind: "Activo", Description: "Determinar si la pretensión corresponde al juicio ordinario civil o debe desviarse a otra vía."},
	{ID: "CIV.ORD.S01", Name: "Preparación de demanda", StageID: "CIV.M01", Kind: "Activo", Description: "Integrar escrito inicial, personalidad, anexos, pruebas y domicilios."},
	{ID: "CIV.ORD.S02", Name: "Demanda presentada", StageID: "CIV.M01", Kind: "Activo", Description: "La demanda ya fue ingresada y espera revisión judicial."},
	{ID: "CIV.ORD.S03", Name: "Revisión judicial inicial", StageID: "CIV.M01", Kind: "Activo", Description: "El órgano jurisdiccional verifica requisitos, personalidad y procedencia formal."},
	{ID: "CIV.ORD.S04", Name: "Prevención pendiente", StageID: "CIV.M01", Kind: "Activo", Description: "Existe requerimiento de aclaración, corrección o complemento."},
	{ID: "CIV.ORD.S05", Name: "Demanda admitida", StageID: "CIV.M01", Kind: "Activo", Description: "Se ordenó correr traslado y emplazar."},
	{ID: "CIV.ORD.S06", Name: "Emplazamiento en preparación", StageID: "CIV.M01", Kind: "Activo", Description: "Se prepara o diligencia la primera notificación."},
	{ID: "CIV.ORD.S07", Name: "Emplazamiento eficaz", StageID: "CIV.M01", Kind: "Activo", Description: "El demandado quedó legalmente emplazado y corre plazo para contestar."},
	{ID: "CIV.ORD.S08", Name: "Plazo de contestación en curso", StageID: "CIV.M01", Kind: "Activo", Description: "Corre plazo de nueve días para contestar."},
	{ID: "CIV.ORD.S09", Name: "Contestación registrada", StageID: "CIV.M01", Kind: "Activo", Description: "La contestación fue ingresada y debe validarse."},
	{ID: "CIV.ORD.S10", Name: "Contestación precluida o no presentada", StageID: "CIV.M01", Kind: "Activo", Description: "Venció el plazo; la demanda se tiene contestada en sentido negativo."},
	{ID: "CIV.ORD.S11", Name: "Vista para réplica", StageID: "CIV.M01", Kind: "Activo", Description: "El actor debe pronunciarse sobre contestación, excepciones o hechos nuevos."},
	{ID: "CIV.ORD.S12", Name: "Contestación de reconvención pendiente", StageID: "CIV.M01", Kind: "Activo", Description: "El actor original debe contestar la reconvención."},
	{ID: "CIV.ORD.S13", Name: "Réplica en curso", StageID: "CIV.M01", Kind: "Activo", Description: "La parte correspondiente debe formular réplica."},
	{ID: "CIV.ORD.S14", Name: "Dúplica en curso", StageID: "CIV.M01", Kind: "Activo", Description: "La contraparte debe formular dúplica."},
	{ID: "CIV.ORD.S15", Name: "Litis fijada", StageID: "CIV.M01", Kind: "Activo", Description: "Concluyó la etapa de planteamientos y se delimitó la controversia."},
	{ID: "CIV.ORD.S16", Name: "Calificación y preparación de pruebas", StageID: "CIV.M02", Kind: "Activo", Description: "Se califican, admiten y preparan medios de prueba."},
	{ID: "CIV.ORD.S17", Name: "Pruebas en desahogo", StageID: "CIV.M03", Kind: "Activo", Description: "Se reciben los medios de prueba admitidos."},
	{ID: "CIV.ORD.S18", Name: "Audiencia programada", StageID: "CIV.M03", Kind: "Activo", Description: "Existe fecha para audiencia de pruebas y alegatos."},
	{ID: "CIV.ORD.S19", Name: "Audiencia diferida", StageID: "CIV.M03", Kind: "Activo", Description: "No se desahogaron todas las pruebas y se señaló continuación."},
	{ID: "CIV.ORD.S20", Name: "Audiencia y alegatos concluidos", StageID: "CIV.M04", Kind: "Activo", Description: "El asunto quedó en estado de sentencia."},
	{ID: "CIV.ORD.S21", Name: "Sentencia pendiente", StageID: "CIV.M04", Kind: "Activo", Description: "Corre plazo judicial para dictar sentencia."},
	{ID: "CIV.ORD.S22", Name: "Sentencia notificada", StageID: "CIV.M05", Kind: "Activo", Description: "La resolución definitiva fue comunicada y abre decisiones de impugnación o ejecución."},
	{ID: "CIV.ORD.S23", Name: "Decisión de impugnación o cumplimiento", StageID: "CIV.M05", Kind: "Activo", Description: "Se analiza y prepara el medio de defensa correspondiente."},
	{ID: "CIV.ORD.S24", Name: "Cumplimiento y ejecución", StageID: "CIV.M07", Kind: "Activo", Description: "Se busca el cumplimiento de sentencia o convenio."},
	{ID: "CIV.ORD.S25", Name: "Suspendido", StageID: "TRANSVERSAL", Kind: "Activo", Description: "El avance principal está detenido por causa jurídica o material registrada."},
	{ID: "CIV.ORD.S26", Name: "Incidente activo", StageID: "TRANSVERSAL", Kind: "Activo", Description: "Existe una cuestión incidental con flujo paralelo."},
	{ID: "CIV.ORD.S27", Name: "Concluido por sentencia", StageID: "CIV.M07", Kind: "Terminal", Description: "El procedimiento terminó por sentencia firme o cumplimiento."},
	{ID: "CIV.ORD.S28", Name: "Concluido por convenio", StageID: "TRANSVERSAL", Kind: "Terminal", Description: "El procedimiento terminó por convenio aprobado o cumplido."},
	{ID: "CIV.ORD.S29", Name: "Concluido por desistimiento", StageID: "TRANSVERSAL", Kind: "Terminal", Description: "El procedimiento terminó por desistimiento válido."},
	{ID: "CIV.ORD.S30", Name: "Concluido por caducidad", StageID: "TRANSVERSAL", Kind: "Terminal", Description: "La instancia caducó por inactividad procesal."},
	{ID: "CIV.ORD.S31", Name: "Demanda desechada", StageID: "CIV.M01", Kind: "Terminal", Description: "La demanda fue repelida o desechada."},
	{ID: "CIV.ORD.S32", Name: "Sobreseído", StageID: "TRANSVERSAL", Kind: "Terminal", Description: "El juicio terminó por actualización de causal de sobreseimiento."},
}

// SeedMachineStates upserts MachineStates into lexon.machine_states.
func SeedMachineStates(ctx context.Context, pool *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for _, s := range MachineStates {
		batch.Queue(
			`INSERT INTO lexon.machine_states (id, name, description, stage_id, kind)
			 VALUES ($1, $2, $3, $4, $5)
			 ON CONFLICT (id) DO UPDATE SET
			   name = EXCLUDED.name,
			   description = EXCLUDED.description,
			   stage_id = EXCLUDED.stage_id,
			   kind = EXCLUDED.kind`,
			s.ID, s.Name, s.Description, s.StageID, s.Kind,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range MachineStates {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("seeds: insert machine state: %w", err)
		}
	}

	return nil
}
