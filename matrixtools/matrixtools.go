package matrixtools

import "matrices/arraytools"

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {
	return m[i]
}

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	return Transposed(m)[j]
}

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {
	arraytools.Add(m[i], m[j])
}

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {
	arraytools.ScalarMult(m[i], factor)
}

// Transpose erwartet eine Matrix und liefert ihre Transponierte.
// D.h. alle Zeilen der ersten Matrix werden zu Spalten der Transponierten und umgekehrt.
func Transposed(m [][]float64) [][]float64 {

	// Transponierte Matrix wird initialisiert, die Länge ist die Länge der ersten Reihe
	transposed := make([][]float64, len(m[0]))

	//ranged durch die Reihen
	for r := range transposed {

		//Spaltenanzahl von Transposed initialisieren
		transposed[r] = make([]float64, len(m))

		//ranged durch die Elemente einer Bestimmten Zeile von Transposed
		for s := range transposed[r] {

			//Element in Reihe s und Spalte r in m wird in die Transponierte Matrix geschrieben
			//Reihenwert wird mit Spaltenwert vertauscht
			transposed[r][s] = m[s][r]
		}
	}

	return transposed
}
