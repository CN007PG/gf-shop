package model

type ItemCatCreateInput struct {
	ParentId  int64
	Name      string
	Status    int
	SortOrder int
	IsParent  int
}
