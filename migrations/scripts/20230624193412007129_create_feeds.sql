CREATE TABLE rss_feeds (
  feed_id UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
  xml_url TEXT,
  html_url TEXT,
  title TEXT,
  type VARCHAR(255),
  description TEXT,
  text TEXT,
  opml_id VARCHAR(255),
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users (user_id)
  FOREIGN KEY (opml_id) REFERENCES opml_store (opml_id)
);

CREATE TABLE opml_store (
  opml_id UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
  filename TEXT,
  title TEXT NOT NULL,
  version VARCHAR(200),
  entries_count INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT NOW()
)

CREATE feed_subscriptions (
  feed_id UUID NOT NULL,
  user_id UUID NOT NULL,
  status INTEGER DEFAULT 1,
  created_at TIMESTAMP DEFAULT NOW()
)

ALTER TABLE feed_subscriptions
ADD CONSTRAINT uq_feed_id_user_id UNIQUE (feed_id, user_id);