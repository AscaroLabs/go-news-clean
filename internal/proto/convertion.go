package proto

import (
	"go-news-clean/internal/domain/entity/news"

	"github.com/google/uuid"
)

func (p *NewsRequestParams) ToGetDTO() (news.GetDTO, error) {
	sort_field, err := news.SortFieldFromString(p.Sort)
	if err != nil {
		return news.GetDTO{}, err
	}
	filter, err := p.Filter.ToFilterDTO()
	if err != nil {
		return news.GetDTO{}, err
	}
	return news.GetDTO{
		Offset: p.Offset,
		Limit:  p.Limit,
		Sort:   sort_field,
		Order:  p.Order,
		Query:  p.Query,
		Filter: filter,
	}, nil
}

func (l *ListRequestFilter) ToFilterDTO() (news.FilterDTO, error) {
	mode, err := news.FilterModeFromSting(l.Mode)
	if err != nil {
		return news.FilterDTO{}, err
	}

	id, err := FromStringToUUID(l.UserId)
	if err != nil {
		return news.FilterDTO{}, err
	}
	return news.FilterDTO{
		UserId: id,
		Mode:   mode,
	}, nil
}

func FromStringToUUID(s string) (uuid.UUID, error) {
	return uuid.FromBytes([]byte(s))
}
