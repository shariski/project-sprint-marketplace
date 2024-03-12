CREATE TABLE payments (
  id INT GENERATED BY DEFAULT AS IDENTITY,
  product_id int NOT NULL,
  bank_account_id int NOT NULL,
  payment_proof_image_url varchar(255) NOT NULL,
  quantity int NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk_product 
    FOREIGN KEY(product_id) 
    REFERENCES products(id)
	  ON DELETE SET NULL,
  CONSTRAINT fk_bank_account
    FOREIGN KEY(bank_account_id) 
    REFERENCES bank_accounts(id)
	  ON DELETE SET NULL
);