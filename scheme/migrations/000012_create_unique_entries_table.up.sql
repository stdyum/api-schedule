-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS schedule.unique_entries
(
    id             VARCHAR,
    study_place_id UUID,
    group_id       UUID,
    subject_id     UUID,
    teacher_id     UUID,
    created_at     timestamp,
    updated_at     timestamp,

    PRIMARY KEY ((study_place_id), id)
) WITH CLUSTERING ORDER BY (id ASC);

-- +goose StatementEnd
