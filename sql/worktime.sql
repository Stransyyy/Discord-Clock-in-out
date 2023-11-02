CREATE DATABASE 'clock_in';

DROP TABLE IF EXIST 'Worktime';
CREATE TABLE 'Worktime'(
    'ID' INT,
    'Time' Time,
    PRIMARY KEY('ID')
);

INSERT INTO 'Worktime' ('ID', 'TIME') VALUES 
    (0001, '10:00:00'),
    (0002, '9:50:00'),
    (0003, '10:10:00'),
    (0004, '10:05:00'),
    (0005, '10:00:00'),
    (0006, '9:55:00');

DROP TABLE IF EXIST 'clock_out';
CREATE TABLE 'clock_out' (
    'ID' INT,
    'Time' TIME,
    PRIMARY KEY('ID')
);

INSERT INTO 'clock_out' ('ID', 'TIME') VALUES 
(0001, 17:00:00),
(0002, 17:00:00),
(0003, 17:00:00),
(0004, 17:00:00),
(0005, 17:00:00),
(0006, 17:00:00);


