package repository

const (
	insertRecord = `INSERT INTO fizzbuzz.stats(int1, int2, max_limit, str1, str2) 
					VALUES ($1, $2, $3, $4, $5);`
	getStats = `
					SELECT int1, int2, max_limit, str1, str2, COUNT(*) AS count
					FROM fizzbuzz.stats
					GROUP BY int1, int2, max_limit, str1, str2
					ORDER BY count DESC
					LIMIT 1;`
)
