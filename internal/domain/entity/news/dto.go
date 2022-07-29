package news

import (
	"errors"

	"github.com/google/uuid"
)

// type NewsRequestParams struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// Номер строки, которой начинается выборка
// 	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
// 	// Количество возвращаемых объектов на странице
// 	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
// 	// Поле для сортировки (active_from, date_create)
// 	Sort string `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
// 	// Направление сортировки (desc, asc)
// 	Order string `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
// 	// Поиск по строке
// 	Query string `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
// 	// Параметры фильтрации
// 	Filter *ListRequestFilter `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"` // Объект фильтарции
// }

// Параметры фильтрации
// type ListRequestFilter struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// фильтр по пользователю
// 	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
// 	// фильтр по активности
// 	//   1. не передано - опубликованные и черновики
// 	//   2. active - только опубликованные
// 	//   3. inactive черновики)
// 	Mode string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
// }

var (
	ErrBadArgument error = errors.New("bad argument")
)

type SortField string

const (
	FieldActiveFrom SortField = "active_from"
	FieldDateCreate SortField = "date_create"
)

func SortFieldFromString(s string) (SortField, error) {
	switch s {
	case "active_from":
		return FieldActiveFrom, nil
	case "date_create":
		return FieldDateCreate, nil
	default:
		return SortField(""), ErrBadArgument
	}
}

type GetDTO struct {
	Offset int64
	Limit  int64
	Sort   SortField
	Order  string
	Query  string
	Filter FilterDTO
}

type FilterMode string

const (
	FilterModeEmpty    FilterMode = ""
	FilterModeActive   FilterMode = "active"
	FilterModeInactive FilterMode = "inactive"
)

func FilterModeFromSting(s string) (FilterMode, error) {
	switch s {
	case "active":
		return FilterModeActive, nil
	case "inactive":
		return FilterModeInactive, nil
	case "":
		return FilterModeEmpty, nil
	default:
		return FilterMode(""), ErrBadArgument
	}
}

type FilterDTO struct {
	UserId uuid.UUID
	Mode   FilterMode
}

type CreateDTO struct {
}

type UpdateDTO struct {
	Id uuid.UUID
}

type DeleteDTO struct {
	Id uuid.UUID
}
