CREATE TABLE IF NOT EXISTS attendances (
       id INTEGER(255) PRIMARY KEY AUTO_INCREMENT,
       user_id integer NOT NULL,
       clock_in_time timestamp NOT NULL,
       clock_out_time timestamp,
       created_at timestamp,
       updated_at timestamp
)