package main

type iPublication interface {
	getName() string
	setName(name string)
	getPages() int
	setPages(pages int)
	getPublisher() string
	setPublisher(publisher string)
}

type publication struct {
	name      string
	pages     int
	publisher string
}

func (p *publication) getName() string {
	return p.name
}
func (p *publication) setName(name string) {
	p.name = name
}

func (p *publication) getPages() int {
	return p.pages
}
func (p *publication) setPages(pages int) {
	p.pages = pages
}

func (p *publication) getPublisher() string {
	return p.publisher
}
func (p *publication) setPublisher(publisher string) {
	p.publisher = publisher
}
