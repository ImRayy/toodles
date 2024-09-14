-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT,
  status TEXT NOT NULL DEFAULT 'pending' CHECK (
    status IN ('pending', 'done', 'cancelled', 'not_done')
  ),
  priority TEXT NOT NULL DEFAULT 'normal' CHECK (priority IN ('normal', 'mid', 'high')),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  done_at DATETIME
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;

-- +goose StatementEnd
