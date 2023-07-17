CREATE TABLE opml_store (
  opml_id UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
  filename TEXT,
  title TEXT NOT NULL,
  version VARCHAR(200),
  entries_count INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE rss_feeds (
  feed_id UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
  xml_url TEXT,
  html_url TEXT,
  title TEXT,
  type VARCHAR(255),
  description TEXT,
  text TEXT,
  opml_id UUID,
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (opml_id) REFERENCES opml_store (opml_id)
);

CREATE TABLE feed_subscriptions (
  feed_id UUID NOT NULL,
  user_id UUID NOT NULL,
  status INTEGER DEFAULT 1,
  created_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE feed_subscriptions
ADD CONSTRAINT uq_feed_id_user_id UNIQUE (feed_id, user_id);

CREATE TABLE categories (
  category_id UUID UNIQUE NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  user_id UUID,
  feed_id UUID
);

create table recommended_feeds_list (
  feed_id UUID PRIMARY KEY NOT NULL UNIQUE,
  score INT NOT NULL,
  FOREIGN KEY (feed_id) REFERENCES rss_feeds (feed_id)
);