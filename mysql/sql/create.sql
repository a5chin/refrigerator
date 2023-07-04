CREATE DATABASE IF NOT EXISTS ref;
USE ref;

CREATE TABLE IF NOT EXISTS ingredients(
  id         VARCHAR(26) PRIMARY KEY,
  name       VARCHAR(40) NOT NULL,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME     NULL
);

CREATE TABLE IF NOT EXISTS nutritions(
  id         VARCHAR(26) PRIMARY KEY,
  name       VARCHAR(40) NOT NULL
);
