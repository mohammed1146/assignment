CREATE TABLE spacecrafts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    class VARCHAR(100) NULL,
    crew INT NOT NULL default 1,
    image VARCHAR(255) NOT NULL,
    value DECIMAL(12,2) NOT NULL DEFAULT 0.00,
    status ENUM('damaged', 'operational', 'maintenance') NOT NULL DEFAULT 'operational',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,


    INDEX index_spacecraft_name (name),
    INDEX index_spacecraft_class (class),
    INDEX index_spacecraft_status (status)
);