BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "box" (
                                     "ID_box"	INTEGER,
                                     "name"	TEXT,
                                     PRIMARY KEY("ID_box" AUTOINCREMENT)
    );
CREATE TABLE IF NOT EXISTS "products" (
                                          "ID_product"	INTEGER,
                                          "name"	TEXT,
                                          "price"	INTEGER,
                                          PRIMARY KEY("ID_product" AUTOINCREMENT)
    );
CREATE TABLE IF NOT EXISTS "orders" (
                                        "ID_order"	INTEGER,
                                        "client_name"	TEXT,
                                        "phone"	TEXT,
                                        PRIMARY KEY("ID_order" AUTOINCREMENT)
    );
CREATE TABLE IF NOT EXISTS "box-products" (
                                              "ID_box"	INTEGER,
                                              "ID_products"	INTEGER,
                                              FOREIGN KEY("ID_box") REFERENCES "box"("ID_box"),
    PRIMARY KEY("ID_box","ID_products")
    );
CREATE TABLE IF NOT EXISTS "orders-products" (
                                                 "ID_order"	INTEGER,
                                                 "ID_product"	INTEGER,
                                                 FOREIGN KEY("ID_product") REFERENCES "products"("ID_product"),
    PRIMARY KEY("ID_order","ID_product")
    );
INSERT INTO "box" VALUES (1,'A');
INSERT INTO "box" VALUES (2,'B');
INSERT INTO "products" VALUES (1,'SAMSUNG TV',200000);
INSERT INTO "products" VALUES (2,'PHILIPS TV',100000);
INSERT INTO "products" VALUES (3,'SONY TV',150000);
INSERT INTO "orders" VALUES (1,'self','88005553535');
INSERT INTO "box-products" VALUES (1,1);
INSERT INTO "orders-products" VALUES (1,1);
COMMIT;