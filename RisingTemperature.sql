select w2.Id from Weather w1 
inner join Weather w2 
on TO_DAYS(w2.Date) = TO_DAYS(w1.Date) + 1 
where w2.Temperature> w1.Temperature;
