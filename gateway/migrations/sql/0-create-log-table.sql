-- name: create-log-table
CREATE TABLE IF NOT EXISTS log (
    id int(11) NOT NULL AUTO_INCREMENT,
    title varchar(255) DEFAULT NULL,
    year varchar(255) DEFAULT NULL,
    imdb_id varchar(255) DEFAULT NULL,
    content_type varchar(255) DEFAULT NULL,
    poster varchar(255) DEFAULT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    updated_at timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
);

-- name: drop-log-table
DROP TABLE IF EXISTS log;