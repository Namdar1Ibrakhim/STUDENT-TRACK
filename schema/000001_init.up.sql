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
    description   varchar(255)
);

CREATE TABLE direction
(
    id                serial       NOT NULL UNIQUE,
    direction_name    varchar(255) NOT NULL,
    description       varchar(255)
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
    direction_id      INT          NOT NULL,
    created_at        timestamp    DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (direction_id) REFERENCES direction(id) ON DELETE CASCADE

);

-- Insert Courses
INSERT INTO course (course_name, description)
VALUES 
('Operating System', 'Study of operating systems and their functionality'),
('Analysis of Algorithm', 'Study of algorithms and their analysis'),
('Programming Concept', 'Basic programming concepts and principles'),
('Software Engineering', 'Software development processes and practices'),
('Computer Network', 'Study of computer networks and communication'),
('Applied Mathematics', 'Mathematics for computer science applications'),
('Computer Security', 'Study of computer security and cryptography');


-- Insert Directions
INSERT INTO direction (direction_name, description)
VALUES 
('Database Administrator', 'Manage and maintain database systems'),
('Data Scientist', 'Analyze data and build predictive models'),
('IT Project Manager', 'Manage and oversee IT projects'),
('Systems Administrator', 'Maintain and support system infrastructure'),
('Cybersecurity Specialist', 'Specialize in protecting information systems'),
('Software Developer', 'Develop software applications'),
('DevOps Engineer', 'Integrate and manage development and operations systems'),
('Network Engineer', 'Design and maintain computer networks');