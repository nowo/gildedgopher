package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	tests := []struct {
		name            string
		sellIn          int
		startQuality    int
		expectedQuality int
	}{
		{"+5 Dexterity Vest", 10, 20, 18},
		{"+5 Dexterity Vest", 0, 1, 18},
		{"Aged Brie", 2, 0, 2},
		{"Elixir of the Mongoose", 5, 7, 5},
		{"Sulfuras, Hand of Ragnaros", 0, 80, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 15, 19},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 10, 16},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 5, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 48, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 40, 42},
		{"Conjured Mana Cake", 3, 6, 2},
		{"Conjured Mana Cake", 0, 6, 0},
	}

	days := 2
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := &gildedrose.Item{Name: test.name, SellIn: test.sellIn, Quality: test.startQuality}

			for day := 0; day < days; day++ {
				gildedrose.UpdateQuality([]*gildedrose.Item{item})
			}

			if item.Quality != test.expectedQuality {
				t.Errorf("Expected quality for %d days: %d, got: %d", days, test.expectedQuality, item.Quality)
			}
		})
	}
}
