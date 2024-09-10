package utils

import (
	"fmt"
	"time"
)

func FormatTime(pastTime time.Time) string {
	now := time.Now()
	duration := now.Sub(pastTime)

	var (
		secs   = int(duration.Seconds())
		mins   = int(duration.Minutes())
		hours  = duration.Hours()
		days   = int(hours / 24)
		weeks  = int(hours / (24 * 7))
		months = int(hours / (24 * 30))
		years  = int(hours / (24 * 365))
		icon   = "󰔚"
	)

	switch {
	case secs < 60:
		return fmt.Sprintf("%s %d seconds ago", icon, secs)
	case mins < 60:
		return fmt.Sprintf("%s %d minutes ago", icon, mins)
	case hours < 24:
		return fmt.Sprintf("%s %d hours ago", icon, int(hours))
	case days < 7:
		return fmt.Sprintf("%s %d days ago", icon, days)
	case weeks < 4:
		return fmt.Sprintf("%s %d weeks ago", icon, weeks)
	case months < 12:
		return fmt.Sprintf("%s %d months ago", icon, months)
	default:
		return fmt.Sprintf("%s %d years ago", icon, years)
	}
}
