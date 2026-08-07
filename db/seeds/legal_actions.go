package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// LegalAction mirrors a row of lexon.legal_actions.
type LegalAction struct {
	ID          string
	Category    string
	SubCategory string
	ActionName  string
	Via         string
}

// LegalActions is the seed data for the civil legal actions catalog.
var LegalActions = []LegalAction{
	{ID: "ACT.CIV.0001", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción reivindicatoria", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0002", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción declarativa de propiedad", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0003", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción confesoria", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0004", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción negatoria", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0005", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción de deslinde", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0006", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción de apeo", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0007", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción de división de cosa común", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0008", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción de partición", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0009", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción de entrega de inmueble", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0010", Category: "Derechos Reales", SubCategory: "Propiedad", ActionName: "Acción de restitución derivada de derecho real", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0011", Category: "Obligaciones", SubCategory: "Cumplimiento", ActionName: "Cumplimiento de contrato", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0012", Category: "Obligaciones", SubCategory: "Cumplimiento", ActionName: "Cumplimiento de convenio", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0013", Category: "Obligaciones", SubCategory: "Cumplimiento", ActionName: "Cumplimiento de obligación de dar", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0014", Category: "Obligaciones", SubCategory: "Cumplimiento", ActionName: "Cumplimiento de obligación de hacer", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0015", Category: "Obligaciones", SubCategory: "Cumplimiento", ActionName: "Cumplimiento de obligación de no hacer", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0016", Category: "Contratos", SubCategory: "Rescisión", ActionName: "Rescisión de compraventa", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0017", Category: "Contratos", SubCategory: "Rescisión", ActionName: "Rescisión de arrendamiento", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0018", Category: "Contratos", SubCategory: "Rescisión", ActionName: "Rescisión de prestación de servicios", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0019", Category: "Contratos", SubCategory: "Rescisión", ActionName: "Rescisión de obra", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0020", Category: "Contratos", SubCategory: "Rescisión", ActionName: "Rescisión de suministro", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0021", Category: "Contratos", SubCategory: "Rescisión", ActionName: "Rescisión de promesa de compraventa", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0022", Category: "Contratos", SubCategory: "Resolución", ActionName: "Resolución contractual", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0023", Category: "Contratos", SubCategory: "Resolución", ActionName: "Terminación anticipada", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0024", Category: "Contratos", SubCategory: "Resolución", ActionName: "Terminación por incumplimiento", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0025", Category: "Nulidades", SubCategory: "Actos Jurídicos", ActionName: "Nulidad absoluta", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0026", Category: "Nulidades", SubCategory: "Actos Jurídicos", ActionName: "Nulidad relativa", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0027", Category: "Nulidades", SubCategory: "Actos Jurídicos", ActionName: "Inexistencia de acto jurídico", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0028", Category: "Nulidades", SubCategory: "Actos Jurídicos", ActionName: "Simulación absoluta", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0029", Category: "Nulidades", SubCategory: "Actos Jurídicos", ActionName: "Simulación relativa", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0030", Category: "Cobro", SubCategory: "Obligaciones", ActionName: "Pago de pesos", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0031", Category: "Cobro", SubCategory: "Obligaciones", ActionName: "Cobro de adeudo", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0032", Category: "Cobro", SubCategory: "Obligaciones", ActionName: "Cobro de saldo insoluto", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0033", Category: "Cobro", SubCategory: "Obligaciones", ActionName: "Cobro de honorarios", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0034", Category: "Cobro", SubCategory: "Obligaciones", ActionName: "Cobro de rentas", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0035", Category: "Cobro", SubCategory: "Obligaciones", ActionName: "Cobro de intereses", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0036", Category: "Cobro", SubCategory: "Obligaciones", ActionName: "Cobro de cláusula penal", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0037", Category: "Responsabilidad Civil", SubCategory: "Daños", ActionName: "Daños y perjuicios", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0038", Category: "Responsabilidad Civil", SubCategory: "Daños", ActionName: "Daño moral", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0039", Category: "Responsabilidad Civil", SubCategory: "Daños", ActionName: "Responsabilidad civil contractual", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0040", Category: "Responsabilidad Civil", SubCategory: "Daños", ActionName: "Responsabilidad civil extracontractual", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0041", Category: "Obligaciones", SubCategory: "Enriquecimiento", ActionName: "Enriquecimiento ilegítimo", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0042", Category: "Obligaciones", SubCategory: "Enriquecimiento", ActionName: "Pago de lo indebido", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0043", Category: "Declarativas", SubCategory: "Declaración", ActionName: "Declaración de existencia", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0044", Category: "Declarativas", SubCategory: "Declaración", ActionName: "Declaración de inexistencia", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0045", Category: "Declarativas", SubCategory: "Declaración", ActionName: "Declaración de validez", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0046", Category: "Declarativas", SubCategory: "Declaración", ActionName: "Declaración de nulidad", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0047", Category: "Declarativas", SubCategory: "Declaración", ActionName: "Declaración de prescripción", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0048", Category: "Declarativas", SubCategory: "Declaración", ActionName: "Declaración de usucapión", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0049", Category: "Constitutivas", SubCategory: "Derechos Reales", ActionName: "Constitución de servidumbre", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0050", Category: "Constitutivas", SubCategory: "Derechos Reales", ActionName: "Extinción de servidumbre", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0051", Category: "Constitutivas", SubCategory: "Derechos Reales", ActionName: "Constitución de usufructo", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0052", Category: "Constitutivas", SubCategory: "Derechos Reales", ActionName: "Cancelación de gravamen", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0053", Category: "Constitutivas", SubCategory: "Derechos Reales", ActionName: "Cancelación de hipoteca", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0054", Category: "Constitutivas", SubCategory: "Derechos Reales", ActionName: "Constitución de copropiedad", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0055", Category: "Posesorias", SubCategory: "Posesión", ActionName: "Conservación de posesión", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0056", Category: "Posesorias", SubCategory: "Posesión", ActionName: "Recuperación de posesión", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0057", Category: "Posesorias", SubCategory: "Posesión", ActionName: "Restitución posesoria", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0058", Category: "Contratos", SubCategory: "Compraventa", ActionName: "Cumplimiento de compraventa", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0059", Category: "Contratos", SubCategory: "Compraventa", ActionName: "Rescisión de compraventa", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0060", Category: "Contratos", SubCategory: "Compraventa", ActionName: "Nulidad de compraventa", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0061", Category: "Contratos", SubCategory: "Compraventa", ActionName: "Entrega de inmueble", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0062", Category: "Contratos", SubCategory: "Compraventa", ActionName: "Otorgamiento de escritura", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0063", Category: "Contratos", SubCategory: "Arrendamiento", ActionName: "Cumplimiento de arrendamiento", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0064", Category: "Contratos", SubCategory: "Arrendamiento", ActionName: "Pago de rentas", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0065", Category: "Contratos", SubCategory: "Arrendamiento", ActionName: "Daños derivados del arrendamiento", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0066", Category: "Contratos", SubCategory: "Mutuo", ActionName: "Cobro de mutuo", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0067", Category: "Contratos", SubCategory: "Mutuo", ActionName: "Cobro de intereses", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0068", Category: "Contratos", SubCategory: "Prestación de servicios", ActionName: "Cumplimiento", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0069", Category: "Contratos", SubCategory: "Prestación de servicios", ActionName: "Pago", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0070", Category: "Contratos", SubCategory: "Prestación de servicios", ActionName: "Rescisión", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0071", Category: "Contratos", SubCategory: "Obra", ActionName: "Cumplimiento de obra", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0072", Category: "Contratos", SubCategory: "Obra", ActionName: "Vicios ocultos", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0073", Category: "Contratos", SubCategory: "Obra", ActionName: "Defectos de construcción", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0074", Category: "Contratos", SubCategory: "Obra", ActionName: "Pago", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0075", Category: "Contratos", SubCategory: "Mandato", ActionName: "Rendición de cuentas", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0076", Category: "Contratos", SubCategory: "Mandato", ActionName: "Responsabilidad del mandatario", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0077", Category: "Contratos", SubCategory: "Depósito", ActionName: "Entrega de bienes", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0078", Category: "Contratos", SubCategory: "Depósito", ActionName: "Restitución", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0079", Category: "Contratos", SubCategory: "Depósito", ActionName: "Daños", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0080", Category: "Garantías", SubCategory: "Fianza/Hipoteca", ActionName: "Cobro de fianza", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0081", Category: "Garantías", SubCategory: "Fianza/Hipoteca", ActionName: "Cumplimiento de fianza", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0082", Category: "Garantías", SubCategory: "Fianza/Hipoteca", ActionName: "Cancelación de hipoteca", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0083", Category: "Obligaciones", SubCategory: "Extinción", ActionName: "Mora", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0084", Category: "Obligaciones", SubCategory: "Extinción", ActionName: "Novación", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0085", Category: "Obligaciones", SubCategory: "Extinción", ActionName: "Compensación", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0086", Category: "Obligaciones", SubCategory: "Extinción", ActionName: "Confusión", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0087", Category: "Obligaciones", SubCategory: "Extinción", ActionName: "Remisión", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0088", Category: "Obligaciones", SubCategory: "Extinción", ActionName: "Prescripción", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0089", Category: "Sociedades Civiles", SubCategory: "Socios", ActionName: "Exclusión de socio", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0090", Category: "Sociedades Civiles", SubCategory: "Socios", ActionName: "Liquidación", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0091", Category: "Sociedades Civiles", SubCategory: "Socios", ActionName: "Responsabilidad entre socios", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0092", Category: "Responsabilidad Profesional", SubCategory: "Profesionales", ActionName: "Responsabilidad médica", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0093", Category: "Responsabilidad Profesional", SubCategory: "Profesionales", ActionName: "Responsabilidad de arquitecto", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0094", Category: "Responsabilidad Profesional", SubCategory: "Profesionales", ActionName: "Responsabilidad de ingeniero", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0095", Category: "Responsabilidad Profesional", SubCategory: "Profesionales", ActionName: "Responsabilidad de contador", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0096", Category: "Responsabilidad Profesional", SubCategory: "Profesionales", ActionName: "Responsabilidad de abogado", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0097", Category: "Responsabilidad Profesional", SubCategory: "Profesionales", ActionName: "Responsabilidad de notario", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0098", Category: "Responsabilidad Objetiva", SubCategory: "Especial", ActionName: "Daños por animales", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0099", Category: "Responsabilidad Objetiva", SubCategory: "Especial", ActionName: "Daños por cosas peligrosas", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0100", Category: "Responsabilidad Objetiva", SubCategory: "Especial", ActionName: "Productos defectuosos", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0101", Category: "Responsabilidad Objetiva", SubCategory: "Especial", ActionName: "Accidentes", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0102", Category: "Derechos Reales", SubCategory: "Otros", ActionName: "Usucapión", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0103", Category: "Derechos Reales", SubCategory: "Otros", ActionName: "Servidumbres", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0104", Category: "Derechos Reales", SubCategory: "Otros", ActionName: "Usufructo", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0105", Category: "Derechos Reales", SubCategory: "Otros", ActionName: "Uso", Via: "Juicio Ordinario Civil"},
	{ID: "ACT.CIV.0106", Category: "Derechos Reales", SubCategory: "Otros", ActionName: "Habitación", Via: "Juicio Ordinario Civil"},
}

// SeedLegalActions upserts LegalActions into lexon.legal_actions.
func SeedLegalActions(ctx context.Context, pool *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for _, a := range LegalActions {
		batch.Queue(
			`INSERT INTO lexon.legal_actions (id, category, sub_category, action_name, via)
			 VALUES ($1, $2, $3, $4, $5)
			 ON CONFLICT (id) DO UPDATE SET
			   category = EXCLUDED.category,
			   sub_category = EXCLUDED.sub_category,
			   action_name = EXCLUDED.action_name,
			   via = EXCLUDED.via`,
			a.ID, a.Category, a.SubCategory, a.ActionName, a.Via,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range LegalActions {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("seeds: insert legal action: %w", err)
		}
	}

	return nil
}
