CREATE TABLE nugget(
  id SERIAL PRIMARY KEY,
  author_id SERIAL,
  category_id SERIAL,
  title VARCHAR,
  body TEXT,
  created_at TIMESTAMP,
  edited_at TIMESTAMP,

  FOREIGN KEY (author_id) REFERENCES author(id),
  FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE author(
  id SERIAL PRIMARY KEY,
  username VARCHAR,
  email VARCHAR
);

CREATE TABLE category(
  id SERIAL PRIMARY KEY,
  name VARCHAR
);
