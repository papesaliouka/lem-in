package helper

import (
	"math"
)

func Elector(eligeables [][]string, thePromised []string, found int, min int, elected [][]string, end string, groups Group) [][]string {
	if len(eligeables) > 0 {

		allPossibilities := GenerateAllPossibilities(thePromised, eligeables)
		for _, v := range allPossibilities {
			if !HasCommonElements2(v) {
				found++
				flatened, flat := Flat2DArray(v)
				if len(flat) < min {
					min = len(flat)
					elected = flatened
				}
			}
		}

		if found == 0 {
			min := math.MaxInt32
			choosen := [][]string{}
			for key, subPath := range groups {
				if key != thePromised[0] {
					_, flat := Flat2DArray(subPath)
					if len(flat) > 0 && len(flat) < min {
						choosen = subPath
						if len(subPath) > 1 {
							for _, val := range subPath {
								if !HasCommonElements(thePromised, val) {
									choosen = [][]string{val}
								}
							}
						}
					}
				}
			}
			elected = append(elected, choosen...)

		}
	}
	return elected
}