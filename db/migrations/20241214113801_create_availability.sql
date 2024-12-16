-- migrate:up
CREATE TYPE day_enum AS ENUM (
    'monday',
    'tuesday',
    'wednesday',
    'thursday',
    'friday',
    'saturday',
    'sunday'
);

CREATE TABLE day_availabilities (
    id UUID PRIMARY KEY,
    day day_enum NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    slots JSONB NOT NULL, -- [{"start": 720, "end": 900}] stores time in "minutes since midnight"
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, day)
);

CREATE TABLE date_availabilities (
    id UUID PRIMARY KEY,
    date DATE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    slots JSONB NOT NULL, -- same as slots in day_availabilities
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, date)
);

-- migrate:down
DROP TABLE IF EXISTS date_availabilities;
DROP TABLE IF EXISTS day_availabilities;
DROP TYPE IF EXISTS day_enum;
