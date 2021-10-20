package warnings

import (
	"time"

	"github.com/mishamyrt/checode/v1/pkg/config"
	"github.com/mishamyrt/checode/v1/pkg/dates"
	"github.com/mishamyrt/checode/v1/pkg/format"
)

func handleDue(m *Match, argument, message string) (result string) {
	result = message + format.Grey(" | ")
	date, err := dates.ParseShort(argument)
	if err != nil {
		result += format.Yellow("Wrong date format. Expected yyyy-mm-dd")
		return
	}
	if time.Now().After(date) {
		m.Flags |= config.ErrFlag
		result += format.Red("Time to fix it is up")
	} else {
		result += format.Grey("Should be fixed before " + dates.Format(date))
	}
	return result
}

var DueHandler = commandHandler{"due", handleDue}
