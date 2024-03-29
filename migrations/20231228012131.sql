CREATE TABLE IF NOT EXISTS "users_plants" (
    user_id TEXT NOT NULL REFERENCES users (id),
    plant_id TEXT NOT NULL REFERENCES plants (id),
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
