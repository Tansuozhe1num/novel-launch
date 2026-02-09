-- 查询完整分类树（使用递归查询，MySQL 8.0+）
WITH RECURSIVE category_tree AS (
    -- 基础查询：根节点
    SELECT
        id,
        name,
        parent_id,
        1 as level,
        CAST(id AS CHAR(500)) as path
    FROM category
    WHERE parent_id IS NULL

    UNION ALL

    -- 递归查询：子节点
    SELECT
        c.id,
        c.name,
        c.parent_id,
        ct.level + 1,
        CONCAT(ct.path, '->', c.id)
    FROM category c
    INNER JOIN category_tree ct ON c.parent_id = ct.id
)
SELECT
    id,
    LPAD('', (level-1)*4, ' ') || name as name_with_indent,
    parent_id,
    level,
    path
FROM category_tree
ORDER BY path;