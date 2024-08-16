-- THIS IS HOW WE CREATE AN FTS5 TABLE + A TABLE FOR BIBLIOGRAPHIC RECORDS FOR BAENDE
-- TODO: Anmerkungen ist HTML: Tags müssen gelöscht werden // genauso bei: Inhalte
INSERT INTO baende_fts (
	rowid,
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
	Erscheinungsorte
)

SELECT 
    Baende.rowid, 
		Baende.id,
	Titelangabe as TA_B,
	Jahr as J_B,
	Verantwortlichkeitsangabe as VA_B,
	Ortsangabe as OA_B,
	Ausgabebezeichnung as ABZ_B,
	Anmerkungen as A_B,

	BRT as BRT_B,
	ART as ART_B,
    FRT as FRT_B,
    DRT as DRT_B,
    ATBRT as ATBRT_B,
	HTART as HTART_B,
    TAVRT as TAVRT_B,

    Herausgeber, 
    Verleger, 
    Drucker, 
	Vertreiber, 
    Orte
   
    
 FROM Baende

-- Herausgeber
FULL JOIN
	(SELECT
		Baende.id as id,
		group_concat(hrsg.Name, '; ') as Herausgeber
	FROM 
		Baende, 
		json_each(Baende.Herausgabe, '$') as herausgabe
	LEFT JOIN 
		Akteure hrsg ON hrsg.id = herausgabe.value
	GROUP BY
		Baende.id) hrsgtab 
ON Baende.id = hrsgtab.id

-- Verleger
FULL JOIN 
	(SELECT
		Baende.id as id,
		group_concat(Akteure.Name, '; ') as Verleger
	FROM 
		Baende, 
		json_each(Baende.Verlag, '$')
	LEFT JOIN 
		Akteure ON Akteure.id = value
	GROUP BY
		Baende.id) verltab 
ON Baende.id = verltab.id

-- Vertrieb
FULL JOIN 
	(SELECT
		Baende.id as id,
		group_concat(Akteure.Name, '; ') as Vertreiber
	FROM 
		Baende, 
		json_each(Baende.Vertrieb, '$')
	LEFT JOIN 
		Akteure ON Akteure.id = value
	GROUP BY
		Baende.id) verttab 
ON Baende.id = verttab.id

-- Druck
FULL JOIN
	(SELECT
		Baende.id as id,
		group_concat(Akteure.Name, '; ') as Drucker
	FROM 
		Baende, 
		json_each(Baende.Druck, '$')
	LEFT JOIN 
		Akteure ON Akteure.id = value
	GROUP BY
		Baende.id) drucktab 
ON Baende.id = drucktab.id

-- Orte
FULL JOIN 
	(SELECT
		Baende.id as id,
		group_concat(Orte.Name, '; ') as Orte
	FROM 
		Baende, 
		json_each(Baende.Erscheinungsorte, '$')
	LEFT JOIN 
		Orte ON Orte.id = value
	GROUP BY
		Baende.id) orttab 
ON Baende.id = orttab.id

-- Bev. Reihentitel
FULL JOIN 
	(SELECT
		Baende.id as id,
		Reihentitel.Titel as BRT
	FROM 
		Baende
	LEFT JOIN
		Reihentitel ON Reihentitel.id = Baende.Bevorzugter_Reihentitel
	GROUP BY
		Baende.id) rttab 
ON Baende.id = rttab.id

-- Frz. Reihentitel
FULL JOIN 
	(SELECT
		Baende.id as id,
    Reihentitel.Titel as FRT
	FROM 
		Baende
	LEFT JOIN
		Reihentitel ON Reihentitel.id = Baende.Franzoesischer_Reihentitel
	GROUP BY
		Baende.id) frttab 
ON Baende.id = frttab.id

-- Dt. Reihentitel
FULL JOIN 
	(SELECT
		Baende.id as id,
    Reihentitel.Titel as DRT
	FROM 
		Baende
	LEFT JOIN
		Reihentitel ON Reihentitel.id = Baende.Deutscher_Reihentitel
	GROUP BY
		Baende.id) drttab 
ON Baende.id = drttab.id

-- Alt. Reihentitel
FULL JOIN 
	(SELECT
		Baende.id as id,
		group_concat(Reihentitel.Titel, '; ') as ART
	FROM 
		Baende, 
		json_each(Baende.Alternativer_Reihentitel, '$')
	LEFT JOIN 
		Reihentitel ON Reihentitel.id = value
	GROUP BY
		Baende.id) altrttab 
ON Baende.id = altrttab.id

-- Alt. Titelblatt
FULL JOIN 
	(SELECT
		Baende.id as id,
		group_concat(Reihentitel.Titel, '; ') as ATBRT
	FROM 
		Baende, 
		json_each(Baende.Alternatives_Titelblatt, '$')
	LEFT JOIN 
		Reihentitel ON Reihentitel.id = value
	GROUP BY
		Baende.id) alttbrttab 
ON Baende.id = alttbrttab.id

-- TA von
FULL JOIN 
	(SELECT
		Baende.id as id,
		group_concat(Reihentitel.Titel, '; ') as TAVRT
	FROM 
		Baende, 
		json_each(Baende.TA_von, '$')
	LEFT JOIN 
		Reihentitel ON Reihentitel.id = value
	GROUP BY
		Baende.id) tavrttab 
ON Baende.id = tavrttab.id

-- hat TA
FULL JOIN
	(SELECT
		Baende.id as id,
		group_concat(Reihentitel.Titel, '; ') as HTART
	FROM 
		Baende, 
		json_each(Baende.hat_TA, '$')
	LEFT JOIN 
		Reihentitel ON Reihentitel.id = value
	GROUP BY
		Baende.id) htarttab 
ON Baende.id = htarttab.id

GROUP BY
	Baende.id


