package reporters

import (
	"errors"
	"io/ioutil"
	"time"

	"github.com/mishamyrt/checode/v1/pkg/types"
)

func getTimeStamp() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02-15-04-05")
}

// STOPSHIP: Add description

// CreateReport lol
func CreateReport(extension string, name string, data types.ParsingResult) error {
	var reportContent string
	if len(name) == 0 {
		name = "report-" + getTimeStamp() + "." + extension
	}
	switch extension {
	case "md":
		reportContent = FormatMarkdown(data)
	default:
		return errors.New("Unknown report format")
	}
	return ioutil.WriteFile(name, []byte(reportContent), 0644)
}
