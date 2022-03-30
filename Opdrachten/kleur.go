package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	// Functiecall
	fmt.Println("Kies een kleur")
	fmt.Println("Opties: Blauw, Groen, Rood, Geel, Paars, Roze & Wit")
	var kleurKeuze string
	fmt.Scanln(&kleurKeuze)
	kleur(kleurKeuze)
}

func kleur(kleur string) {
	if kleur == "Blauw" || kleur == "blauw" {
		fmt.Println(kleur, " zoals de lucht")

	} else if kleur == "Groen" || kleur == "groen" {
		fmt.Println(kleur, " van de natuur")

	} else if kleur == "Rood" || kleur == "rood" {
		fmt.Println(kleur, " met passie")

	} else if kleur == "Geel" || kleur == "geel" {
		fmt.Println(kleur, " als de stralen van de zon")

	} else if kleur == "Paars" || kleur == "paars" {
		fmt.Println(kleur, " zoals de bloembedden in het veld")

	} else if kleur == "Roze" || kleur == "roze" {
		fmt.Println(kleur, " zoals de biggetjes bij de boer")

	} else if kleur == "Wit" || kleur == "wit" {
		fmt.Println(kleur, " zoals het sneeuw in de bergen")

	} else {
		fmt.Println("Ongeldige kleur, kies alleen uit het rijtje")
		main()
	}
}
