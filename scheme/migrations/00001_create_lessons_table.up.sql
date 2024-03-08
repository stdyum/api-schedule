-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS schedule.lessons
(
    id              uuid,
    study_place_id  uuid,
    group_id        uuid,
    room_id         uuid,
    subject_id      uuid,
    teacher_id      uuid,
    lesson_index    int,
    date            date,
    start_time      timestamp,
    end_time        timestamp ,
    primary_color   varchar,
    secondary_color varchar,
    created_at      timestamp,
    updated_at      timestamp,

    PRIMARY KEY ((study_place_id), date, id),
) WITH CLUSTERING ORDER BY (date DESC);

-- +goose StatementEnd
