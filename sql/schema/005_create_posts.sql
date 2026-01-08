-- +goose Up

CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    description TEXT,
    published_at TIMESTAMPTZ,
    feed_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', now()),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', now()),

    CONSTRAINT posts_feeds_fkey
        FOREIGN KEY (feed_id)
        REFERENCES feeds(id)

);

-- +goose Down

DROP TABLE posts;