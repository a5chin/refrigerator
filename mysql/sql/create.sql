CREATE DATABASE IF NOT EXISTS ref;
USE ref;

CREATE TABLE IF NOT EXISTS ingredients(
  id         VARCHAR(26)  PRIMARY KEY,
  name       VARCHAR(40)  NOT NULL,
  weight      INT UNSIGNED NOT NULL DEFAULT 0,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME     NULL
);

CREATE TABLE IF NOT EXISTS nutritions(
  id         VARCHAR(26) PRIMARY KEY,
  name       VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredient_nutritions(
  ingredient_id VARCHAR(26) NOT NULL,
  nutrition_id  VARCHAR(26) NOT NULL
);
