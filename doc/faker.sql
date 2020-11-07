DROP TABLE IF EXISTS `tests`;

CREATE TABLE `tests` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `value` int(5) NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);

INSERT INTO tests(value)
VALUES (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);

SET FOREIGN_KEY_CHECKS=0;
TRUNCATE admins;
TRUNCATE posts;
TRUNCATE archived_posts;
TRUNCATE comments;
TRUNCATE archived_comments;
TRUNCATE categories;
TRUNCATE tags;
TRUNCATE tag_post;
SET FOREIGN_KEY_CHECKS=1;

INSERT INTO admins(name, email, password)
SELECT
  CONCAT(@rownum := @rownum + 1, 'admin'),
  CONCAT(@rownum := @rownum, 'admin@example.com'),
  '$2a$10$2XMiLL9FXbmmre7IUOl4R.YQgKY0NnP/L5GhQ8ZRpEmfB4ovL.aP6' /* password */
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

INSERT INTO categories(name)
SELECT
  CONCAT(@rownum := @rownum + 1, 'category')
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

INSERT INTO posts(admin_id, category_id, title, md_body, html_body)
SELECT
  (@rownum := @rownum + 1),
  @rownum,
  CONCAT(@rownum, 'title'),
  CONCAT(@rownum, 'md_body'),
  CONCAT(@rownum, 'html_body')
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

INSERT INTO archived_posts(admin_id, category_id, title, md_body, html_body)
SELECT
  (@rownum := @rownum + 1),
  @rownum,
  CONCAT(@rownum, 'title'),
  CONCAT(@rownum, 'md_body'),
  CONCAT(@rownum, 'html_body')
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

INSERT INTO comments(post_id, body)
SELECT
  (@rownum := @rownum + 1),
  CONCAT(@rownum, 'body')
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

INSERT INTO archived_comments(archived_post_id, body)
SELECT
  (@rownum := @rownum + 1),
  CONCAT(@rownum, 'body')
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

INSERT INTO tags(name)
SELECT
  CONCAT(@rownum := @rownum + 1, 'tag')
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

INSERT INTO tag_post(tag_id, post_id)
SELECT
  (@rownum := @rownum + 1),
  @rownum
FROM
  /* create 1000 rows */
  tests AS t1,
  tests AS t2,
  tests AS t3,
  (SELECT @rownum := 0) AS v;

DROP TABLE IF EXISTS tests;