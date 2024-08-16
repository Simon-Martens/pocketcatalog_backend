CREATE VIRTUAL TABLE inhalte_fts USING fts5 (
	id,

	Seite,
	Incipit,
	Titelangabe,
	Urheberangabe,
	Anmerkungen,
	Typ,
	Band,
	Geschaffen,
	Geschrieben,
	Gestochen,
	Gezeichnet,
	content=Inhalte,
	tokenize="trigram remove_diacritics 1"
)
