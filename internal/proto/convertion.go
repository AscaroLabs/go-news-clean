package proto

import (
	"go-news-clean/internal/domain/entity/news"
	"log"

	"github.com/google/uuid"
)

func (p *NewsRequestParams) ToGetDTO() (news.GetDTO, error) {

	log.Printf("[Convertor] Convert NewsRequestParams to GetDTO")

	sort_field, err := news.SortFieldFromString(p.Sort)
	if err != nil {

		log.Printf("[Convertor] Error: %v", err)

		return news.GetDTO{}, err
	}

	log.Printf("[Convertor] Sort fied: %s", sort_field)

	order, err := news.OrderFromString(p.Order)
	if err != nil {

		log.Printf("[Convertor] Error: %v", err)

		return news.GetDTO{}, err
	}

	log.Printf("[Convertor] Order: %s", order)

	var filter news.FilterDTO

	if p.Filter != nil {
		filter, err = p.Filter.ToFilterDTO()
		if err != nil {

			log.Printf("[Convertor] Error: %v", err)

			return news.GetDTO{}, err
		}
	}

	log.Printf("[Convertor] Filter: %v", filter)

	return news.GetDTO{
		Offset: p.Offset,
		Limit:  p.Limit,
		Sort:   sort_field,
		Order:  order,
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
	if s == "" {
		return uuid.Nil, nil
	}
	return uuid.FromBytes([]byte(s))
}
