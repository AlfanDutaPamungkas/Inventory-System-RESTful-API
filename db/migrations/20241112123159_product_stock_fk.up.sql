ALTER TABLE product_stock
ADD CONSTRAINT fk_product_stock_products FOREIGN KEY (SKU)
REFERENCES products (SKU);