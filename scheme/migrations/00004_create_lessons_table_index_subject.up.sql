-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons (subject_id);

-- +goose StatementEnd
