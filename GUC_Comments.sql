CREATE TABLE Courses
(
    CourseId INT PRIMARY KEY, 
    CourseName VARCHAR(30)
    
);

CREATE TABLE Comments
(
    CourseId INT FOREIGN KEY REFERENCES Courses, 
    Comment VARCHAR(50)
   
);