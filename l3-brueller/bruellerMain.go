package main

func main() {
	hund := *new(hund)
	katze := *new(katze)
	maus := *new(maus)
	bruellen(hund)
	bruellen(katze)
	bruellen(maus)
}

func bruellen(b brueller) {
	b.bruellen()
}
