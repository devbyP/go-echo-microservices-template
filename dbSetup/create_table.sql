CREATE TABLE users(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  username VARCHAR(40) UNIQUE,
  hash CHAR(60) NOT NULL,
  join_date TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE roles(
  id uuid PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE user_roles(
  user_id uuid references users(id),
  role_id uuid references roles(id),
  PRIMARY KEY(role_id, user_id)
);