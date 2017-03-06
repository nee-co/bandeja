-- +goose Up
CREATE TABLE trays (
  id INTEGER NOT NULL AUTO_INCREMENT,
  title VARCHAR(50) NOT NULL,
  endpoint VARCHAR (50) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE users_trays (
  user_id INTEGER NOT NULL,
  tray_id INTEGER NOT NULL,
  space_id INTEGER NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
DROP TABLE trays;
DROP TABLE users_trays;