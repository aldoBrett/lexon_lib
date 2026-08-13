package engine

import "time"

type ExportHistoryParams struct {
	LegalProcedureID    string
	MachineStateStageID *string
}

// ExportHistoryEvent carries the minimal event data needed by the UI timeline.
type ExportHistoryEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExportHistoryItem represents a single transition in the machine instance's history.
type ExportHistoryItem struct {
	ID        string             `json:"id"`
	From      string             `json:"from"`
	To        string             `json:"to"`
	Event     ExportHistoryEvent `json:"event"`
	CreatedAt time.Time          `json:"createdAt"`
}

// This function will export the transitions history with data necessary for the UI
// to plot the timeline.
// Each struct in the slice should have the following:
// - id: the transition id
// - from: the source state id
// - to: the target state id
// - event: the event that triggered the transition, with its id and name
// - createdAt: the timestamp when the transition occurred
func (e *EngineHandler) ExportHistory(params ExportHistoryParams) ([]ExportHistoryItem, error) {
	machineInstance, err := e.repos.MachineInstances.GetMachineInstanceByLegalProcedureID(params.LegalProcedureID)
	if err != nil {
		return nil, err
	}

	var history []*MachineStateTransitionHistory
	if params.MachineStateStageID != nil {
		history, err = e.repos.MachineTransitionsHistory.GetMachineStateTransitionsHistoryByMachineStateStageID(machineInstance.ID, *params.MachineStateStageID)
	} else {
		history, err = e.repos.MachineTransitionsHistory.GetMachineStateTransitionsHistoryByMachineInstanceID(machineInstance.ID)
	}
	if err != nil {
		return nil, err
	}
	if len(history) == 0 {
		return []ExportHistoryItem{}, nil
	}

	transitionIDs := make([]string, len(history))
	for i, h := range history {
		transitionIDs[i] = h.TransitionID
	}

	transitions, err := e.repos.MachineTransitions.GetMachineStateTransitionsByHistoryIDs(transitionIDs)
	if err != nil {
		return nil, err
	}
	transitionsByID := make(map[string]*MachineStateTransition, len(transitions))
	eventIDsSeen := make(map[string]bool)
	eventIDs := make([]string, 0, len(transitions))
	for _, t := range transitions {
		transitionsByID[t.ID] = t
		if !eventIDsSeen[t.EventID] {
			eventIDsSeen[t.EventID] = true
			eventIDs = append(eventIDs, t.EventID)
		}
	}

	events, err := e.repos.MachineEvents.GetMachineEventsByIDs(eventIDs)
	if err != nil {
		return nil, err
	}
	eventsByID := make(map[string]*MachineEvent, len(events))
	for _, ev := range events {
		eventsByID[ev.ID] = ev
	}

	result := make([]ExportHistoryItem, 0, len(history))
	for _, h := range history {
		transition, ok := transitionsByID[h.TransitionID]
		if !ok {
			continue
		}

		item := ExportHistoryItem{
			ID:        transition.ID,
			From:      transition.SourceStateID,
			To:        transition.TargetStateID,
			CreatedAt: h.CreatedAt,
			Event: ExportHistoryEvent{
				ID: transition.EventID,
			},
		}
		if event, ok := eventsByID[transition.EventID]; ok {
			item.Event.Name = event.Name
		}

		result = append(result, item)
	}

	return result, nil
}
