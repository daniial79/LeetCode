SELECT
    stock_name,
    SUM(IF(operation = 'Sell', price, -1 * price)) AS capital_gain_loss
FROM stocks
GROUP BY stock_name;