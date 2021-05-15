BEGIN;

DROP INDEX books_name_uniq_idx;
DROP TRIGGER ipmonitors_updated_at on ipmonitors;
DROP TABLE books;
DROP TABLE genres;

COMMIT;
