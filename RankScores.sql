select s.Score ,count(distinct t.score)Rank
From Scores s Join Scores t On s.Score <= t.score
Group by s.Id
Order by s.Score desc;
