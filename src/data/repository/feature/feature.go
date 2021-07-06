package feature

import (
	Model "sawonersm/geonames-processor/data/model"
	Di "sawonersm/geonames-processor/kernel/di"
)

func GetByCode(di *Di.Di, code string) *Model.Feature {
	feature := &Model.Feature{}
	di.GetDb().GetConnection().Where("code = ?", code).Find(feature)

	if feature.Id == 0 {
		return nil
	}

	return feature
}
