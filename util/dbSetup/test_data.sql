INSERT INTO users (
  id,
  first_name,
  last_name,
  username,
  hash
) VALUES (
  uuid_generate_v4(),
  'test',
  'user',
  'user1',
  '1234fakeHash'
) RETURNING id;

INSERT INTO roles (
  id,
  name
) VALUES (
  uuid_generate_v4(),
  'testRole'
) RETURNING id;
