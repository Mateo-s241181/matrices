package matrix

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {
	// TODO
}

// EliminateBelow erwartet eine Spaltennummer `col`.
// Multipliziert alle Zeilen row unter der Zeile col mit -1/Matrix[row][col] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile col ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(col int) {
	// TODO
}

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(row int) {
	// TODO
}

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {
	// TODO
}

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {
	// TODO
}

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {
	// TODO
}
