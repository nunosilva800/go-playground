/*
We are working on a security system for a badged-access room in our company's building.

We want to find employees who badged into our secured room unusually often. We have an unordered
list of names and entry times over a single day. Access times are given as numbers up to four digits
in length using 24-hour time, such as "800" or "2250".

Write a function that finds anyone who badged into the room three or more times in a one-hour
period. Your function should return each of the employees who fit that criteria, plus the times that
they badged in during the one-hour period. If there are multiple one-hour periods where this was
true for an employee, just return the earliest one for that employee.

badge_times = [
  ["Paul",      "1355"], ["Jennifer",  "1910"], ["Jose",    "835"],
  ["Jose",       "830"], ["Paul",      "1315"], ["Chloe",     "0"],
  ["Chloe",     "1910"], ["Jose",      "1615"], ["Jose",   "1640"],
  ["Paul",      "1405"], ["Jose",       "855"], ["Jose",    "930"],
  ["Jose",       "915"], ["Jose",       "730"], ["Jose",    "940"],
  ["Jennifer",  "1335"], ["Jennifer",   "730"], ["Jose",   "1630"],
  ["Jennifer",     "5"], ["Chloe",     "1909"], ["Zhang",     "1"],
  ["Zhang",       "10"], ["Zhang",      "109"], ["Zhang",   "110"],
  ["Amos",         "1"], ["Amos",         "2"], ["Amos",    "400"],
  ["Amos",       "500"], ["Amos",       "503"], ["Amos",    "504"],
  ["Amos",       "601"], ["Amos",       "602"], ["Paul",   "1416"],
];

Expected output (in any order)
   Paul: 1315 1355 1405
   Jose: 830 835 855 915 930
   Zhang: 10 109 110
   Amos: 500 503 504

n: length of the badge records array

*/

package main

import "fmt"

func main() {
	badge_times := [][]string{
		[]string{"Paul", "1355"},
		[]string{"Jennifer", "1910"},
		[]string{"Jose", "835"},
		[]string{"Jose", "830"},
		[]string{"Paul", "1315"},
		[]string{"Chloe", "0"},
		[]string{"Chloe", "1910"},
		[]string{"Jose", "1615"},
		[]string{"Jose", "1640"},
		[]string{"Paul", "1405"},
		[]string{"Jose", "855"},
		[]string{"Jose", "930"},
		[]string{"Jose", "915"},
		[]string{"Jose", "730"},
		[]string{"Jose", "940"},
		[]string{"Jennifer", "1335"},
		[]string{"Jennifer", "730"},
		[]string{"Jose", "1630"},
		[]string{"Jennifer", "5"},
		[]string{"Chloe", "1909"},
		[]string{"Zhang", "1"},
		[]string{"Zhang", "10"},
		[]string{"Zhang", "109"},
		[]string{"Zhang", "110"},
		[]string{"Amos", "1"},
		[]string{"Amos", "2"},
		[]string{"Amos", "400"},
		[]string{"Amos", "500"},
		[]string{"Amos", "503"},
		[]string{"Amos", "504"},
		[]string{"Amos", "601"},
		[]string{"Amos", "602"},
		[]string{"Paul", "1416"},
	}

	fmt.Println(hours(badge_times))
}

func hours(input [][]string) map[string][]string {
	// map names into entrie times (ordered)
	data := make(map[string][]string)

	// 1. accumulate data into map
	for _, entrie := range input {
		data[entrie[0]] = append(data[entrie[0]], entrie[1])
	}

	// 2. sort all values of map
	// for _, vals := range data {
	// sort...
	// }

	// 3. analyse

	return data
}
