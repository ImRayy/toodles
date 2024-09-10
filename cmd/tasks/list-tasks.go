package tasks

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type TableFields string

const (
	Pending   TableFields = "pending"
	Done      TableFields = "done"
	Cancelled TableFields = "cancelled"
	NotDone   TableFields = "not_done"
)

const (
	border_color = lipgloss.Color("#90a4ae")
	text         = lipgloss.Color("#d0d0d0")
	subtext      = lipgloss.Color("#cdd6f4")
)

// NOTE: This table isn't dynamic, means if columns like status showing colored
// text it's only gonna work only if col number of "status" is == 2 (by index)
// I can make it dynamic but not going to cause not worth it, maybe I'll think
// about it if data structure gets complicated

func RenderTable(rows [][]string, headers []string, priority []string) {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		HeaderStyle = re.NewStyle().Foreground(border_color).
				Bold(true).Align(lipgloss.Center)
		CellStyle    = re.NewStyle().Padding(0, 1)
		EvenRowStyle = CellStyle.Foreground(text)
		OddRowStyle  = CellStyle.Foreground(lipgloss.Color(subtext))
		BorderStyle  = lipgloss.NewStyle().Foreground(border_color)
	)

	statusColors := map[string]lipgloss.Color{
		" Pending": lipgloss.Color("#FF875F"),
		" Done":    lipgloss.Color("#00E2C7"),
	}

	priorityColors := map[string]lipgloss.Color{
		"mid":  lipgloss.Color("#FDFF90"),
		"high": lipgloss.Color("#FF7698"),
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}

			switch col {
			case 0:
				width := (func() int {
					idLen := strings.Count(rows[row-1][0], "") - 2
					return 3 + idLen
				})()
				style = style.Width(width).Align(lipgloss.Center)
			case 1:
				color := priorityColors[fmt.Sprint(priority[row-1])]
				style = CellStyle.Foreground(color)
			case 2:
				color := statusColors[fmt.Sprint(rows[row-1][2])]
				style = CellStyle.Foreground(color)
			}

			return style
		}).
		Headers(headers...).
		Rows(rows...)

	fmt.Println(t)
}
