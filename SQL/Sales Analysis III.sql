SELECT
    p.product_id,
    p.product_name
FROM product p 
LEFT JOIN sales s ON  p.product_id = s.product_id
GROUP BY 1, 2
HAVING MIN(s.sale_date) >= '2019-01-01' AND MAX(s.sale_date) <= '2019-03-31';