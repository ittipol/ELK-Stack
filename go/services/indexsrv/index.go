package indexsrv

type index struct {
	Health       string `json:"health,omitempty"`
	Status       string `json:"status,omitempty"`
	Index        string `json:"index,omitempty"`
	Uuid         string `json:"uuid,omitempty"`
	Pri          string `json:"pri,omitempty"`
	Rep          string `json:"rep,omitempty"`
	DocsCount    string `json:"docs.count,omitempty"`
	DocsDeleted  string `json:"docs.deleted,omitempty"`
	StoreSize    string `json:"store.size,omitempty"`
	PriStoreSize string `json:"pri.store.size,omitempty"`
}

type IndexResponse struct {
	Index string
}

type searchResult struct {
	Hits searchHits `json:"hits"`
}

type searchHits struct {
	Total total `json:"total"`
}

type total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type IndexService interface {
	List() (response []IndexResponse, err error)
	Search(index string) error
}
