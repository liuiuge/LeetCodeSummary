-- https://leetcode.com/problems/trips-and-users/
-- # Write your MySQL query statement below
select t2.Day, cast(ifnull(t1.cnt, 0)/t2.cnt as decimal(6,2)) as 'Cancellation Rate' from
(select Request_at as Day, count(*) as cnt from Trips t 
inner join Users c on t.Client_Id=c.Users_Id 
inner join Users d on t.Driver_Id=d.Users_id 
where c.Banned='No' and d.Banned='No' and t.Status like 'cancelled%' 
group by Request_at) t1 right join 
(select Request_at as Day, count(*) as cnt from Trips t 
inner join Users c on t.Client_Id=c.Users_Id 
inner join Users d on t.Driver_Id=d.Users_id 
where c.Banned='No' and d.Banned='No'
group by Request_at) t2 on t2.day=t1.day where t2.day between '2013-10-01' and '2013-10-03';