CREATE TRIGGER tbl_ai AFTER INSERT ON tbl BEGIN
  INSERT INTO fts_idx(rowid, b, c) VALUES (new.a, new.b, new.c);
END