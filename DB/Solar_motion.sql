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
                            FOREIGN KEY (user_id) REFERENCES manager(id)
);
CREATE TABLE UserSport(
                            user_id INT,
                            sport_date DATE,
                            FOREIGN KEY (user_id) REFERENCES manager(id)
);
CREATE TABLE  Prize(
                       id INT PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       describ TEXT,
                       sum  INT,
                       start_date DATE,
                       end_date DATE,
                       status INT
);
CREATE TABLE UserDate(
                          user_id INT,
                          draw_date DATE,
                          FOREIGN KEY (user_id) REFERENCES manager(id)
);
CREATE TABLE UserApply(
                         user_id INT,
                         prize_id INT,
                         name VARCHAR(255) NOT NULL,
                         create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE carry_prize(
                            user_id INT,
                            name VARCHAR(255) NOT NULL,
                            create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE manager(
                    id INT AUTO_INCREMENT PRIMARY KEY,
                     username VARCHAR(255) NOT NULL,
                     password VARCHAR(255) NOT NULL,
                     phone_number VARCHAR(255) NOT NULL ,
                     avatar VARCHAR(255),
                     create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     INDEX phone_number_index (phone_number)
);
CREATE TABLE carry_integral(
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            user_id INT,
                            active_id INT,
                            name VARCHAR(255) NOT NULL,
                            create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);