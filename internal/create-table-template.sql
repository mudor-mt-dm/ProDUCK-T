CREATE TABLE table_1(  
    id SERIAL NOT NULL PRIMARY KEY,
    create_time DATE,
    update_time DATE,
    content VARCHAR(255)
);
COMMENT ON TABLE table_1 IS 'Это таблица 1';
COMMENT ON COLUMN table_1.content IS 'Здесь содержится содежимое';