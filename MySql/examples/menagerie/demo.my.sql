/* 【1】选择price>10的所有assets，并显示对应的owner的species和name */
SELECT
    assets.name AS asset_name,
    pets.name AS owner_name ,
    pets.species AS owner_species,
    assets.price
FROM pets
    RIGHT JOIN ( /* RIGHT JOIN：右边的必显示，无对应左边则留空 */
        SELECT * FROM assets
        WHERE price > 10
    ) AS assets
    ON pets.name = assets.owner 
    ORDER BY assets.price; /*记得每条语句末尾必须有分号*/

/* 【2】选择pets中所有可配对的（一雄一雌，物种一样，且都在世） */
SELECT p1.name, p1.gender, p2.name, p2.gender
FROM pets AS p1 INNER JOIN pets AS p2 /* 表可以JOIN自己 */
    ON p1.gender = 1 AND p2.gender = 0
    AND p1.species = p2.species
    AND p1.death IS NULL AND p2.death IS NULL;

/* 【3】选择pets的所有行，及每个pet拥有的asset的最高价值&平均价值&价值最高的asset名称（没有asset的留NULL） */
WITH owner_asset_stats AS ( /* WITH xxx AS ... 语法：定义一个CTE（临时结果集） */
    SELECT 
        assets.owner AS owner_name,
        MAX(assets.price) AS price_max,
        AVG(assets.price) AS price_avg,
        (  /* 相关子查询（correlated subquery）获取当前asset.owner(已被group去重)拥有的价值最高的asset的名称 */
            SELECT a2.name 
            FROM assets a2 
            WHERE a2.owner = assets.owner 
            ORDER BY a2.price DESC 
            LIMIT 1 /* 相关子查询必须返回一行一列（类似于聚合函数），否则报错 */
        ) AS highest_priced_asset_name
    FROM assets
    GROUP BY assets.owner
)
SELECT 
    p.name AS owner_name,
    p.species AS owner_species,
    oa.price_max AS asset_max_price,
    oa.highest_priced_asset_name AS asset_highest_name,
    oa.price_avg AS asset_avg_price
FROM pets p /* 起别名，其他地方可使用p引用pets */
    LEFT JOIN owner_asset_stats oa /* 起别名 */
    ON oa.owner_name = p.name;

/* 【4】查询每个asset的信息，及其owner拥有的asset.price总和 */
SELECT
    assets.id AS asset_id,
    assets.name,
    assets.owner,
    assets.price,
    SUM(price) /* OVER：窗口函数，在不改变原行数的情况下进行分组聚合等计算，每行都得到一个结果 */
        OVER (PARTITION BY assets.owner)
        AS total_price_of_its_owner
FROM assets
ORDER BY asset_id;
/* 以下为功能等价的写法，使用相关子查询（性能上不如第一种写法） */
SELECT 
    assets.id AS asset_id,
    assets.name,
    assets.owner,
    assets.price,
    (
        SELECT SUM(a2.price) 
        FROM assets a2
        WHERE a2.owner = assets.owner
    ) AS total_price_of_its_owner
FROM assets
ORDER BY asset_id;

/* 【5】同第3条的前半部分，但性能更优化（将相关子查询改为窗口函数） */
WITH `ranked_assets` AS (
    SELECT
        `owner`,
        `name`,
        `price`,
        ROW_NUMBER() OVER ( /*每个owner内部把资产按价格从高到低排，最贵的那条 rn = 1，后面依次递增 */
            PARTITION BY `owner`
            ORDER BY `price` DESC
        ) AS `rn`
    FROM `assets`
),
`owner_asset_stats` AS (
    SELECT
        `owner` AS owner_name,
        MAX(`price`) AS price_max,
        AVG(`price`) AS price_avg,
        /* CASE可过滤出rn=1的唯一一个asset名称（其他的都是NULL），MAX可作用于字符串列取字典序最大值（取到唯一一个非NULL的） */
        MAX(CASE WHEN `rn` = 1 THEN `name` END) AS highest_priced_asset_name
    FROM `ranked_assets`
    GROUP BY `owner`
)
SELECT * FROM `owner_asset_stats`;