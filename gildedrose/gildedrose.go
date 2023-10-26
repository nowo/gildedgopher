package gildedrose

import (
	"strings"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch {
		case strings.Contains(item.Name, "Sulfuras"):
			continue
		case item.Name == "Aged Brie":
			updateBrie(item)
		case strings.Contains(item.Name, "Backstage passes"):
			updateBackstagePasses(item)
		case strings.Contains(item.Name, "Conjured"):
			updateConjured(item)
		default:
			updateRegularItem(item)
		}
		item.SellIn--
	}
}

func updateBrie(item *Item) {
	if item.Quality < 50 {
		item.Quality++
	}
}

func updateBackstagePasses(item *Item) {
	switch {
	case item.SellIn <= 0:
		item.Quality = 0
	case item.SellIn < 6:
		item.Quality += 3
	case item.SellIn < 11:
		item.Quality += 2
	default:
		item.Quality++
	}

	if item.Quality > 50 {
		item.Quality = 50
	}
}

func updateConjured(item *Item) {
	decrease := 2
	if item.SellIn <= 0 {
		decrease = 4
	}
	item.Quality -= decrease
	if item.Quality < 0 {
		item.Quality = 0
	}
}

func updateRegularItem(item *Item) {
	decrease := 1
	if item.SellIn <= 0 {
		decrease = 2
	}
	item.Quality -= decrease
	if item.Quality < 0 {
		item.Quality = 0
	}
}
