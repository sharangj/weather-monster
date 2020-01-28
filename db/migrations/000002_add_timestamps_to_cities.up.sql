ALTER TABLE cities
ADD COLUMN created_at timestamp NOT NULL,
ADD COLUMN updated_at timestamp NOT NULL,
ADD COLUMN deleted_at timestamp NOT NULL;
