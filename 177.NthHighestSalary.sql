-- https://leetcode.com/problems/nth-highest-salary/

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
declare m int;
set m = n-1;
  RETURN (
      # Write your MySQL query statement below.
      SELECT DISTINCT Salary from Employee
      ORDER BY Salary DESC LIMIT 1 OFFSET m
  );
END