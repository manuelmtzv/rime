CREATE TABLE IF NOT EXISTS tag_writing (
  tag_id uuid NOT NULL,
  writing_id uuid NOT NULL,
  PRIMARY KEY (tag_id, writing_id),
  FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE,
  FOREIGN KEY (writing_id) REFERENCES writings (id) ON DELETE CASCADE
);