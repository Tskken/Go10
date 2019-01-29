package player

import (
	"Go10/deck"
	"log"
	"sort"
	"strconv"
)

type Hand []deck.Card

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) Less(i, j int) bool {
	iInt, err := strconv.Atoi(h[i].Name)
	if err != nil {
		if h[i].Name == deck.Wild || h[i].Name == deck.Skip {
			iInt = 13
		}
	}

	jInt, err := strconv.Atoi(h[j].Name)
	if err != nil {
		if h[j].Name == deck.Wild || h[j].Name == deck.Skip {
			jInt = 13
		}
	}

	return iInt < jInt
}

func (h Hand) SetCount(setCount int, usWild ...string) (bool, Hand) {
	cards := make(map[string]int)

	for _, h := range h {
		cards[h.Name]++
	}

	if wCount, ok := cards[deck.Wild]; ok {
		if usWild != nil && len(usWild) <= wCount{
			for _, u := range usWild {
				cards[u]++
				wCount--
			}
		}
		if wCount > 0 {
			for k, count:= range cards {
				if k != deck.Skip && k != deck.Wild {
					if count + wCount >= 3 {
						wUsed := wCount - ((count + wCount) - 3)
						if wUsed > 0 {
							cards[k] += wUsed
							wCount -= wUsed
						}
					}

					if wCount == 0 {
						break
					}
				}
			}
		}
	}

	sets := 0

	setCards := make(map[string]int)

	for k, count := range cards {
		if k != deck.Skip {
			if count >= 3 {
				setCards[k] = 3
				sets++
			}
		}
	}

	newHand := make(Hand, 0)
	if sets >= setCount {
		for _, c := range h {
			if _, ok := setCards[c.Name]; !ok {
				newHand = append(newHand, c)
			} else {
				setCards[c.Name]--
			}

			if setCards[c.Name] <= 0 {
				delete(setCards, c.Name)
			}
		}
		return true, newHand
	}

	return false, nil
}

/*
	TODO: hand.RunCount()
		- Finished run check
 */

func (h Hand) RunCount(runCount, runLength int, usWild ...string) (bool, Hand) {

	sort.Sort(h)

	current := -1
	count := 0

	for _, c := range h {
		iCard, err := strconv.Atoi(c.Name)
		if err == nil {
			if iCard == current + 1 {
				current++
				count++
			} else if iCard != current {
					current = iCard
					count = 0
			}
		}

		if count >= runCount {
			log.Println("count matched!!!")
			return true, h
		}
	}

	return false, nil
}
