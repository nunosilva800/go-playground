/*

Given records of people either entering or leaving throught a day, output the name of those that had
an invalid entrie or invalid exit.

An invalid entry is one that has no preceeding exit.
An invalid exit is one that has no preceeding entry.
When the day starts everyone is outside.
When the day ends everyone is outside.

inputA: [
	["Paul", "enter"],
	["Paul", "exit"],
];
outputA: [], []

inputB: [
	["Paul", "exit"],
	["Paul", "enter"],
];
outputB: ["Paul"], ["Paul"]

inputC: [
	["Paul", "enter"],
	["Anna", "enter"],
	["Paul", "enter"],
	["Anna", "exit"],
	["Paul", "exit"],
];
outputC: ["Paul"], []
*/

package main

func main() {
}

func check(input [][]string) ([]string, []string) {
	log := make(map[string]string)
	invalidEntries := make(map[string]struct{})
	invalidExits := make(map[string]struct{})

	for idx, entrie := range input {
		_ = idx
		//     fmt.Println("at", idx, log)

		name := entrie[0]
		event := entrie[1]

		curr, ok := log[name]
		if !ok {
			// exit without entering - invalid
			if event == "exit" {
				invalidExits[name] = struct{}{}
				//         fmt.Println("exit without entering", name, curr, event)
			}
			log[name] = event

			continue
		}

		if event == "enter" {
			if curr != "exit" {
				invalidEntries[name] = struct{}{}
				log[name] = event
				//         fmt.Println("enter without exit", name, curr, event)
				continue
			}
			log[name] = event
		}

		if event == "exit" {
			if curr != "enter" {
				invalidExits[name] = struct{}{}
				log[name] = event
				//         fmt.Println("exit without entering", name, curr, event)
				continue
			}
			log[name] = event
		}
	}

	for name, curr := range log {
		if curr == "enter" {
			invalidEntries[name] = struct{}{}
		}
	}

	invalidEntriesNames := []string{}
	for name, _ := range invalidEntries {
		invalidEntriesNames = append(invalidEntriesNames, name)
	}

	invalidExistsNames := []string{}
	for name, _ := range invalidExits {
		invalidExistsNames = append(invalidExistsNames, name)
	}

	return invalidEntriesNames, invalidExistsNames
}
