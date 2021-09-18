package chat

type entry struct {
	user    string
	message string
}

type spyPublisher struct {
	published []entry
}

func (p *spyPublisher) Publish(user, message string) error {
	p.published = append(p.published, entry{user: user, message: message})
	return nil
}
