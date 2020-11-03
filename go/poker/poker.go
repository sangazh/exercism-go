package poker

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type PokerHandType int

const (
	highCard PokerHandType = iota
	onePair
	twoPair
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
)

type PokerHand struct {
	hand  string
	cards []*PokerCard
	kind  PokerHandType
	nums  map[int]int   // k:num v:count
	suits map[rune]int  // k:suit v:count
	score map[int][]int // k:count v:num  = nums reverse
}

type PokerCard struct {
	card   string
	suit   rune
	number int
}

func NewCard(card string) (pc *PokerCard, err error) {
	r, size := utf8.DecodeLastRuneInString(card)
	suit := string(r)
	number := card[:len(card)-size]
	var n int

	switch number {
	case "A":
		n = 14
	case "J":
		n = 11
	case "Q":
		n = 12
	case "K":
		n = 13
	default:
		n, err = strconv.Atoi(number)
		if err != nil {
			return nil, errors.New("invalid card " + err.Error())
		}
		if n > 10 || n < 2 {
			return nil, errors.New("invalid card 1")
		}
	}

	switch suit {
	case "♢", "♡", "♧", "♤":
	default:
		return nil, errors.New("invalid card suit")
	}

	return &PokerCard{
		card:   card,
		suit:   r,
		number: n,
	}, nil
}

func NewPokerHand(hand string) (*PokerHand, error) {
	cards, err := check(strings.Split(hand, " "))
	if err != nil {
		return nil, err
	}

	p := &PokerHand{
		hand:  hand,
		cards: cards,
		suits: make(map[rune]int),
		nums:  make(map[int]int),
		score: make(map[int][]int),
	}
	p.parse()
	return p, nil
}

func check(cards []string) ([]*PokerCard, error) {
	if len(cards) != 5 {
		return nil, errors.New("cards number wrong")
	}
	pcs := make([]*PokerCard, 0)
	for _, card := range cards {
		pc, err := NewCard(card)
		if err != nil {
			return nil, err
		}

		pcs = append(pcs, pc)
	}
	return pcs, nil
}

func (p *PokerHand) parse() {
	for _, card := range p.cards {
		if _, ok := p.suits[card.suit]; ok {
			p.suits[card.suit] += 1
		} else {
			p.suits[card.suit] = 1
		}
		if _, ok := p.nums[card.number]; ok {
			p.nums[card.number] += 1
		} else {
			p.nums[card.number] = 1
		}
	}

	for n, count := range p.nums {
		if _, ok := p.score[count]; ok {
			p.score[count] = append(p.score[count], n)
		} else {
			p.score[count] = []int{n}
		}
	}

	for _, s := range p.score {
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
	}

	p.kind = p.Kind()

}

func (p *PokerHand) Kind() PokerHandType {
	sort.Slice(p.cards, func(i, j int) bool {
		return p.cards[i].number < p.cards[j].number
	})

	//only one suit, could be flush or straightFlush
	if len(p.suits) == 1 {
		if len(p.nums) == 5 {
			// max - min = 5 -> straightFlush
			if p.cards[4].number-p.cards[0].number == 4 {
				return straightFlush
			}
			return flush
		}
	}

	//two numbers, 1 4 or 2 3
	if len(p.nums) == 2 {
		for _, count := range p.nums {
			switch count {
			case 1, 4:
				return fourOfAKind
			case 2, 3:
				return fullHouse
			}
		}
	}

	//three numbers, 1 1 3 or 1 2 2
	if len(p.nums) == 3 {
		for _, count := range p.nums {
			switch count {
			case 3:
				return threeOfAKind
			case 2:
				return twoPair
			}
		}
	}

	if len(p.nums) == 4 {
		return onePair
	}

	if len(p.nums) == 5 && p.cards[4].number-p.cards[0].number == 4 {
		return straight
	}

	// aces can start a straight (A 2 3 4 5)
	if len(p.nums) == 5 && p.cards[3].number-p.cards[0].number == 3 && p.cards[4].number == 14 {
		p.cards[4].number = 1
		p.cards = append([]*PokerCard{p.cards[4]}, p.cards[:4]...)
		return straight
	}

	return highCard
}

func (p PokerHand) Equal(p2 PokerHand) bool {
	if p.kind != p2.kind {
		return false
	}

	for n, count := range p.nums {
		p2count, ok := p2.nums[n]
		if !ok {
			return false
		}

		if p2count != count {
			return false
		}
	}

	return true
}
func BestHand(hands []string) (actual []string, err error) {
	pokerHands := make([]*PokerHand, 0)
	for _, h := range hands {
		ph, err := NewPokerHand(h)
		if err != nil {
			return nil, err
		}
		pokerHands = append(pokerHands, ph)
	}

	if len(pokerHands) <= 1 {
		return hands, nil
	}

	sort.Slice(pokerHands, func(i, j int) bool {
		a, b := pokerHands[i], pokerHands[j]
		if a.kind > b.kind {
			return true
		} else if a.kind < b.kind {
			return false
		}

		switch a.kind {
		case fourOfAKind:
			if a.score[4][0] > b.score[4][0] {
				return true
			} else if a.score[4][0] < b.score[4][0] {
				return false
			} else {
				return a.score[1][0] > b.score[1][0]
			}
		case fullHouse:
			if a.score[3][0] > b.score[3][0] {
				return true
			} else if a.score[3][0] < b.score[3][0] {
				return false
			} else {
				return a.score[2][0] > b.score[2][0]
			}
		case threeOfAKind:
			if a.score[3][0] > b.score[3][0] {
				return true
			} else if a.score[3][0] < b.score[3][0] {
				return false
			}
			for i := 0; i < 2; i++ {
				if a.score[1][i] == b.score[1][i] {
					continue
				}
				return a.score[1][i] > b.score[1][i]
			}
		}

		for i := 4; i >= 0; i-- {
			if a.cards[i].number == b.cards[i].number {
				continue
			}
			return a.cards[i].number > b.cards[i].number
		}
		return false
	})

	if pokerHands[0].Equal(*pokerHands[1]) {
		return []string{pokerHands[0].hand, pokerHands[1].hand}, nil
	}

	return []string{pokerHands[0].hand}, nil
}
