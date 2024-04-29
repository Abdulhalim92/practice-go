package same_fct2

import "golang.org/x/exp/constraints"

// проверка нахождения элемента в срезе

func containsUint8(needle uint8, haystack []uint8) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func containsInt(needle int, haystack []int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func contains[E constraints.Ordered](needle E, haystack []E) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
