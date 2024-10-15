CREATE TABLE users
(
    id            serial       primary key ,
    firstname     varchar(255) NOT NULL,
    lastname      varchar(255) NOT NULL,
    username      varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    role          INT          NOT NULL,
    CHECK ( role >= 1 and role <=3 )
);

CREATE TABLE course
(
    id            serial       primary key ,
    course_name   varchar(255) NOT NULL,
    description   varchar(255) default 'Description will be added soon...'
);

CREATE TABLE direction
(
    id                serial       primary key ,
    direction_name    varchar(255) NOT NULL,
    description       varchar(255) default 'Description will be added soon...'
);

CREATE TABLE student_course
(
    id            serial       primary key ,
    student_id    INT          NOT NULL,
    course_id     INT          NOT NULL,
    grades        INT          NOT NULL,
    FOREIGN KEY (student_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (course_id)  REFERENCES course(id) ON DELETE CASCADE,
    CHECK ( grades >=0 )
);

CREATE TABLE prediction
(
    id                serial       primary key ,
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
('Network Engineer', 'Design and maintain computer networks'),
('IT Support Specialist');

CREATE INDEX idx_course_name ON course(course_name);
CREATE INDEX idx_direction_name ON direction(direction_name);