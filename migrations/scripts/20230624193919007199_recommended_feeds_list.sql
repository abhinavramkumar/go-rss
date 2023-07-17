create table recommended_feeds_list (
  feed_id UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
  feed_url VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  score INT NOT NULL
);