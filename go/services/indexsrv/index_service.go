package indexsrv

import (
	"elkstack/helper"
	"fmt"
)

type indexService struct {
}

func NewIndexService() IndexService {
	return &indexService{}
}

func (obj indexService) List() (response []IndexResponse, err error) {

	indices := []index{}

	err = helper.HttpGet("http://localhost:9200/_cat/indices?format=json", &indices)

	if err != nil {
		return
	}

	for _, v := range indices {
		fmt.Printf("%v\n", v.Index)

		response = append(response, IndexResponse{
			Index: v.Index,
		})
	}

	return
}

func (obj indexService) Search(index string) error {

	fmt.Printf("index: %v \n", index)

	search := searchResult{}

	err := helper.HttpGet(fmt.Sprintf("http://localhost:9200/%v/_search", index), &search)

	if err != nil {
		return err
	}

	fmt.Printf("%#v", search)

	return nil
}
