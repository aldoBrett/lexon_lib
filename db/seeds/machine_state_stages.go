package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MachineStateStage mirrors a row of lexon.machine_state_stages.
type MachineStateStage struct {
	ID          string
	Code        string
	Name        string
	Description string
	Content     string
	StageOrder  int
}

// MachineStateStages is the seed data for the CIV.ORD (civil ordinario) master stages.
var MachineStateStages = []MachineStateStage{
	{ID: "CIV.M01", Code: "POSTULATORIA", Name: "Postulatoria / fijación de la litis", Description: "Desde preparación/presentación de demanda hasta que quedan definidas las posiciones de las partes.", Content: "Demanda; prevención; admisión; emplazamiento; contestación; reconvención; réplica; dúplica; fijación de litis.", StageOrder: 1},
	{ID: "CIV.M02", Code: "DEPURACION_PROBATORIA", Name: "Depuración procesal y calificación probatoria", Description: "Control de cuestiones procesales y determinación de qué pruebas ya ofrecidas serán admitidas y cómo se prepararán.", Content: "Excepciones procesales; depuración; admisión/desechamiento de pruebas ofrecidas.", StageOrder: 2},
	{ID: "CIV.M03", Code: "PROBATORIA", Name: "Probatoria", Description: "Preparación y desahogo de medios de prueba.", Content: "Documental; confesional; testimonial; pericial; inspección; electrónica; incidencias de desahogo.", StageOrder: 3},
	{ID: "CIV.M04", Code: "CONCLUSIVA", Name: "Conclusiva", Description: "Cierre de la instrucción y formulación de alegatos antes de resolver.", Content: "Cierre de pruebas; alegatos; asunto en estado de sentencia.", StageOrder: 4},
	{ID: "CIV.M05", Code: "RESOLUTIVA", Name: "Resolutiva", Description: "Emisión y comunicación de la sentencia y determinación de sus efectos inmediatos.", Content: "Sentencia; notificación; análisis de resultado; firmeza potencial.", StageOrder: 5},
	{ID: "CIV.M06", Code: "IMPUGNATIVA", Name: "Impugnativa", Description: "Trámite de recursos o medios de impugnación contra resoluciones.", Content: "Apelación; revocación cuando proceda; alzada; confirmación; modificación; revocación; reposición.", StageOrder: 6},
	{ID: "CIV.M07", Code: "EJECUCION", Name: "Ejecución / cumplimiento", Description: "Obtención del cumplimiento voluntario o forzoso de la resolución.", Content: "Solicitud de ejecución; plazo de cumplimiento; requerimiento; embargo; avalúo; remate/adjudicación; pago; conclusión.", StageOrder: 7},
	{ID: "TRANSVERSAL", Code: "TRANSVERSAL", Name: "Transversal", Description: "REVISAR CON ERNESTO COMO FUNCIONA ESTA ETAPA", Content: "", StageOrder: 8},
}

// SeedMachineStateStages upserts MachineStateStages into lexon.machine_state_stages.
func SeedMachineStateStages(ctx context.Context, pool *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for _, s := range MachineStateStages {
		batch.Queue(
			`INSERT INTO lexon.machine_state_stages (id, code, name, description, content, stage_order)
			 VALUES ($1, $2, $3, $4, $5, $6)
			 ON CONFLICT (id) DO UPDATE SET
			   code = EXCLUDED.code,
			   name = EXCLUDED.name,
			   description = EXCLUDED.description,
			   content = EXCLUDED.content,
			   stage_order = EXCLUDED.stage_order`,
			s.ID, s.Code, s.Name, s.Description, s.Content, s.StageOrder,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range MachineStateStages {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("seeds: insert machine state stage: %w", err)
		}
	}

	return nil
}
