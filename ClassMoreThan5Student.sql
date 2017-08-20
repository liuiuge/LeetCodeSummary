SELECT class 
FROM Courses 
GROUP BY Courses.class 
HAVING COUNT(DISTINCT student)>=5;
