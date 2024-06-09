package models

func GetDailyObservationStructForPlace(place string) ObservationQuery {
	query := ObservationQuery{}
	query.Id = "fmi::observations::weather::daily::timevaluepair"
	query.Place = place
	return query
}

func GetDailyObservationStructForFmisid(fmisid int) ObservationQuery {
	query := ObservationQuery{}
	query.Id = "fmi::observations::weather::daily::timevaluepair"
	query.Fmisid = fmisid
	return query
}