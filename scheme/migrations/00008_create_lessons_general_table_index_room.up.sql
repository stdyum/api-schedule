-- +goose Up
-- +goose StatementBegin

CREATE INDEX ON schedule.lessons_general (room_id);

-- +goose StatementEnd
