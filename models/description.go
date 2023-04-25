package models

import (
	"DAB-SSH/styling"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

const useHighPerformanceRenderer = false

type DescriptionPage struct {
	projectName             string         // Watermark in top left corner of page
	summary                 string         // Short summary of project at top of page
	description             string         // Actual description of the project
	help                    help.Model     // Help bar at bottom of page
	termWidth, termHeight   int            // Size of the terminal
	modelWidth, modelHeight int            // Size of the model (not including help model)
	content                 string         // Just temporary holding of the description
	ready                   bool           // Bool if the program is ready for page scroller
	viewport                viewport.Model // Pager to fit the summarizing text
}

func CreateDescriptionPage(projectAddress int, projectName string, summary string) DescriptionPage {

	// Sets the description passed through
	description := Descriptions[projectAddress]

	// Return our created model
	return DescriptionPage{
		projectName: projectName,
		summary:     summary,
		description: description,
	}
}

// Initializes our struct as a bubble tea model
func (d DescriptionPage) Init() tea.Cmd {
	return nil
}

// Updates our model everytime a key event happens, mainly window resizes and key presses
func (d DescriptionPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets cmd as a tea command that can be easily changed later
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// Sets msg as a switch for all events
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(d.headerView())
		footerHeight := lipgloss.Height(d.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !d.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			d.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			d.viewport.YPosition = headerHeight
			d.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			d.viewport.SetContent(d.content)
			d.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			d.viewport.YPosition = headerHeight + 1
		} else {
			d.viewport.Width = msg.Width
			d.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(d.viewport))
		}

		// All key presses
	case tea.KeyMsg:

		// Converts the press into a string
		switch msg.String() {

		// Back to project page
		case "esc":
			return CreateProjectPage(), nil

		}
	}

	// Handle keyboard and mouse events in the viewport
	d.viewport, cmd = d.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return d, tea.Batch(cmds...)

	//return d, cmd
}

// Renders our model formatted to be viewed, then returns as a string
func (d DescriptionPage) View() string {

	// Our s string to build our model
	var s string

	// Size to return our model later
	var width, height int

	// Logic for setting terminal width to not break model
	if d.termWidth <= d.modelWidth {
		width = d.modelWidth
	} else {
		width = d.termWidth
	}

	// Logic for setting terminal height to not break model
	if d.termHeight <= d.modelHeight {
		height = d.modelHeight
	} else {
		height = d.termHeight
	}

	// RENDERING OUR MODEL |*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|*|
	// Addds the watermark
	s += d.projectName + "\n\n"

	s += d.summary

	s += "\n\n\n"

	s += d.headerView() + "\n"
	s += d.viewport.View() + "\n"
	s += d.footerView()

	s += d.description

	return styling.BorderStyle.Width(width).Height(height).Render(s)
}

func (d DescriptionPage) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, d.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (d DescriptionPage) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", d.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, d.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Descriptions = []string{
	"deez",
	"coconuts",
	"rule",
	"yo"}
