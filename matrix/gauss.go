package matrix

import "matrices/arraytools"

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {

	if !AlmostEqual(m[col][col], 0) {
		arraytools.ScalarMult(m[col], 1/m[col][col])
	} else {
		panic("STOP")
	}
}

// EliminateBelow erwartet eine Spaltennummer `col`.
// Multipliziert alle Zeilen unter der Zeile row mit -1/Matrix[row][col] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile col ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(col int) {
	//Zeile col normieren
	m.Normalize(col)

	//Zeile Row mit -1/m[row][col] multiplizieren
	for i := range m {

		//Nur Reihen unter m[col] betrachten
		if i > col {

			//Checken ob zahl unter dem Diagonalelement 0 ist
			if !AlmostEqual(m[i][col], 0) {
				arraytools.ScalarMult(m[i], -1/m[i][col])

				//m[row] + m[col]
				arraytools.Add(m[i], m[col])
			}
		}

	}

}

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(col int) {
	//Zeile col normieren
	m.Normalize(col)

	//Zeile Row mit -1/m[row][col] multiplizieren
	for i := range m {

		//Nur Reihen unter m[col] betrachten
		if i < col {

			//reihen mit faktor multiplizieren
			if !AlmostEqual(m[i][col], 0) {
				arraytools.ScalarMult(m[i], -1/m[i][col])

				//m[row] + m[col]
				arraytools.Add(m[i], m[col])
			}
		}

	}
}

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {

	for i := range m {
		m.EliminateBelow(i)
	}

	FixZeros(m)
}

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {

	for i := range m {
		//Muss die Reihenanzah len(m) nutzen und nicht len(m[0]), weil in Eliminierungsmatritzen die Spaltenanzahl höher als die Reihenanzahl ist
		m.EliminateAbove(len(m) - 1 - i)
	}

	FixZeros(m)

}

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {

	m.UpperTriangular()
	m.LowerTriangular()

}

// almostEqual ist eine Hilfsfunktion, die zwei float64-Werte auf Gleichheit prüft.
// Da float64-Werte nicht exakt vergleichbar sind, wird ein Toleranzwert von 1e-10 verwendet.
func AlmostEqual(a, b float64) bool {
	return a-b < 1e-10 && b-a < 1e-10
}

// fixZeros erwartet eine Matrix und setzt alle Elemente, die kleiner als 1e-10 sind, auf 0.
func FixZeros(m Matrix) {
	for i := range m {
		for j := range m[i] {
			if AlmostEqual(m[i][j], 0) {
				m[i][j] = 0
			}
		}
	}
}
