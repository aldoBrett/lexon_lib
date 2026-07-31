package seeds

// MachineStateTransition mirrors a row of lexon.machine_state_transitions.
//
// SourceState and TargetState are not always real lexon.machine_states IDs:
// "*" denotes a wildcard source (any state), and values such as
// FUERA_DE_ALCANCE, FUJO_RECURSO, ALERTA_CADUCIDAD, ESTADO_PRINCIPAL_PREVIO,
// ESTADO_PREVIO and ESTADO_ACTUAL are pseudo-state exit markers rather than
// rows in machine_states. Loading this data therefore requires resolving
// those cases before/instead of relying on the table's foreign keys.
type MachineStateTransition struct {
	ID          string
	SourceState string
	TargetState string
	Condition   string
	Actions     string
	Risk        string
	Note        string
}

// MachineStateTransitions is the seed data for the CIV.ORD (civil ordinario) machine state transitions.
var MachineStateTransitions = []MachineStateTransition{
	{ID: "T001", SourceState: "CIV.ORD.S00", Condition: "pretension_registrada = TRUE", TargetState: "CIV.ORD.S00", Actions: "Guardar pretensión; solicitar documentos base", Risk: "Bajo", Note: "No crea término"},
	{ID: "T002", SourceState: "CIV.ORD.S00", Condition: "via_ordinaria = TRUE", TargetState: "CIV.ORD.S01", Actions: "Crear checklist de demanda ordinaria", Risk: "Medio", Note: "Art. 638"},
	{ID: "T003", SourceState: "CIV.ORD.S00", Condition: "via_ordinaria = FALSE", TargetState: "FUERA_DE_ALCANCE", Actions: "Desviar a procedimiento aplicable", Risk: "Alto", Note: "No forzar flujo ordinario"},
	{ID: "T004", SourceState: "CIV.ORD.S01", Condition: "demanda_y_anexos_registrados = TRUE", TargetState: "CIV.ORD.S02", Actions: "Crear expediente; conservar acuse y fecha/hora", Risk: "Alto", Note: "Arts. 612, 614, 623"},
	{ID: "T005", SourceState: "CIV.ORD.S02", Condition: "prevencion_emitida = TRUE", TargetState: "CIV.ORD.S04", Actions: "Crear tarea; calcular plazo indicado; elevar riesgo", Risk: "Crítico", Note: "Arts. 616, 621"},
	{ID: "T006", SourceState: "CIV.ORD.S04", Condition: "presentado_en_plazo = TRUE AND cumplimiento_suficiente = TRUE", TargetState: "CIV.ORD.S03", Actions: "Adjuntar escrito; solicitar nueva revisión", Risk: "Alto", Note: "Revisión judicial"},
	{ID: "T007", SourceState: "CIV.ORD.S04", Condition: "presentado_en_plazo = TRUE AND cumplimiento_suficiente = UNKNOWN", TargetState: "CIV.ORD.S03", Actions: "Marcar revisión humana obligatoria", Risk: "Alto", Note: "No decidir suficiencia automáticamente"},
	{ID: "T008", SourceState: "CIV.ORD.S04", Condition: "desechamiento_notificado = TRUE", TargetState: "CIV.ORD.S31", Actions: "Cerrar plazo; habilitar análisis de apelación", Risk: "Crítico", Note: "Arts. 621-622"},
	{ID: "T009", SourceState: "CIV.ORD.S02", Condition: "admision_emitida = TRUE", TargetState: "CIV.ORD.S05", Actions: "Crear misión de emplazamiento", Risk: "Medio", Note: "Arts. 624, 639"},
	{ID: "T010", SourceState: "CIV.ORD.S05", Condition: "modalidad_emplazamiento_definida = TRUE", TargetState: "CIV.ORD.S06", Actions: "Registrar modalidad y responsable", Risk: "Alto", Note: "Personal/exhorto/edictos/autoridad"},
	{ID: "T011", SourceState: "CIV.ORD.S06", Condition: "emplazamiento_valido = TRUE", TargetState: "CIV.ORD.S07", Actions: "Guardar acta; activar plazo de contestación", Risk: "Crítico", Note: "Arts. 624-628, 639"},
	{ID: "T012", SourceState: "CIV.ORD.S06", Condition: "emplazamiento_valido = FALSE", TargetState: "CIV.ORD.S06", Actions: "Crear tarea de corrección; no iniciar contestación", Risk: "Crítico", Note: "Evitar cómputo prematuro"},
	{ID: "T013", SourceState: "CIV.ORD.S06", Condition: "nuevo_domicilio_registrado = TRUE", TargetState: "CIV.ORD.S06", Actions: "Programar nueva diligencia", Risk: "Alto", Note: "Nueva cita"},
	{ID: "T014", SourceState: "CIV.ORD.S06", Condition: "domicilio_desconocido_acreditado = TRUE", TargetState: "CIV.ORD.S06", Actions: "Crear checklist de publicaciones", Risk: "Alto", Note: "Efectos según última publicación aplicable"},
	{ID: "T015", SourceState: "CIV.ORD.S06", Condition: "fuera_jurisdiccion = TRUE", TargetState: "CIV.ORD.S06", Actions: "Crear seguimiento de exhorto", Risk: "Alto", Note: "Art. 626 y reglas de comunicaciones"},
	{ID: "T016", SourceState: "CIV.ORD.S07", Condition: "fecha_emplazamiento_confirmada = TRUE", TargetState: "CIV.ORD.S08", Actions: "Calcular nueve días; alertas escalonadas", Risk: "Crítico", Note: "Art. 639"},
	{ID: "T017", SourceState: "CIV.ORD.S08", Condition: "presentado_en_plazo = TRUE AND requisitos_contestacion = TRUE", TargetState: "CIV.ORD.S09", Actions: "Clasificar defensas, hechos y anexos", Risk: "Alto", Note: "Arts. 629-630"},
	{ID: "T018", SourceState: "CIV.ORD.S08", Condition: "contestacion_valida = FALSE", TargetState: "CIV.ORD.S10", Actions: "Registrar contestación en sentido negativo", Risk: "Crítico", Note: "Art. 631"},
	{ID: "T019", SourceState: "CIV.ORD.S09", Condition: "requisitos_contestacion = FALSE", TargetState: "CIV.ORD.S10", Actions: "Advertir que puede tenerse por no contestada", Risk: "Crítico", Note: "Art. 632"},
	{ID: "T020", SourceState: "CIV.ORD.S09", Condition: "reconvencion = TRUE", TargetState: "CIV.ORD.S12", Actions: "Crear plazo de nueve días para contestar reconvención", Risk: "Crítico", Note: "Arts. 635, 637"},
	{ID: "T021", SourceState: "CIV.ORD.S09", Condition: "compensacion = TRUE", TargetState: "CIV.ORD.S11", Actions: "Dar vista y vincular a sentencia conjunta", Risk: "Alto", Note: "Arts. 629, 635, 637"},
	{ID: "T022", SourceState: "CIV.ORD.S09", Condition: "allanamiento_total = TRUE", TargetState: "CIV.ORD.S20", Actions: "Solicitar revisión judicial de efectos y sentencia", Risk: "Alto", Note: "Intervención judicial obligatoria"},
	{ID: "T023", SourceState: "CIV.ORD.S09", Condition: "vista_ordenada = TRUE", TargetState: "CIV.ORD.S11", Actions: "Calcular plazo de tres días", Risk: "Alto", Note: "Flujo ordinario de vistas"},
	{ID: "T024", SourceState: "CIV.ORD.S11", Condition: "replica_presentada = TRUE", TargetState: "CIV.ORD.S14", Actions: "Crear plazo de dúplica de tres días", Risk: "Alto", Note: "Flujo postulatorio"},
	{ID: "T025", SourceState: "CIV.ORD.S11", Condition: "replica_presentada = FALSE", TargetState: "CIV.ORD.S15", Actions: "Cerrar oportunidad y continuar", Risk: "Medio", Note: "Preclusión"},
	{ID: "T026", SourceState: "CIV.ORD.S12", Condition: "contestacion_reconvencion_en_plazo = TRUE", TargetState: "CIV.ORD.S13", Actions: "Crear réplica correspondiente", Risk: "Alto", Note: "Nueve días"},
	{ID: "T027", SourceState: "CIV.ORD.S12", Condition: "contestacion_reconvencion = FALSE", TargetState: "CIV.ORD.S13", Actions: "Registrar falta y continuar", Risk: "Crítico", Note: "Consecuencia procesal a validar"},
	{ID: "T028", SourceState: "CIV.ORD.S13", Condition: "replica_presentada = TRUE", TargetState: "CIV.ORD.S14", Actions: "Calcular tres días de dúplica", Risk: "Alto", Note: "Tres días"},
	{ID: "T029", SourceState: "CIV.ORD.S14", Condition: "duplica_presentada = TRUE", TargetState: "CIV.ORD.S15", Actions: "Cerrar etapa postulatoria", Risk: "Medio", Note: "Tres días"},
	{ID: "T030", SourceState: "CIV.ORD.S14", Condition: "duplica_presentada = FALSE", TargetState: "CIV.ORD.S15", Actions: "Cerrar oportunidad y continuar", Risk: "Medio", Note: "Preclusión"},
	{ID: "T031", SourceState: "CIV.ORD.S15", Condition: "litis_fijada = TRUE", TargetState: "CIV.ORD.S16", Actions: "Generar matriz de hechos controvertidos y pruebas", Risk: "Medio", Note: "Inicio fase probatoria"},
	{ID: "T032", SourceState: "CIV.ORD.S16", Condition: "prueba_identificada = TRUE", TargetState: "CIV.ORD.S16", Actions: "Registrar medio, objeto y oferente", Risk: "Medio", Note: "Una fila por prueba"},
	{ID: "T033", SourceState: "CIV.ORD.S16", Condition: "prueba_admitida = TRUE", TargetState: "CIV.ORD.S16", Actions: "Crear subtareas de preparación", Risk: "Alto", Note: "Por tipo de prueba"},
	{ID: "T034", SourceState: "CIV.ORD.S16", Condition: "prueba_admitida = FALSE", TargetState: "CIV.ORD.S16", Actions: "Cerrar medio; habilitar análisis de impugnación", Risk: "Alto", Note: "No eliminar evidencia histórica"},
	{ID: "T035", SourceState: "CIV.ORD.S16", Condition: "preparacion_completa = TRUE", TargetState: "CIV.ORD.S17", Actions: "Marcar lista para desahogo", Risk: "Medio", Note: "Citas, oficios, peritos, exhortos"},
	{ID: "T036", SourceState: "CIV.ORD.S16", Condition: "fuera_lugar_juicio = TRUE", TargetState: "CIV.ORD.S16", Actions: "Crear seguimiento de comunicación judicial", Risk: "Alto", Note: "Art. 643"},
	{ID: "T037", SourceState: "CIV.ORD.S17", Condition: "resultado_prueba_registrado = TRUE", TargetState: "CIV.ORD.S17", Actions: "Guardar acta, dictamen o documento", Risk: "Medio", Note: "Actualizar porcentaje de preparación"},
	{ID: "T038", SourceState: "CIV.ORD.S17", Condition: "fecha_audiencia_confirmada = TRUE", TargetState: "CIV.ORD.S18", Actions: "Agregar calendario y recordatorios", Risk: "Alto", Note: "Audiencia"},
	{ID: "T039", SourceState: "CIV.ORD.S18", Condition: "todas_pruebas_desahogadas = TRUE", TargetState: "CIV.ORD.S20", Actions: "Registrar alegatos y cierre", Risk: "Alto", Note: "Art. 644"},
	{ID: "T040", SourceState: "CIV.ORD.S18", Condition: "pruebas_pendientes = TRUE", TargetState: "CIV.ORD.S19", Actions: "Calcular continuación máximo ocho días", Risk: "Alto", Note: "Art. 644"},
	{ID: "T041", SourceState: "CIV.ORD.S19", Condition: "continuacion_celebrada = TRUE", TargetState: "CIV.ORD.S20", Actions: "Cerrar audiencia", Risk: "Alto", Note: "Art. 644"},
	{ID: "T042", SourceState: "CIV.ORD.S20", Condition: "instruccion_cerrada = TRUE", TargetState: "CIV.ORD.S21", Actions: "Crear control judicial de quince días", Risk: "Medio", Note: "Art. 644"},
	{ID: "T043", SourceState: "CIV.ORD.S21", Condition: "sentencia_emitida = TRUE", TargetState: "CIV.ORD.S22", Actions: "Clasificar sentido y puntos resolutivos", Risk: "Crítico", Note: "Arts. 644-645"},
	{ID: "T044", SourceState: "CIV.ORD.S22", Condition: "notificacion_valida = TRUE", TargetState: "CIV.ORD.S23", Actions: "Crear decisión: impugnar o ejecutar", Risk: "Crítico", Note: "Plazo del medio aplicable"},
	{ID: "T045", SourceState: "CIV.ORD.S23", Condition: "recurso_presentado = TRUE", TargetState: "FUJO_RECURSO", Actions: "Crear expediente relacionado de impugnación", Risk: "Crítico", Note: "Desviar a BMEP de recurso"},
	{ID: "T046", SourceState: "CIV.ORD.S23", Condition: "sentencia_firme = TRUE", TargetState: "CIV.ORD.S24", Actions: "Habilitar ejecución", Risk: "Alto", Note: "Firmeza confirmada"},
	{ID: "T047", SourceState: "CIV.ORD.S24", Condition: "ejecucion_solicitada = TRUE", TargetState: "CIV.ORD.S24", Actions: "Crear flujo de requerimiento, embargo o cumplimiento", Risk: "Alto", Note: "BMEP de ejecución"},
	{ID: "T048", SourceState: "CIV.ORD.S24", Condition: "cumplimiento_total = TRUE", TargetState: "CIV.ORD.S27", Actions: "Cerrar expediente con evidencia", Risk: "Bajo", Note: "Estado terminal"},
	{ID: "T049", SourceState: "*", Condition: "convenio_valido_y_aprobado = TRUE", TargetState: "CIV.ORD.S28", Actions: "Cerrar términos pendientes y conservar convenio", Risk: "Medio", Note: "Transición transversal"},
	{ID: "T050", SourceState: "CIV.ORD.S01", Condition: "desistimiento_presentado = TRUE", TargetState: "CIV.ORD.S29", Actions: "Cerrar preparación", Risk: "Bajo", Note: "Antes del emplazamiento"},
	{ID: "T051", SourceState: "CIV.ORD.S07", Condition: "consentimiento_demandado = TRUE OR desistimiento_accion = TRUE", TargetState: "CIV.ORD.S29", Actions: "Registrar costas/efectos según supuesto", Risk: "Alto", Note: "Art. 3"},
	{ID: "T052", SourceState: "*", Condition: "inactividad_dias_naturales >= 120 AND primera_instancia = TRUE AND impedimento_suspensivo = FALSE AND sentencia_ordenada = FALSE", TargetState: "ALERTA_CADUCIDAD", Actions: "Alerta crítica; revisión humana", Risk: "Crítico", Note: "Art. 3"},
	{ID: "T053", SourceState: "*", Condition: "caducidad_decretada = TRUE", TargetState: "CIV.ORD.S30", Actions: "Cerrar instancia; conservar posibilidad de nueva acción si procede", Risk: "Crítico", Note: "Art. 3"},
	{ID: "T054", SourceState: "*", Condition: "incidente_admitido = TRUE", TargetState: "CIV.ORD.S26", Actions: "Crear subflujo paralelo sin borrar estado principal", Risk: "Alto", Note: "Incidente"},
	{ID: "T055", SourceState: "CIV.ORD.S26", Condition: "incidente_resuelto = TRUE", TargetState: "ESTADO_PRINCIPAL_PREVIO", Actions: "Cerrar subflujo y ejecutar efectos", Risk: "Alto", Note: "Retorno controlado"},
	{ID: "T056", SourceState: "*", Condition: "suspension_decretada = TRUE", TargetState: "CIV.ORD.S25", Actions: "Congelar términos afectados; conservar no afectados", Risk: "Crítico", Note: "Requiere regla individual"},
	{ID: "T057", SourceState: "CIV.ORD.S25", Condition: "reanudacion_decretada = TRUE", TargetState: "ESTADO_PREVIO", Actions: "Recalcular términos conforme resolución", Risk: "Crítico", Note: "No asumir reanudación automática"},
	{ID: "T058", SourceState: "*", Condition: "falla_confirmada_por_administrador = TRUE", TargetState: "ESTADO_ACTUAL", Actions: "Adjuntar reporte y someter efecto al tribunal", Risk: "Crítico", Note: "Art. 75"},
}
