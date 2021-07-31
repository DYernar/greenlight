package data

import "greenlight.yernar.net/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, filter Filters) {

	v.Check(filter.Page > 0, "page", "must be greater than 0")
	v.Check(filter.Page < 10_000_000, "page", "must be a maximum of 10 million")
	v.Check(filter.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(filter.PageSize <= 100, "page_size", "must be a maximum of 100")

	v.Check(validator.In(filter.Sort, filter.SortSafeList...), "sort", "invalid sort value")

}
