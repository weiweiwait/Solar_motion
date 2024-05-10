CREATE TABLE user(
                      id INT PRIMARY KEY,
                      username VARCHAR(255) NOT NULL,
                      password VARCHAR(255) NOT NULL,
                      phone_number VARCHAR(255) NOT NULL ,
                      qq VARCHAR(255) NOT NULL ,
                      avatar VARCHAR(255),
                      create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      integral INT,
                      awardHistory TEXT,
                      INDEX phone_number_index (phone_number),
);
CREATE TABLE UserCheckin(
                            user_id INT,
                            checkin_date DATE,
                            FOREIGN KEY (user_id) REFERENCES user(id)
);