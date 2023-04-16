package tui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type modelState int

const (
	selectionState modelState = iota
	progressState
)

type Model struct {
	state     modelState
	selection tea.Model
	progress  tea.Model
}

func NewModel() Model {
	return Model{
		selection: newSelectionModel(),
		progress:  newProgress(),
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case itemSelectionMsg:
		m.state = progressState
		cmd = m.progressSelectedCmd()
	case progressSelectedMsg:
		m.progress, cmd = m.progress.Update(msg)
	case progressIncrementMsg:
		m.progress, cmd = m.progress.Update(msg)
	case progress.FrameMsg:
		m.progress, cmd = m.progress.Update(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			if m.state == selectionState {
				m.selection, cmd = m.selection.Update(msg)
			}
			if m.state == progressState {
				m.progress, cmd = m.progress.Update(msg)
			}
		}
	}

	return m, tea.Batch(cmd)
}

func (m Model) View() string {
	switch m.state {
	case progressState:
		return m.progress.View()
	default:
		return m.selection.View()
	}
}

func (m Model) progressSelectedCmd() tea.Cmd {
	return func() tea.Msg {
		return progressSelectedMsg{}
	}
}
