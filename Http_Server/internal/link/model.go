package link

import (
	"gorm.io/gorm"
	"math/rand"
)

type Link struct {
	gorm.Model
	Url  string `json:"url,omitempty"`
	Hash string `json:"hash,omitempty" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = randStringRunes(6)
}

var validRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = validRunes[rand.Intn(len(validRunes))]
	}
	return string(b)
}
