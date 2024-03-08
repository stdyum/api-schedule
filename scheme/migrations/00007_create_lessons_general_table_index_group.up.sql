-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons_general (group_id);

-- +goose StatementEnd
