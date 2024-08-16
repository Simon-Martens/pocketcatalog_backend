-- SELECT *
-- FROM Baende
-- WHERE rowid IN (SELECT rowid
--              FROM baende_fts
--              WHERE baende_fts MATCH 'Triest 1811');


SELECT *
    FROM baende_fts
    WHERE baende_fts MATCH 'Triest 1811';