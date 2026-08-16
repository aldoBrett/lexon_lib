package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MachineEvent mirrors a row of lexon.machine_events.
type MachineEvent struct {
	ID          string
	Name        string
	Description string
	Kind        string
	Issuer      string
}

// MachineEvents is the seed data for the CIV.ORD (civil ordinario) machine events.
var MachineEvents = []MachineEvent{
	{ID: "CIV.ORD.E001", Name: "Seleccionar pretensión", Kind: "CAPTURA_MANUAL", Issuer: "Usuario", Description: "Define la finalidad jurídica perseguida."},
	{ID: "CIV.ORD.E002", Name: "Confirmar vía ordinaria", Kind: "VALIDACIÓN", Issuer: "Sistema/abogado", Description: "Confirma que la controversia no tiene atención especial."},
	{ID: "CIV.ORD.E003", Name: "Elaborar demanda", Kind: "GENERACION_DOCUMENTO", Issuer: "Abogado", Description: "Elabora demanda con la información del expediente."},
	{ID: "CIV.ORD.E004", Name: "Presentar demanda", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Actor", Description: "Ingreso del escrito inicial y anexos."},
	{ID: "CIV.ORD.E005", Name: "Demanda ingresada", Kind: "", Issuer: "", Description: ""},
	{ID: "CIV.ORD.E006", Name: "Auto de prevención", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Requiere aclarar, corregir o completar la demanda."},
	{ID: "CIV.ORD.E007", Name: "Prevención ingresada", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Actor", Description: "Se presenta respuesta al requerimiento."},
	{ID: "CIV.ORD.E008", Name: "Desahogar prevención", Kind: "", Issuer: "Juzgado", Description: "Se acepta la prevención"},
	{ID: "CIV.ORD.E009", Name: "Desechar demanda", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se rechaza la demanda."},
	{ID: "CIV.ORD.E010", Name: "Admitir demanda", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se admite y ordena traslado."},
	{ID: "CIV.ORD.E011", Name: "Preparar emplazamiento", Kind: "ACTUACIÓN_JUDICIAL", Issuer: "Juzgado/actuario", Description: "Se genera cita, exhorto, oficio o publicación."},
	{ID: "CIV.ORD.E012", Name: "Emplazamiento válido", Kind: "NOTIFICACIÓN", Issuer: "Actuario/secretario", Description: "La primera notificación queda legalmente practicada."},
	{ID: "CIV.ORD.E013", Name: "Emplazamiento fallido", Kind: "NOTIFICACIÓN_FALLIDA", Issuer: "Actuario/secretario", Description: "No se localiza o no se logra diligenciar."},
	{ID: "CIV.ORD.E014", Name: "Promover nuevo domicilio", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Actor", Description: "Se aporta información para nueva diligencia."},
	{ID: "CIV.ORD.E015", Name: "Ordenar edictos", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se autoriza notificación por publicaciones."},
	{ID: "CIV.ORD.E016", Name: "Enviar exhorto", Kind: "COMUNICACIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se solicita auxilio a otro órgano."},
	{ID: "CIV.ORD.E017", Name: "Contestar demanda", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Demandado", Description: "Se formula contestación y defensas."},
	{ID: "CIV.ORD.E018", Name: "Vencer plazo de contestación", Kind: "VENCIMIENTO", Issuer: "Sistema", Description: "Concluyen nueve días sin contestación."},
	{ID: "CIV.ORD.E019", Name: "Detectar contestación defectuosa", Kind: "VALIDACIÓN", Issuer: "Sistema/juzgado", Description: "El escrito no cumple requisitos de contestación."},
	{ID: "CIV.ORD.E020", Name: "Formular reconvención", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Demandado", Description: "Se plantea pretensión contra el actor."},
	{ID: "CIV.ORD.E021", Name: "Oponer compensación", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Demandado", Description: "Se hace valer compensación."},
	{ID: "CIV.ORD.E022", Name: "Allanarse", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Demandado", Description: "Acepta total o parcialmente las pretensiones."},
	{ID: "CIV.ORD.E023", Name: "Dar vista al actor", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se abre plazo para manifestaciones."},
	{ID: "CIV.ORD.E024", Name: "Presentar réplica", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Actor/reconventor", Description: "Se responde a planteamientos de contraparte."},
	{ID: "CIV.ORD.E025", Name: "Presentar dúplica", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Demandado/actor reconvenido", Description: "Se responde a la réplica."},
	{ID: "CIV.ORD.E026", Name: "Vencer vista", Kind: "VENCIMIENTO", Issuer: "Sistema", Description: "Concluye plazo sin promoción."},
	{ID: "CIV.ORD.E027", Name: "Fijar litis", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Queda delimitada la controversia."},
	{ID: "CIV.ORD.E028", Name: "Ofrecer pruebas", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Parte", Description: "Se identifican los medios de prueba."},
	{ID: "CIV.ORD.E029", Name: "Admitir pruebas", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se califican y admiten medios de prueba."},
	{ID: "CIV.ORD.E030", Name: "Desechar prueba", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se excluye un medio probatorio."},
	{ID: "CIV.ORD.E031", Name: "Preparar prueba", Kind: "ACTUACIÓN_PROBATORIA", Issuer: "Parte/juzgado", Description: "Se cita, requiere, designa o gira oficio."},
	{ID: "CIV.ORD.E032", Name: "Desahogar prueba", Kind: "ACTUACIÓN_PROBATORIA", Issuer: "Parte/perito/testigo/juzgado", Description: "Se incorpora el resultado probatorio."},
	{ID: "CIV.ORD.E033", Name: "Solicitar prueba fuera del lugar", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Parte", Description: "Se requiere exhorto o medio de comunicación."},
	{ID: "CIV.ORD.E034", Name: "Señalar audiencia", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se fija fecha y hora."},
	{ID: "CIV.ORD.E035", Name: "Celebrar audiencia", Kind: "AUDIENCIA", Issuer: "Juzgado/partes", Description: "Se desahogan pruebas y alegatos."},
	{ID: "CIV.ORD.E036", Name: "Diferir audiencia", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se difiere por pruebas pendientes."},
	{ID: "CIV.ORD.E037", Name: "Presentar alegatos", Kind: "DOCUMENTO_O_MANIFESTACIÓN", Issuer: "Parte", Description: "Se formulan conclusiones."},
	{ID: "CIV.ORD.E038", Name: "Cerrar instrucción", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "El asunto queda en estado de sentencia."},
	{ID: "CIV.ORD.E039", Name: "Dictar sentencia", Kind: "SENTENCIA", Issuer: "Juzgado", Description: "Se resuelve el fondo o la terminación procesal."},
	{ID: "CIV.ORD.E040", Name: "Notificar sentencia", Kind: "NOTIFICACIÓN", Issuer: "Juzgado", Description: "Se comunica la sentencia."},
	{ID: "CIV.ORD.E041", Name: "Interponer recurso", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Parte", Description: "Se promueve medio de impugnación."},
	{ID: "CIV.ORD.E042", Name: "Consentir sentencia", Kind: "DECISIÓN", Issuer: "Parte/sistema", Description: "No se interpone medio de defensa dentro del plazo aplicable."},
	{ID: "CIV.ORD.E043", Name: "Solicitar ejecución", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Parte vencedora", Description: "Se pide cumplimiento forzoso."},
	{ID: "CIV.ORD.E044", Name: "Cumplir sentencia", Kind: "CUMPLIMIENTO", Issuer: "Parte obligada", Description: "Se acredita cumplimiento."},
	{ID: "CIV.ORD.E045", Name: "Celebrar convenio", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Partes", Description: "Se acuerda terminar o modificar la controversia."},
	{ID: "CIV.ORD.E046", Name: "Desistirse de demanda", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Actor", Description: "Abandona la instancia o acción conforme al momento procesal."},
	{ID: "CIV.ORD.E047", Name: "Detectar inactividad", Kind: "CONTROL_TEMPORAL", Issuer: "Sistema", Description: "No existe promoción que impulse el procedimiento."},
	{ID: "CIV.ORD.E048", Name: "Decretar caducidad", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se declara caducidad de la instancia."},
	{ID: "CIV.ORD.E049", Name: "Promover incidente", Kind: "DOCUMENTO_PRESENTADO", Issuer: "Parte", Description: "Se abre cuestión accesoria."},
	{ID: "CIV.ORD.E050", Name: "Resolver incidente", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se decide la cuestión incidental."},
	{ID: "CIV.ORD.E051", Name: "Suspender procedimiento", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se detiene temporalmente el curso."},
	{ID: "CIV.ORD.E052", Name: "Reanudar procedimiento", Kind: "RESOLUCIÓN_JUDICIAL", Issuer: "Juzgado", Description: "Se levanta suspensión."},
	{ID: "CIV.ORD.E053", Name: "Reportar falla Tribunal Virtual", Kind: "INCIDENTE_TÉCNICO", Issuer: "Usuario/administrador", Description: "Se acredita interrupción que pudo afectar término."},
}

// SeedMachineEvents upserts MachineEvents into lexon.machine_events.
func SeedMachineEvents(ctx context.Context, pool *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for _, e := range MachineEvents {
		batch.Queue(
			`INSERT INTO lexon.machine_events (id, name, description, kind, issuer)
			 VALUES ($1, $2, $3, $4, $5)
			 ON CONFLICT (id) DO UPDATE SET
			   name = EXCLUDED.name,
			   description = EXCLUDED.description,
			   kind = EXCLUDED.kind,
			   issuer = EXCLUDED.issuer`,
			e.ID, e.Name, e.Description, e.Kind, e.Issuer,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range MachineEvents {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("seeds: insert machine event: %w", err)
		}
	}

	return nil
}
