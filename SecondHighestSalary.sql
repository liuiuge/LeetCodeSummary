select max(Salary) as SecondHighestSalary 
from Employee 
where Salary < (select max(Salary) as SecondHighestSalary 
from Employee)
