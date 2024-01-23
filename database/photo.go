package database

type photo struct {
	likes []*user
}

func (p *photo) addLike(user *user) {
	p.likes = append(p.likes, user)
}
