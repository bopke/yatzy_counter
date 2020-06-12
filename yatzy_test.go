package yatzy

import (
	"testing"
)

func TestRollsFilledMap(t *testing.T) {
	nums := []int{1, 1, 1, 2, 4}
	m := rollsFilledMap(nums)
	if m[1] != 3 {
		t.Error("rollsFilledMap function for 1 from", nums, "should be 3 but is", m[1])
	}
	if m[2] != 1 {
		t.Error("rollsFilledMap function for 2 from", nums, "should be 1 but is", m[2])
	}
	if m[3] != 0 {
		t.Error("rollsFilledMap function for 3 from", nums, "should be 0 but is", m[3])
	}
}

func TestCheckNumberCorrectness(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3}
	ret := checkRollsCorrectness(nums, 0, 5)
	if ret != nil {
		t.Error("checkRollsCorrectness function for number from the middle of the range should return nil but returned ", ret)
	}
	ret = checkRollsCorrectness(nums, 1, 5)
	if ret != nil {
		t.Error("checkRollsCorrectness function for number from the lower end of the range should return \"nil\" but returned \"", ret, "\"")
	}
	ret = checkRollsCorrectness(nums, 1, 5)
	if ret != nil {
		t.Error("checkRollsCorrectness function for number from the higher end of the range should return \"nil\" but returned \"", ret, "\"")
	}
	ret = checkRollsCorrectness(nums, 1, 3)
	if ret != ErrIncorrectRollValue {
		t.Error("checkRollsCorrectness function for number above the range should return \"", ErrIncorrectRollValue, "\" but returned \"", ret, "\"")
	}
	ret = checkRollsCorrectness(nums, 2, 5)
	if ret != ErrIncorrectRollValue {
		t.Error("checkRollsCorrectness function for number below the range should return \"", ErrIncorrectRollValue, "\" but returned \"", ret, "\"")
	}
	nums = []int{1, 2}
	ret = checkRollsCorrectness(nums, 1, 6)
	if ret != ErrIncorrectRollsLength {
		t.Error("checkRollsCorrectness function for rolls slice with incorrect length should return \"", ErrIncorrectRollsLength, "\" but returned \"", ret, "\"")
	}
}

func TestScoreYatzy(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	ret := scoreYatzy(nums)
	if ret != 50 {
		t.Error("scoreYatzy function for yatzy should return 50 but returned ", ret)
	}
	nums = []int{1, 1, 1, 2, 1}
	ret = scoreYatzy(nums)
	if ret != 0 {
		t.Error("scoreYatzy function for definitely not yatzy should return 0 but returned ", ret)
	}
}

func TestScorePair(t *testing.T) {
	nums := []int{1, 1, 2, 3, 4}
	ret := scorePair(nums)
	if ret != 2 {
		t.Error("scorePair function for ordered pair of 1 should return 4 but returned ", ret)
	}
	nums = []int{1, 2, 1, 4, 3}
	ret = scorePair(nums)
	if ret != 2 {
		t.Error("scorePair function for unordered pair of 1 should return 4 but returned ", ret)
	}
	nums = []int{1, 5, 2, 3, 4}
	ret = scorePair(nums)
	if ret != 0 {
		t.Error("scorePair function for definitly not pair of any kind should return 0 but returned ", ret)
	}
	nums = []int{3, 3, 1, 1, 2}
	ret = scorePair(nums)
	if ret != 6 {
		t.Error("scorePair function for two pairs of 1 and 3 should return 6 but returned ", ret)
	}
	nums = []int{3, 3, 3, 1, 2}
	ret = scorePair(nums)
	if ret != 0 {
		t.Error("scorePair function for three of 3 should return 0 but returned ", ret)
	}
}

func TestScoreTwoPairs(t *testing.T) {
	nums := []int{1, 1, 2, 3, 4}
	ret := scoreTwoPairs(nums)
	if ret != 0 {
		t.Error("scoreTwoPairs function for ordered pair of 1 should return 0 but returned ", ret)
	}
	nums = []int{1, 2, 1, 4, 3}
	ret = scoreTwoPairs(nums)
	if ret != 0 {
		t.Error("scoreTwoPairs function for unordered pair of 1 should return 0 but returned ", ret)
	}
	nums = []int{1, 5, 2, 3, 4}
	ret = scoreTwoPairs(nums)
	if ret != 0 {
		t.Error("scoreTwoPairs function for definitly not pair of any kind should return 0 but returned ", ret)
	}
	nums = []int{3, 3, 1, 1, 2}
	ret = scoreTwoPairs(nums)
	if ret != 8 {
		t.Error("scoreTwoPairs function for two pairs of 1 and 3 should return 8 but returned ", ret)
	}
	nums = []int{3, 3, 1, 1, 1}
	ret = scoreTwoPairs(nums)
	if ret != 0 {
		t.Error("scoreTwoPairs function for full house of 1 and 3 should return 0 but returned ", ret)
	}
}

func TestScoreThree(t *testing.T) {
	nums := []int{1, 1, 1, 2, 4}
	ret := scoreThree(nums)
	if ret != 3 {
		t.Error("scoreThree function for ordered three of 1 should return 3 but returned ", ret)
	}
	nums = []int{1, 2, 1, 4, 1}
	ret = scoreThree(nums)
	if ret != 3 {
		t.Error("scoreThree function for unordered three of 1 should return 3 but returned ", ret)
	}
	nums = []int{1, 1, 2, 3, 4}
	ret = scoreThree(nums)
	if ret != 0 {
		t.Error("scoreThree function for definitly not three of any kind should return 0 but returned ", ret)
	}
	nums = []int{1, 1, 1, 1, 4}
	ret = scoreThree(nums)
	if ret != 0 {
		t.Error("scoreThree function for four of 1 should return 0 but returned ", ret)
	}
}

func TestScoreFour(t *testing.T) {
	nums := []int{1, 1, 1, 1, 4}
	ret := scoreFour(nums)
	if ret != 4 {
		t.Error("scoreFour function for ordered four of 1 should return 4 but returned ", ret)
	}
	nums = []int{1, 1, 1, 4, 1}
	ret = scoreFour(nums)
	if ret != 4 {
		t.Error("scoreFour function for unordered four of 1 should return 4 but returned ", ret)
	}
	nums = []int{1, 1, 2, 3, 4}
	ret = scoreFour(nums)
	if ret != 0 {
		t.Error("scoreFour function for definitly not four of any kind should return 0 but returned ", ret)
	}
	nums = []int{1, 1, 1, 1, 1}
	ret = scoreFour(nums)
	if ret != 0 {
		t.Error("scoreFour function for yatzy of 1 should return 0 but returned ", ret)
	}
}

func TestScoreSmallStraight(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ret := scoreSmallStraight(nums)
	if ret != 15 {
		t.Error("scoreSmallStraight function for ordered small straight should return 15 but returned ", ret)
	}
	nums = []int{4, 5, 3, 1, 2}
	ret = scoreSmallStraight(nums)
	if ret != 15 {
		t.Error("scoreSmallStraight function for unordered small straight should return 15 but returned ", ret)
	}
	nums = []int{1, 2, 3, 4, 2}
	ret = scoreSmallStraight(nums)
	if ret != 0 {
		t.Error("scoreSmallStraight function for definitely not small straight should return 0 but returned ", ret)
	}
}

func TestScoreLargeStraight(t *testing.T) {
	nums := []int{2, 3, 4, 5, 6}
	ret := scoreLargeStraight(nums)
	if ret != 20 {
		t.Error("scoreLargeStraight function for ordered large straight should return 15 but returned ", ret)
	}
	nums = []int{4, 5, 3, 6, 2}
	ret = scoreLargeStraight(nums)
	if ret != 20 {
		t.Error("scoreLargeStraight function for unordered large straight should return 15 but returned ", ret)
	}
	nums = []int{3, 2, 3, 4, 6}
	ret = scoreLargeStraight(nums)
	if ret != 0 {
		t.Error("scoreLargeStraight function for definitely not large straight should return 0 but returned ", ret)
	}
}

func TestScoreFullHouse(t *testing.T) {
	nums := []int{2, 2, 2, 3, 3}
	ret := scoreFullHouse(nums)
	if ret != 12 {
		t.Error("scoreFullHouse function for full house (2,2,2,3,3) should return 12 but returned ", ret)
	}
	nums = []int{2, 1, 3, 4, 5}
	ret = scoreFullHouse(nums)
	if ret != 0 {
		t.Error("scoreFullHouse function for definitely not full house should return 0 but returned ", ret)
	}
	nums = []int{2, 2, 3, 3, 5}
	ret = scoreFullHouse(nums)
	if ret != 0 {
		t.Error("scoreFullHouse function for definitely not full house should return 0 but returned ", ret)
	}
	nums = []int{2, 1, 3, 3, 3}
	ret = scoreFullHouse(nums)
	if ret != 0 {
		t.Error("scoreFullHouse function for definitely not full house should return 0 but returned ", ret)
	}
	nums = []int{3, 3, 3, 3, 3}
	ret = scoreFullHouse(nums)
	if ret != 0 {
		t.Error("scoreFullHouse function for yatzy should return 0 but returned ", ret)
	}
}

func TestAnswer(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1}
	_, err := answer(nums, CategoryYatzy)
	if err != ErrIncorrectRollsLength {
		t.Error("answer function for rolls slice with incorrect length should return \"", ErrIncorrectRollsLength, "\" but returned \"", err, "\"")
	}
	nums = []int{1, 1, 1, 10, 1}
	_, err = answer(nums, CategoryYatzy)
	if err != ErrIncorrectRollValue {
		t.Error("answer function for rolls slice with incorrect content should return \"", ErrIncorrectRollValue, "\" but returned \"", err, "\"")
	}
	nums = []int{1, 1, 1, 1, 1}
	ret, err := answer(nums, CategoryYatzy)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 50 {
		t.Error("answer function for correct yatzy should return 50 but returned ", ret)
	}
	nums = []int{1, 1, 5, 4, 3}
	ret, err = answer(nums, CategoryPair)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 2 {
		t.Error("answer function for correct pair of 1 should return 2 but returned ", ret)
	}
	nums = []int{1, 1, 2, 3, 2}
	ret, err = answer(nums, CategoryTwoPairs)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 6 {
		t.Error("answer function for correct two pairs of 1 and 2 should return 6 but returned ", ret)
	}
	nums = []int{1, 1, 1, 2, 3}
	ret, err = answer(nums, CategoryThree)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 3 {
		t.Error("answer function for correct three should return 3 but returned ", ret)
	}
	nums = []int{1, 1, 1, 1, 2}
	ret, err = answer(nums, CategoryFour)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 4 {
		t.Error("answer function for correct four should return 4 but returned ", ret)
	}
	nums = []int{1, 2, 3, 4, 5}
	ret, err = answer(nums, CategorySmallStraight)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 15 {
		t.Error("answer function for correct small straight should return 15 but returned ", ret)
	}
	nums = []int{2, 3, 4, 5, 6}
	ret, err = answer(nums, CategoryLargeStraight)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 20 {
		t.Error("answer function for correct large straight should return 20 but returned ", ret)
	}
	nums = []int{1, 1, 1, 2, 2}
	ret, err = answer(nums, CategoryFullHouse)
	if err != nil {
		t.Error("answer function for rolls slice with correct content should return \"nil\" but returned \"", err, "\"")
	}
	if ret != 7 {
		t.Error("answer function for correct full house should return 7 but returned ", ret)
	}
}
