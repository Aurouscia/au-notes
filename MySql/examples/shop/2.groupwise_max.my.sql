-- For each article, find the dealer or dealers with the most expensive price.

/* 【1】相关子查询 */
SELECT s1.article, s1.dealer, s1.price
FROM shop s1
WHERE price = ( /* 对于外层查询的每一行，内层查询都会针对当前行的s1.article执行一次（性能不行） */
    SELECT MAX(s2.price)
    FROM shop s2
    WHERE s1.article = s2.article
)
ORDER BY article;

/* 【2】连接子查询 */
SELECT s1.article, s1.dealer, s1.price
FROM shop s1
    JOIN ( /* 相当于INNER JOIN */
        SELECT article, MAX(price) AS price /* group by时，dealer列无法聚合 */
        FROM shop
        GROUP BY article
    ) AS s2
    ON s1.article = s2.article AND s1.price = s2.price /* 只能通过join的方式重新得到dealer */
ORDER BY article;

/* 【3】LEFT JOIN，join出article相同的情况下price最高的（没更高的了） */
SELECT s1.article, s1.dealer, s1.price
FROM shop s1
    LEFT JOIN shop s2
    ON s1.article = s2.article AND s1.price < s2.price
WHERE s2.article IS NULL
ORDER BY article;

/* 【4】CTE，按每个article分组并排序，最后只取排名为1的那些 */
WITH s1 AS (
    SELECT 
        article, dealer, price,
        RANK() OVER ( /*窗口函数RANK与ROW_NUMBER的区别：前者会给排名相同的行相同排名，后者会保证不重复*/
            PARTITION BY article
            ORDER BY price DESC
        ) AS `rank`
    FROM shop
)
SELECT article, dealer, price
    FROM s1
    WHERE `rank` = 1
ORDER BY article;