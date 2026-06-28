package presenter

import "trec/internal/domain"

func MapFromOrderBy(by domain.OrderBy) string {
	return string(by)
}
