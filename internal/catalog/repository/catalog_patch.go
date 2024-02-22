package repository

import (
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/catalog/domain"
)

// PatchToDBProductGroup patches ProductGroup into the database model
func PatchToDBProductGroup(domain *domain.ProductGroup, model *prophet_19_1_3668.ProductGroup) error {

	if domain.Name != "" {
		model.ProductGroupDesc = domain.Name
	}

	return nil

}
