-- https://leetcode.com/problems/delete-duplicate-emails/

DELETE FROM Person WHERE Id NOT IN (SELECT Id FROM (SELECT MIN(Id) AS Id FROM Person GROUP BY Email) t)