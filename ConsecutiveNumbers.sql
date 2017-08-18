select distinct l1.Num as ConsecutiveNums
from Logs as l1, logs as l2, logs as l3
where 
    (l1.id= l2.id-1 and l1.id = l3.id-2)    
and 
    (l1.Num = l2.Num and l1.Num = l3.Num);
