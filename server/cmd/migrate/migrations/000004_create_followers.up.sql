CREATE TABLE
    IF NOT EXISTS followers (
        follower_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        following_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        followed_at TIMESTAMP NOT NULL DEFAULT now (),
        PRIMARY KEY (follower_id, following_id)
    );