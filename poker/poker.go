package poker

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	HighHand = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

var faces = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

var suits = map[string]int{
	"♡": 1,
	"♧": 2,
	"♢": 3,
	"♤": 4,
}

type card struct {
	suit int
	face int
}

type hand struct {
	cards []card
	score int
}

func BestHand(hands []string) ([]string, error) {
	// Single hands always win.
	if len(hands) == 1 {
		return hands, nil
	}

	var bestRawHand []string

	for _, h := range hands {
		// TODO fix holder name
		holder := hand{ }

		hs := strings.Split(h, " ")
		if len(hs) != 5 {
			return nil, fmt.Errorf("error: %d is not a valid amount of cards", len(hs))
		}

		for _, pair := range hs {
			var (
				value string
				suit  string
			)

			// TODO is regex better here?
			switch utf8.RuneCountInString(pair) {
			// The 10 has a Rune count of 3...
			case 3:
				value, suit = string(pair[0:2]), pair[2:]
			case 2:
				value, suit = string(pair[0]), pair[1:]
			default:
				return nil, errors.New("No a valid hand")
			}

			c := card{}

			v, ok := faces[value]
			if !ok {
				return nil, fmt.Errorf("%c is not a valid card", v)
			}
			c.face = v

			v, ok = suits[suit]
			if !ok {
				return nil, fmt.Errorf("%c is not a valid card", v)
			}
			c.suit = v

			holder.cards = append(holder.cards, c)
		}

	}

	return []string{}, nil
}

func hasRoyalFlush(hand []card) (ok bool) {
	if !hasFlush(hand) {
		return false
	}

	var totalScore int
	// Want to make sure that all the cards are unique.
	// TODO Do I need this?
	var set map[int]bool

	for _, c := range hand {
		_, ok := set[c.face]
		if !ok {
			return false
		}
		set[c.face] = true

		totalScore += c.face
	}

	return totalScore == 60
}

func hasFlush(hand []card) (ok bool) {
	if len(hand) == 0 {
		return false
	}

	first, rest := hand[0], hand[1:]

	for _, c := range rest {
		if c.suit != first.suit {
			return false
		}

	}

	return true
}

func hasStraight(hand []card) (ok bool) {
	if len(hand) == 0 {
		return false
	}

	sort.SliceStable(hand, func(i, j int) bool {
		return hand[i].face < hand[j].face
	})
	first, rest := hand[0], hand[1:]

	for i, c := range rest {
		if first.face + 1 + i != c.face {
			return false
		}
	}

	return true
}

func totalScore(hand []card) int {
	var result int
	for _, c := range hand {
		result += c.face * c.face
	}

	return result
}

func determineHandRank(hand []card) int {

}

