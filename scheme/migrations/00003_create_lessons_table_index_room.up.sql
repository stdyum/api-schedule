-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons (room_id);

-- +goose StatementEnd
