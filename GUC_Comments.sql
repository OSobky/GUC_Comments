CREATE TABLE Courses
(
    CourseId SERIAL PRIMARY KEY,
    CourseName VARCHAR(30)
    
);

CREATE TABLE Comments
(
    CourseId INTEGER REFERENCES Courses(CourseId), 
    Comment VARCHAR(50)
   
);