-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS schedule.lessons
(
    id              UUID,
    study_place_id  UUID,
    subject_id      UUID,
    teacher_id      UUID,
    group_id        UUID,
    room_id         UUID,
    lesson_index    int,
    date            date,
    start_time      timestamp,
    end_time        timestamp,
    primary_color   varchar,
    secondary_color varchar,
    created_at      timestamp,
    updated_at      timestamp,

    PRIMARY KEY (id)
);

-- +goose StatementEnd
