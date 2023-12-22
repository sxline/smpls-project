package model

import "github.com/sxline/smpls-project/grpc-service/internal/pb"

// DataModel represents the data model structure
type DataModel struct {
	Id         string          `json:"id"`
	Categories *DataCategories `json:"categories"`
	Title      *DataTitle      `json:"title"`
	Type       string          `json:"type"`
	Posted     float64         `json:"posted"`
}

// DataCategories represents the categories data structure
type DataCategories struct {
	Subcategory string `json:"subcategory"`
}

// DataTitle represents the title data structure
type DataTitle struct {
	Ro string `json:"ro"`
	Ru string `json:"ru"`
}

// GrpcDataToModel converts a gRPC-generated Data object to the DataModel structure
func GrpcDataToModel(p *pb.Data) DataModel {
	return DataModel{
		Id: p.GetXId(),
		Categories: &DataCategories{
			Subcategory: p.GetCategories().GetSubcategory(),
		},
		Title: &DataTitle{
			Ro: p.GetTitle().GetRo(),
			Ru: p.GetTitle().GetRu(),
		},
		Type:   p.GetType(),
		Posted: p.GetPosted(),
	}
}

// ToGrpcData converts a DataModel structure to a gRPC-generated Data object.
func ToGrpcData(p DataModel) *pb.Data {
	return &pb.Data{
		XId: p.Id,
		Categories: &pb.Data_Categories{
			Subcategory: p.Categories.Subcategory,
		},
		Title: &pb.Data_Title{
			Ro: p.Title.Ro,
			Ru: p.Title.Ru,
		},
		Type:   p.Type,
		Posted: p.Posted,
	}
}
