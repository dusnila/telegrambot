package storage

type Storage interface {
	Save(p *Page) error
	PickRandom(UserName string) (*Page, error)
	Remove(p *Page) error
	isExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}
