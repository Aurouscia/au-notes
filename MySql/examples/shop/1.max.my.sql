-- Find the number, dealer, and price of the most expensive article.

/* 【1】使用子查询查出最高的价格，并用于筛选 */
SELECT * 
FROM shop
WHERE price = (SELECT MAX(price) FROM shop);

/* 【2】使用CTE
    较繁琐，虽然mp一行一列，但不能直接被WHERE引用，需要出现在FROM中 */
WITH mp AS(
    SELECT MAX(price) AS max_price
    FROM shop)
SELECT article, dealer, price /* 避免mp.max_price出现在结果中 */
FROM shop
    CROSS JOIN mp /* 只有FROM中提到过的表名，才能被这个语句引用 */
WHERE price = mp.max_price;

/* 【3】按价格从高到低排序并LIMIT
    缺点：如果有数个并列最高的price，则无法全部展示出来 */
SELECT *
FROM shop
ORDER BY price DESC
LIMIT 1;

/* 【4】使用LEFT JOIN，找到“没有比它price更高的其他行”的这样一行
    缺点：不够直观，不易被查询优化器处理*/
SELECT s1.*
FROM shop s1
LEFT JOIN shop s2 ON s1.price < s2.price /* LEFT JOIN：s1必定完整，无法找到对应的s2则以NULL填充 */
WHERE s2.article IS NULL; /* 对应的s2的price应该更高，但不存在（说明此s1的price最高） */