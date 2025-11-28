package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Datenmodell implementieren",
		117: "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren",
		346: "346 Cloud Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[107] = "Codierungs-, Kompressions- und Verschlüsselungsverfahren einsetzen"
	// TODO: replace one
	modules[107] = "ICT-Lösungen mit Blockchain Technologie umsetzen"
	fmt.Println(modules)
}