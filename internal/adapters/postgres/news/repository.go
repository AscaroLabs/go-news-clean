package news

import (
	"context"
	"fmt"
	"go-news-clean/internal/domain/entity/news"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type newsRepository struct {
	connStr string
}

func NewNewsRepository(connStr string) *newsRepository {
	return &newsRepository{
		connStr: connStr,
	}
}

// news.News возвращаются без файлов и тегов
func (nr *newsRepository) GetNews(gdto news.GetDTO, user_id uuid.UUID) ([]*news.News, error) {
	pool, err := pgxpool.Connect(context.Background(), nr.connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	q := makeGetQuery(gdto)

	rows, err := pool.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news_list := make([]*news.News, 0)

	for rows.Next() {
		var n news.News
		err := rows.Scan(
			&n.Id,
			&n.Title,
			&n.Author,
			&n.Active,
			&n.ActiveFrom,
			&n.Text,
			&n.TextJson,
			&n.UserId,
			&n.IsImportant,
		)
		if err != nil {
			return nil, err
		}
		news_list = append(news_list, &n)
	}

	return news_list, nil
}

func makeGetQuery(gdto news.GetDTO) string {
	return fmt.Sprintf(
		`
		select id,title,author,active,activeFrom,text,textJSON,userId,isImportant from news
		%s order by %s %s limit %d offset %d`,
		makeGetFilterPart(gdto.Filter),
		gdto.Sort,
		gdto.Order,
		gdto.Limit,
		gdto.Offset,
	)
}

func makeGetFilterPart(fdto news.FilterDTO) string {

	var user_part string
	var mode_part string
	switch fdto.UserId {
	case uuid.Nil:
		user_part = "true"
	default:
		user_part = fmt.Sprintf("news.userId=%s", fdto.UserId.String())
	}
	switch fdto.Mode {
	case news.FilterModeActive:
		mode_part = "news.active=true"
	case news.FilterModeInactive:
		mode_part = "news.active=false"
	default:
		mode_part = "true"
	}

	return fmt.Sprintf("where %s and %s", user_part, mode_part)
}

// news.News возвращается без файлов и тегов
func (nr *newsRepository) GetOne(id uuid.UUID) (*news.News, error) {
	pool, err := pgxpool.Connect(context.Background(), nr.connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(
		context.Background(),
		`select id,title,author,active,activeFrom,text,textJSON,userId,isImportant from news
		where id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var n news.News
	err = rows.Scan(
		&n.Id,
		&n.Title,
		&n.Author,
		&n.Active,
		&n.ActiveFrom,
		&n.Text,
		&n.TextJson,
		&n.UserId,
		&n.IsImportant,
	)
	if err != nil {
		return nil, err
	}

	return &n, nil
}

// news.News возвращается без файлов и тегов
func (nr *newsRepository) GetOneBySlug(objectSlug string) (*news.News, error) {
	return nil, nil
}

func (nr *newsRepository) Create(cdto news.CreateDTO) error {
	return nil
}

func (nr *newsRepository) Update(udto news.UpdateDTO) error {
	return nil
}

func (nr *newsRepository) Delete(ddto news.DeleteDTO) error {
	return nil
}
