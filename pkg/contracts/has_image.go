package contracts

type HasImage interface {
	GetImagePath() string
	SetImagePath(path string)
}
