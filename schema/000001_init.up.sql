CREATE TABLE users
(
    id            serial       NOT NULL UNIQUE,
    firstname     varchar(255) NOT NULL,
    lastname      varchar(255) NOT NULL,
    username      varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    role          INT          NOT NULL
);

CREATE TABLE course
(
    id            serial       NOT NULL UNIQUE,
    course_name   varchar(255) NOT NULL,
    description   varchar(255) NOT NULL
);

CREATE TABLE student_course
(
    id            serial       NOT NULL UNIQUE,
    student_id    INT          NOT NULL,
    course_id     INT          NOT NULL,  
    grades        INT          NOT NULL,
    FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (course_id)  REFERENCES course(id) ON DELETE CASCADE
);

CREATE TABLE prediction
(
    id                serial       NOT NULL UNIQUE,
    student_id        INT          NOT NULL,
    prediction_text   varchar(255) NOT NULL,
    created_at        timestamp    DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE
);
