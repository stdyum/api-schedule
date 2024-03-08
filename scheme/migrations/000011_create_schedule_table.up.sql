-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS schedule.schedule
(
    id              UUID,
    study_place_id  UUID,
    date            date,
    status          text,
    created_at      timestamp,
    updated_at      timestamp,

    PRIMARY KEY ((study_place_id), date, status)
) WITH CLUSTERING ORDER BY (date DESC);

-- +goose StatementEnd
