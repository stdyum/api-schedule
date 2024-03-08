-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons (group_id);

-- +goose StatementEnd
