-- SELECT * FROM inhalte_fts
-- WHERE inhalte_fts MATCH "Licht Und"


SELECT *
FROM Inhalte
WHERE rowid IN (SELECT rowid
             FROM inhalte_fts
             WHERE inhalte_fts MATCH 'süss* LICHT* UND*')