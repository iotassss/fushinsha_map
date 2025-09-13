package domain

import "fmt"

type Emoji string

func NewEmoji(emoji string) (Emoji, error) {
	runes := []rune(emoji)
	if len(runes) != 1 {
		return "", fmt.Errorf("invalid emoji: only a single face emoji (U+1F600–U+1F64A) is allowed")
	}
	r := runes[0]
	if r >= 0x1F600 && r <= 0x1F64A {
		return Emoji(emoji), nil
	}
	return "", fmt.Errorf("invalid emoji: only a single face emoji (U+1F600–U+1F64A) is allowed")
}

func (e Emoji) String() string { return string(e) }
