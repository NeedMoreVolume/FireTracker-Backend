
-- +migrate Up
CREATE TABLE speed_units (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    unit varchar(255)
);

CREATE TABLE temp_units (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    unit varchar(255)
);

-- +migrate Down
DROP TABLE speed_units;
DROP TABLE temp_units;
