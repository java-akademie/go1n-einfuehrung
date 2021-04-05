module github.com/java-akademie/go1n-einfuehrung/a-first

go 1.16

// -------------------------------------------------------------------------
//
// wenn an myutils gearbeitet werden soll, muss require gelöscht und
// folgende Befehle abgesetzt werden
//
// $ go mod edit -replace=github.com/java-akademie/myutils=../../myutils
// $ go mod tidy
// ----
// wenn die Arbeit an mytools abgeschlossen wurde
// und commit/push durchgeführt ist,
// muessen die beiden Zeilen geaendert werden
//
// replace auskommentieren
// require loeschen
// ----
// anschliessend ist der folgende Befehl durchzufuehren
// $ go mod tidy
// ----
// wenn require geloescht und replace wieder aktiviert wird, und
// $ go mod tidy
// durchgefuehrt wurde, kann an myutils wieder gearbeitet werden
//
// -------------------------------------------------------------------------

replace github.com/java-akademie/myutils => ../../myutils

require github.com/java-akademie/myutils v0.0.0-00010101000000-000000000000
