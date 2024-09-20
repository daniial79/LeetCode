SELECT 
    user_id,
    CONCAT(
        UPPER(SUBSTR(name, 1, 1)),
        LOWER(SUBSTR(name,2 , LENGTH(name)))
    ) AS "name"
FROM users
ORDER BY user_id;