package book

import "time"

type Book struct {
	ID				string		`json:"id"`
	BarCode			string		`json:"barCode"`
	Number			string		`json:"number"`
	ClassNumber		string		`json:"classNumber"`
	Class			string 		`json:"class"`
	Name			string		`json:"name"`
	Author			string		`json:"author"`
	Publisher		string		`json:"publisher"`
	PageNumber		int32		`json:"pageNumber"`
	Size			int32		`json:"size"`
	Price			float64		`json:"price"`
	Location		string		`json:"location"`
	CreationTime	time.Time	`json:"creationTime"`
	Annotation		string		`json:"annotation"`
}