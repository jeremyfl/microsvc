CREATE DATABASE IF NOT EXISTS productsvc;
USE productsvc;

DROP TABLE IF EXISTS products;
CREATE TABLE products(
    id          INTEGER  NOT NULL PRIMARY KEY
    ,name        VARCHAR(38) NOT NULL
    ,price       INTEGER  NOT NULL
    ,description VARCHAR(156) NOT NULL
    ,rating      NUMERIC(4,2) NOT NULL
    ,category    VARCHAR(15) NOT NULL
    ,brand       VARCHAR(26) NOT NULL
    ,thumbnail   VARCHAR(56) NOT NULL
);
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (1,'iPhone 9',549,'An apple mobile which is nothing like apple',4.69,'smartphones','Apple','https://dummyjson.com/image/i/products/1/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (2,'iPhone X',899,'SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...',4.44,'smartphones','Apple','https://dummyjson.com/image/i/products/2/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (3,'Samsung Universe 9',1249,'Samsung''s new variant which goes beyond Galaxy to the Universe',4.09,'smartphones','Samsung','https://dummyjson.com/image/i/products/3/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (4,'OPPOF19',280,'OPPO F19 is officially announced on April 2021.',4.3,'smartphones','OPPO','https://dummyjson.com/image/i/products/4/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (5,'Huawei P30',499,'Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.',4.09,'smartphones','Huawei','https://dummyjson.com/image/i/products/5/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (6,'MacBook Pro',1749,'MacBook Pro 2021 with mini-LED display may launch between September, November',4.57,'laptops','APPle','https://dummyjson.com/image/i/products/6/thumbnail.png');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (7,'Samsung Galaxy Book',1499,'Samsung Galaxy Book S (2020) Laptop With Intel Lakefield Chip, 8GB of RAM Launched',4.25,'laptops','Samsung','https://dummyjson.com/image/i/products/7/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (8,'Microsoft Surface Laptop 4',1499,'Style and speed. Stand out on HD video calls backed by Studio Mics. Capture ideas on the vibrant touchscreen.',4.43,'laptops','Microsoft Surface','https://dummyjson.com/image/i/products/8/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (9,'Infinix INBOOK',1099,'Infinix Inbook X1 Ci3 10th 8GB 256GB 14 Win10 Grey – 1 Year Warranty',4.54,'laptops','Infinix','https://dummyjson.com/image/i/products/9/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (10,'HP Pavilion 15-DK1056WM',1099,'HP Pavilion 15-DK1056WM Gaming Laptop 10th Gen Core i5, 8GB, 256GB SSD, GTX 1650 4GB, Windows 10',4.43,'laptops','HP Pavilion','https://dummyjson.com/image/i/products/10/thumbnail.jpeg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (11,'perfume Oil',13,'Mega Discount, Impression of Acqua Di Gio by GiorgioArmani concentrated attar perfume Oil',4.26,'fragrances','Impression of Acqua Di Gio','https://dummyjson.com/image/i/products/11/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (12,'Brown Perfume',40,'Royal_Mirage Sport Brown Perfume for Men & Women - 120ml',4,'fragrances','Royal_Mirage','https://dummyjson.com/image/i/products/12/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (13,'Fog Scent Xpressio Perfume',13,'Product details of Best Fog Scent Xpressio Perfume 100ml For Men cool long lasting perfumes for Men',4.59,'fragrances','Fog Scent Xpressio','https://dummyjson.com/image/i/products/13/thumbnail.webp');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (14,'Non-Alcoholic Concentrated Perfume Oil',120,'Original Al Munakh® by Mahal Al Musk | Our Impression of Climate | 6ml Non-Alcoholic Concentrated Perfume Oil',4.21,'fragrances','Al Munakh','https://dummyjson.com/image/i/products/14/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (15,'Eau De Perfume Spray',30,'Genuine  Al-Rehab spray perfume from UAE/Saudi Arabia/Yemen High Quality',4.7,'fragrances','Lord - Al-Rehab','https://dummyjson.com/image/i/products/15/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (16,'Hyaluronic Acid Serum',19,'L''OrÃ©al Paris introduces Hyaluron Expert Replumping Serum formulated with 1.5% Hyaluronic Acid',4.83,'skincare','L''Oreal Paris','https://dummyjson.com/image/i/products/16/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (17,'Tree Oil 30ml',12,'Tea tree oil contains a number of compounds, including terpinen-4-ol, that have been shown to kill certain bacteria,',4.52,'skincare','Hemani Tea','https://dummyjson.com/image/i/products/17/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (18,'Oil Free Moisturizer 100ml',40,'Dermive Oil Free Moisturizer with SPF 20 is specifically formulated with ceramides, hyaluronic acid & sunscreen.',4.56,'skincare','Dermive','https://dummyjson.com/image/i/products/18/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (19,'Skin Beauty Serum.',46,'Product name: rorec collagen hyaluronic acid white face serum riceNet weight: 15 m',4.42,'skincare','ROREC White Rice','https://dummyjson.com/image/i/products/19/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (20,'Freckle Treatment Cream- 15gm',70,'Fair & Clear is Pakistan''s only pure Freckle cream which helpsfade Freckles, Darkspots and pigments. Mercury level is 0%, so there are no side effects.',4.06,'skincare','Fair & Clear','https://dummyjson.com/image/i/products/20/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (21,'- Daal Masoor 500 grams',20,'Fine quality Branded Product Keep in a cool and dry place',4.44,'groceries','Saaf & Khaas','https://dummyjson.com/image/i/products/21/thumbnail.png');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (22,'Elbow Macaroni - 400 gm',14,'Product details of Bake Parlor Big Elbow Macaroni - 400 gm',4.57,'groceries','Bake Parlor Big','https://dummyjson.com/image/i/products/22/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (23,'Orange Essence Food Flavou',14,'Specifications of Orange Essence Food Flavour For Cakes and Baking Food Item',4.85,'groceries','Baking Food Items','https://dummyjson.com/image/i/products/23/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (24,'cereals muesli fruit nuts',46,'original fauji cereal muesli 250gm box pack original fauji cereals muesli fruit nuts flakes breakfast cereal break fast faujicereals cerels cerel foji fouji',4.94,'groceries','fauji','https://dummyjson.com/image/i/products/24/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (25,'Gulab Powder 50 Gram',70,'Dry Rose Flower Powder Gulab Powder 50 Gram • Treats Wounds',4.87,'groceries','Dry Rose','https://dummyjson.com/image/i/products/25/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (26,'Plant Hanger For Home',41,'Boho Decor Plant Hanger For Home Wall Decoration Macrame Wall Hanging Shelf',4.08,'home-decoration','Boho Decor','https://dummyjson.com/image/i/products/26/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (27,'Flying Wooden Bird',51,'Package Include 6 Birds with Adhesive Tape Shape: 3D Shaped Wooden Birds Material: Wooden MDF, Laminated 3.5mm',4.41,'home-decoration','Flying Wooden','https://dummyjson.com/image/i/products/27/thumbnail.webp');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (28,'3D Embellishment Art Lamp',20,'3D led lamp sticker Wall sticker 3d wall art light on/off button  cell operated (included)',4.82,'home-decoration','LED Lights','https://dummyjson.com/image/i/products/28/thumbnail.jpg');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (29,'Handcraft Chinese style',60,'Handcraft Chinese style art luxury palace hotel villa mansion home decor ceramic vase with brass fruit plate',4.44,'home-decoration','luxury palace','https://dummyjson.com/image/i/products/29/thumbnail.webp');
INSERT INTO products(id,name,price,description,rating,category,brand,thumbnail) VALUES (30,'Key Holder',30,'Attractive DesignMetallic materialFour key hooksReliable & DurablePremium Quality',4.92,'home-decoration','Golden','https://dummyjson.com/image/i/products/30/thumbnail.jpg');

ALTER TABLE `productsvc`.`products` ADD COLUMN `deleted_at` datetime NULL;

CREATE DATABASE IF NOT EXISTS stocksvc;
USE stocksvc;

CREATE DATABASE IF NOT EXISTS ordersvc;
USE ordersvc;