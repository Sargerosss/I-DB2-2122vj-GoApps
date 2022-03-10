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
		main()

	} else if kleur == "Groen" || kleur == "groen" {
		fmt.Println(kleur, " van de natuur")
		main()

	} else if kleur == "Rood" || kleur == "rood" {
		fmt.Println(kleur, " met passie")
		main()

	} else if kleur == "Geel" || kleur == "geel" {
		fmt.Println(kleur, " als de stralen van de zon")
		main()

	} else if kleur == "Paars" || kleur == "paars" {
		fmt.Println(kleur, " zoals de bloembedden in het veld")
		main()

	} else if kleur == "Roze" || kleur == "roze" {
		fmt.Println(kleur, " zoals de biggetjes bij de boer")
		main()

	} else if kleur == "Wit" || kleur == "wit" {
		fmt.Println(kleur, " zoals het sneeuw in de bergen")
		main()

	} else {
		fmt.Println("Ongeldige kleur, kies alleen uit het rijtje")
		main()
	}
}
