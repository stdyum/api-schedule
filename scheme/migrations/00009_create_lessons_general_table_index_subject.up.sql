-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons_general (subject_id);

-- +goose StatementEnd
