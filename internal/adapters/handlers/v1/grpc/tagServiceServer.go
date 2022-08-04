package grpc

import (
	"context"
	entity "go-news-clean/internal/domain/entity/tags"
	"go-news-clean/internal/domain/usecase/tags"

	"log"

	pb "go-news-clean/internal/proto"
)

type tagServiceServer struct {
	pb.UnimplementedTagServiceServer
	tagService *tags.TagService
}

func (ts *tagServiceServer) Get(ctx context.Context, r *pb.EmptyRequest) (*pb.TagList, error) {

	log.Printf("[REQ] Get all tags")

	entity_tag_list, err := ts.tagService.GetAll()
	if err != nil {
		return nil, err
	}

	// entity_tag_list := make([]*entity.Tag, 0)

	log.Printf("Enitiy tag list: %v", entity_tag_list)

	tag_list := ConvertEntityTagListToPBTagList(entity_tag_list)

	log.Printf("Enitiy tag list coverted: %v", tag_list)
	return &pb.TagList{
		Tag: tag_list,
	}, nil
}

func NewTagServiceServer(ts *tags.TagService) *tagServiceServer {
	return &tagServiceServer{
		tagService: ts,
	}
}

func ConvertEntityTagListToPBTagList(entity_tag_list []*entity.Tag) []*pb.Tag {
	tag_list := make([]*pb.Tag, 0)
	for _, t := range entity_tag_list {
		tag_list = append(tag_list, ConvertEntityTagToPBTag(t))
	}
	return tag_list
}

func ConvertEntityTagToPBTag(t *entity.Tag) *pb.Tag {
	return &pb.Tag{
		Id:   t.Id.String(),
		Name: t.Name,
	}
}
