-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS schedule.lessons_general
(
    id              uuid,
    study_place_id  uuid,
    subject_id      uuid,
    teacher_id      uuid,
    group_id        uuid,
    room_id         uuid,
    lesson_index    int,
    start_time      time,
    end_time        time,
    day_index       int,
    primary_color   varchar,
    secondary_color varchar,
    created_at      timestamp,
    updated_at      timestamp,

    PRIMARY KEY ((study_place_id), day_index, id)
) WITH CLUSTERING ORDER BY (day_index DESC);


-- +goose StatementEnd
