-- https://leetcode.com/problems/department-highest-salary/

select a.Name as Department, b.Name as Employee, b.Salary as Salary 
from Employee as b 
inner join 
Department as a on a.Id = b.DepartmentId
where (b.DepartmentId, b.Salary) in 
(select DepartmentId, max(Salary) from Employee group by DepartmentId)