CREATE TABLE spacecraft_armaments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    spacecraft_id INT,
    name VARCHAR(100) NOT NULL,
    qty INT NOT NULL default 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (spacecraft_id) REFERENCES spacecrafts(id)
);