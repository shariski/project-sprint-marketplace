CREATE TABLE users (
  id INT GENERATED BY DEFAULT AS IDENTITY,
  username varchar(15) UNIQUE NOT NULL,
  name varchar(50) NOT NULL,
  password varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  PRIMARY KEY(id)
);