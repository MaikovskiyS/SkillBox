CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR(255),
    age INTEGER
    );

  CREATE TABLE IF NOT EXISTS friends(
    user_id INTEGER,
    friend_id INTEGER
  )  