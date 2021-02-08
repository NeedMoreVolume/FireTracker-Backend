
-- +migrate Up
CREATE TABLE fires (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    name VARCHAR(255),
    description VARCHAR(255),
    start DATETIME NOT NULL,
    end DATETIME
);

CREATE TABLE weathers (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    fire_id INT,
    log_id INT,
    temp INT,
    temp_unit_id INT,
    humidity INT,
    dew_point INT,
    dew_point_unit_id INT,
    wind_speed INT,
    wind_unit_id INT,
    wind_direction VARCHAR(255),
    type VARCHAR(255),
    weather_time DATETIME NOT NULL
);

CREATE TABLE logs (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,
    fire_id INT,
    weather_id INT,
    name VARCHAR(255),
    size VARCHAR(255),
    added_at DATETIME NOT NULL
);

-- +migrate Down
DROP TABLE fires;
DROP TABLE weathers;
DROP TABLE logs;
