package main

func graph16() map[string][]string {
	graph16 := make(map[string][]string)
	graph16["A"] = []string{"B", "C", "D"}
	graph16["B"] = []string{"C", "E"}
	graph16["C"] = []string{"A", "F"}
	graph16["D"] = []string{"G"}
	graph16["E"] = []string{"H"}
	graph16["F"] = []string{"I"}
	graph16["G"] = []string{"J"}
	graph16["H"] = []string{"K"}
	graph16["I"] = []string{"L"}
	graph16["J"] = []string{"M"}
	graph16["K"] = []string{"N"}
	graph16["L"] = []string{"O"}
	graph16["M"] = []string{"P"}
	graph16["N"] = []string{"Q"}
	graph16["O"] = []string{"R"}
	graph16["P"] = []string{"S"}
	graph16["Q"] = []string{"T"}

	return graph16
}

func graph16NoCycle() map[string][]string {
	graph16NoCycle := make(map[string][]string)
	graph16NoCycle["A"] = []string{"B", "C", "D"}
	graph16NoCycle["B"] = []string{"C", "E"}
	graph16NoCycle["C"] = []string{"F"}
	graph16NoCycle["D"] = []string{"G"}
	graph16NoCycle["E"] = []string{"H"}
	graph16NoCycle["F"] = []string{"I"}
	graph16NoCycle["G"] = []string{"J"}
	graph16NoCycle["H"] = []string{"K"}
	graph16NoCycle["I"] = []string{"L"}
	graph16NoCycle["J"] = []string{"M"}
	graph16NoCycle["K"] = []string{"N"}
	graph16NoCycle["L"] = []string{"O"}
	graph16NoCycle["M"] = []string{"P"}
	graph16NoCycle["N"] = []string{"Q"}
	graph16NoCycle["O"] = []string{"R"}
	graph16NoCycle["P"] = []string{"S"}
	graph16NoCycle["Q"] = []string{"T"}

	return graph16NoCycle
}
