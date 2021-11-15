-- using mysql 5.6
-- Simple Database querying

SELECT u.ID,
  u.UserName,
  parent.UserName AS ParentUserName
FROM user AS u
LEFT JOIN user AS parent ON u.parent = parent.id
