package repository

const (
	insertRecord = `INSERT INTO fizzbuzz.stats(int1, int2, max_limit, str1, str2, created_at) 
					VALUES ($1, $2, $3, $4, $5, now());`
	getStats = `
					SELECT COUNT(*) AS hits, int1, int2, max_limit, str1, str2 
					FROM fizzbuzz.stats
					GROUP BY int1, int2, max_limit, str1, str2
					ORDER BY hits DESC
					LIMIT 1;`
)
