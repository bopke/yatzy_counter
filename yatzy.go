package yatzy

import "errors"

const (
	maxRollValue = 6
	minRollValue = 1
)

type Category byte

const (
	CategoryYatzy Category = iota
	CategoryPair
	CategoryTwoPairs
	CategoryThree
	CategoryFour
	CategorySmallStraight
	CategoryLargeStraight
	CategoryFullHouse
)

func zeroedMap() map[int]int {
	m := make(map[int]int)
	for i := 1; i < 6; i++ {
		m[i] = 0
	}
	return m
}

func rollsFilledMap(rolls []int) map[int]int {
	m := zeroedMap()
	for _, roll := range rolls {
		m[roll]++
	}
	return m
}

var (
	ErrIncorrectRollValue   = errors.New("incorrect roll value")
	ErrIncorrectRollsLength = errors.New("incorrect rolls slice length")
)

func scoreYatzy(rolls []int) int {
	for _, roll := range rolls {
		if roll != rolls[0] {
			return 0
		}
	}
	return 50
}

func scorePair(rolls []int) int {
	rollsmap := rollsFilledMap(rolls)
	highestPairValue := 0
	for rollValue, amount := range rollsmap {
		if amount == 2 {
			if rollValue > highestPairValue {
				highestPairValue = rollValue
			}
		}
	}
	return highestPairValue * 2
}

func scoreTwoPairs(rolls []int) int {
	rollsmap := rollsFilledMap(rolls)
	previousPairValue := 0
	for rollValue, amount := range rollsmap {
		if amount == 2 {
			if previousPairValue == 0 {
				previousPairValue = rollValue
			} else {
				return previousPairValue*2 + rollValue*2
			}
		}
	}
	return 0
}

func scoreThree(rolls []int) int {
	rollsmap := rollsFilledMap(rolls)
	highestThreeValue := 0
	for rollValue, amount := range rollsmap {
		if amount == 3 {
			if rollValue > highestThreeValue {
				highestThreeValue = rollValue
			}
		}
	}
	return highestThreeValue * 3

}

func scoreFour(rolls []int) int {
	rollsmap := rollsFilledMap(rolls)
	highestFourValue := 0
	for rollValue, amount := range rollsmap {
		if amount == 4 {
			if rollValue > highestFourValue {
				highestFourValue = rollValue
			}
		}
	}
	return highestFourValue * 4
}

func scoreSmallStraight(rolls []int) int {
numbersloop:
	for i := 1; i < 6; i++ {
		for _, roll := range rolls {
			if roll == i {
				continue numbersloop
			}
		}
		return 0
	}
	return 15
}

func scoreLargeStraight(rolls []int) int {
numbersloop:
	for i := 2; i < 7; i++ {
		for _, roll := range rolls {
			if roll == i {
				continue numbersloop
			}
		}
		return 0
	}
	return 20
}

func scoreFullHouse(rolls []int) int {
	three := scoreThree(rolls)
	if three == 0 {
		return 0
	}
	pair := scorePair(rolls)
	if pair == 0 {
		return 0
	}
	return three + pair
}

func checkRollsCorrectness(rolls []int, minAcceptableRollValue, maxAcceptableRollValue int) error {
	if len(rolls) != 5 {
		return ErrIncorrectRollsLength
	}
	for _, roll := range rolls {
		if roll < minAcceptableRollValue || roll > maxAcceptableRollValue {
			return ErrIncorrectRollValue
		}
	}
	return nil
}

func answer(rolls []int, category Category) (int, error) {
	err := checkRollsCorrectness(rolls, 1, 6)
	if err != nil {
		return 0, err
	}
	switch category {
	case CategoryYatzy:
		return scoreYatzy(rolls), nil
	case CategoryPair:
		return scorePair(rolls), nil
	case CategoryTwoPairs:
		return scoreTwoPairs(rolls), nil
	case CategoryThree:
		return scoreThree(rolls), nil
	case CategoryFour:
		return scoreFour(rolls), nil
	case CategorySmallStraight:
		return scoreSmallStraight(rolls), nil
	case CategoryLargeStraight:
		return scoreLargeStraight(rolls), nil
	case CategoryFullHouse:
		return scoreFullHouse(rolls), nil
	}
	return -1, nil // should never happen
}
