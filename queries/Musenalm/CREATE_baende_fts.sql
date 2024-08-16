CREATE VIRTUAL TABLE baende_fts USING fts5 (
    id,

    Titelangabe,
    Jahr,
    Verantwortlichkeitsangabe,
    Ortsangabe,
    Ausgabebezeichnung,
    Anmerkungen,

    Bevorzugter_Reihentitel,
    Alternativer_Reihentitel,
    Franzoesischer_Reihentitel,
    Deutscher_Reihentitel,
    Alternatives_Titelblatt,
    hat_TA,
    TA_von,
	
    Herausgabe,
	Verlag,
    Druck,
    Vertrieb,
    Erscheinungsorte,
    Content=Baende,
    tokenize="trigram remove_diacritics 1",
)
