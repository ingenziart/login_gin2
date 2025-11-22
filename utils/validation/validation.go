package validation

import "github.com/ingenziart/myapp/models"

func IsValidateStatus(s models.Status) bool {
	switch s {
	case models.StatusActive, models.StatusInactive, models.StatusDeleted:
		return true

	}
	return false

}
