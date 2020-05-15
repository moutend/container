CREATE TABLE bar (
  id BIGINT AUTO_INCREMENT,
  value VARCHAR(32) NOT NULL,

  PRIMARY KEY(id)
) ENGINE = InnoDB ROW_FORMAT = DYNAMIC CHARSET = utf8mb4 COLLATE utf8mb4_bin;
