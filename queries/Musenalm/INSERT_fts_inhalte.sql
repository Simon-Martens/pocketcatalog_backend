INSERT INTO inhalte_fts (
                rowid,
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
				Gezeichnet
)

SELECT
    Inhalte.rowid,
        Inhalte.id,
    Inhalte.Seite,
    Inhalte.Incipit,
    Inhalte.Titelangabe,
    Inhalte.Urheberangabe,
    Inhalte.Anmerkungen,
    group_concat(typen.value, '; ') as Typ_I,
    -- Herausgaber bei Bänden ergibt zuviele false positives
    -- (Herausgeber || " " ||  Reihentitel.Titel || " " || Baende.Jahr) as Reihe_I,
    (Reihentitel.Titel || " " || Baende.Jahr) as Reihe_I,
    -- Baende.Titelangabe,
    Schöpfer,
    Autor,
    Stecher,
    Künstler
    
FROM Inhalte, json_each(INHALTE.Typ, '$') as typen

LEFT JOIN Baende ON Inhalte.Band = Baende.id

-- Baende.Bevorzugter_Reihentitel
LEFT JOIN Reihentitel ON Baende.Bevorzugter_Reihentitel = Reihentitel.id

-- -- Baende.Herausgeber
-- LEFT JOIN
--     (SELECT
--         Baende.id as id,
--         group_concat(hrsg.Name, '; ') as Herausgeber
--     FROM 
--         Baende, 
--         json_each(Baende.Herausgabe, '$') as herausgabe
--     LEFT JOIN 
--         Akteure hrsg ON hrsg.id = herausgabe.value
--     GROUP BY
--         Baende.id) hrsgtab ON Baende.id = hrsgtab.id

-- Inhalte.Geschaffen
LEFT JOIN
    (SELECT
        Inhalte.id as id,
        group_concat(a.Name, '; ') as Schöpfer
    FROM 
        Inhalte, 
        json_each(Inhalte.Geschaffen, '$') as gesch
    LEFT JOIN 
        Akteure a ON a.id = gesch.value
    GROUP BY
        Inhalte.id) geschaffentab ON Inhalte.id = geschaffentab.id

-- Inhalte.Geschrieben
LEFT JOIN
    (SELECT
        Inhalte.id as id,
        group_concat(a.Name, '; ') as Autor
    FROM 
        Inhalte, 
        json_each(Inhalte.Geschrieben, '$') as gesch
    LEFT JOIN 
        Akteure a ON a.id = gesch.value
    GROUP BY
        Inhalte.id) geschriebentab ON Inhalte.id = geschriebentab.id

-- Inhalte.Gestochen
LEFT JOIN
    (SELECT
        Inhalte.id as id,
        group_concat(a.Name, '; ') as Stecher
    FROM 
        Inhalte, 
        json_each(Inhalte.Gestochen, '$') as gesch
    LEFT JOIN 
        Akteure a ON a.id = gesch.value
    GROUP BY
        Inhalte.id) stechertab ON Inhalte.id = stechertab.id

-- Inhalte.Gezeichnet
LEFT JOIN
    (SELECT
        Inhalte.id as id,
        group_concat(a.Name, '; ') as Künstler
    FROM 
        Inhalte, 
        json_each(Inhalte.Gezeichnet, '$') as gesch
    LEFT JOIN 
        Akteure a ON a.id = gesch.value
    GROUP BY
        Inhalte.id) kuenstlertab ON Inhalte.id = kuenstlertab.id

GROUP BY Inhalte.id
