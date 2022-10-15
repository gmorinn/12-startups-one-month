package utils

import (
	db "12-startups-one-month/internal"
)

func RemoveDuplicateViewer(viewers []db.Viewer) []db.Viewer {
	keys := make(map[string]bool)
	list := []db.Viewer{}
	for _, entry := range viewers {
		if _, value := keys[entry.UserIDViewer.String()]; !value {
			keys[entry.UserIDViewer.String()] = true
			list = append(list, entry)
		}
	}
	return list
}
