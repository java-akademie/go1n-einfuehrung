package main

// Alle Typen die eine Funktion "bruellen()" haben,
// "implementieren" das Interface "brueller".
// In hund, katze, maus sind keine Vorkehrungen
// zu treffen. Das geschieht automatisch.
type brueller interface{ bruellen() }
