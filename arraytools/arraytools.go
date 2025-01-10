package arraytools

// Mult erwartet ein Array und einen skalaren Faktor.
// Multipliziert jedes Element des Arrays mit dem Faktor.
func ScalarMult(a []float64, factor float64) {

	//Durch alle Elemente aus a durchrangen
	for i := range a {

		//Ausgewähltes Element mit faktor multiplizieren
		//el ist nur eine Kopie vom
		a[i] *= factor
	}

}

// Add erwartet zwei Arrays der gleichen Länge.
// Addiert die Elemente der Arrays paarweise.
// Ergebnis wird im Array a gespeichert
func Add(a, b []float64) {

	if len(a) != len(b) {
		panic("Die beiden Slices haben nicht die gleiche Länge!")
	}

	for i := range a {
		a[i] += b[i]
	}
}
