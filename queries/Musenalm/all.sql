-- Merge incipit & titel on Inhalte: v_anfaenge
SELECT id, Anfang, Band
FROM (
    SELECT id, Band, Incipit AS Anfang 
    FROM Inhalte 
    WHERE Anfang != "" AND Anfang NOT LIKE "#%" AND Anfang NOT LIKE "$%"
    UNION
    SELECT id, Band, Titelangabe AS Anfang 
    FROM Inhalte 
    WHERE Anfang != "" AND Anfang NOT LIKE "#%" AND Anfang NOT LIKE "$%"
)
ORDER BY Anfang COLLATE NOCASE ASC

-- Show available Years: v_b_jahre: v_b_jahre
SELECT Baende.id, Baende.Jahr
FROM Baende
WHERE Baende.Jahr > 0
GROUP BY Baende.Jahr
ORDER BY Baende.Jahr;


-- Select unique Persons & join their info on Baende: v_b_herausgabe
SELECT b_id, Akteure.id as id, Akteure.Name, Akteure.Lebensdaten, Akteure.Koerperschaft 
FROM (
    SELECT Baende.id AS b_id, value
    FROM Baende, json_each(Baende.Herausgabe, '$')
    WHERE json_array_length(Herausgabe) > 0 
    GROUP BY value
    UNION
    SELECT Baende.id AS b_id, value
    FROM Baende, json_each(Baende.Verlag, '$')
    WHERE json_array_length(Verlag) > 0 
    GROUP BY value
    UNION
    SELECT Baende.id AS b_id, value
    FROM Baende, json_each(Baende.Vertrieb, '$')
    WHERE json_array_length(Vertrieb) > 0 
    GROUP BY value
    UNION
    SELECT Baende.id AS b_id, value
    FROM Baende, json_each(Baende.Druck, '$')
    WHERE json_array_length(Druck) > 0 
    GROUP BY value
)
LEFT JOIN Akteure ON Akteure.id = value
ORDER BY Akteure.Name

-- Show available first letters of Reihentitel: v_reihentitel_letters
SELECT DISTINCT substr(Titel, 1, 1) as id 
  FROM Reihentitel
  WHERE Reihentitel.Titel IS NOT ""
  ORDER BY id COLLATE NOCASE ASC

-- Show available pseudonyms in Baende: v_b_hrsgpseu
SELECT id, Baende.Verantwortlichkeitsangabe, Baende.Herausgabe
FROM Baende
WHERE Baende.Verantwortlichkeitsangabe IS NOT ""
GROUP BY Baende.Verantwortlichkeitsangabe
ORDER BY Baende.Verantwortlichkeitsangabe COLLATE NOCASE ASC;

-- Select all Inhalte that are deemed to be presentable on the main page: v_vorschau
SELECT id, Inhalte.Band, Inhalte.Seite, Inhalte.Scans, Inhalte.Geschaffen, Inhalte.Geschrieben, Inhalte.Gezeichnet, Inhalte.Gestochen, Inhalte.Objektnummer, Inhalte.Urheberangabe, Inhalte.Anmerkungen, Inhalte.Typ, Inhalte.Titelangabe, Inhalte.Incipit
FROM Inhalte
WHERE Inhalte.Vorschau = TRUE