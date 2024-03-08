-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons_general (teacher_id);

-- +goose StatementEnd
