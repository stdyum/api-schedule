-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons (teacher_id);

-- +goose StatementEnd
