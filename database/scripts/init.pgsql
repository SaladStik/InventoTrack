--Step 0: Drop the database if necessary
--DROP DATABASE IF EXISTS InventoTrack;

-- Step 1: Create the database
-- CREATE DATABASE InventoTrack;

-- Step 2: Connect to the database
-- \c InventoTrack;

-- Step 3: Create tables

-- Companies Table
CREATE TABLE Companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    owner_id INT UNIQUE, -- Owner of the company (linked to Users table)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories Table
CREATE TABLE Categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INT REFERENCES Categories(id) ON DELETE SET NULL,
    company_id INT NOT NULL REFERENCES Companies(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inventory Table
CREATE TABLE Inventory (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL REFERENCES Companies(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    category_id INT REFERENCES Categories(id) ON DELETE SET NULL,
    is_archived BOOLEAN DEFAULT FALSE,
    retention_period INT DEFAULT 30, -- Retention period in days for archival
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users Table
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    company_id INT REFERENCES Companies(id) ON DELETE SET NULL, -- Users can belong to a company
    role VARCHAR(50) NOT NULL CHECK (role IN ('owner', 'admin', 'user', 'viewer')), -- Role management
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Stock Table
CREATE TABLE Stock (
    id SERIAL PRIMARY KEY,
    inventory_id INT NOT NULL REFERENCES Inventory(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    threshold INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- SerialNumbers Table
CREATE TABLE SerialNumbers (
    id SERIAL PRIMARY KEY,
    stock_id INT NOT NULL REFERENCES Stock(id) ON DELETE CASCADE,
    serial_number VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Custom_Fields Table
CREATE TABLE Custom_Fields (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL REFERENCES Companies(id) ON DELETE CASCADE,
    table_name VARCHAR(50) NOT NULL, -- Table the field applies to (e.g., Inventory, Stock)
    field_name VARCHAR(50) NOT NULL, -- Name of the custom field
    field_type VARCHAR(50) NOT NULL CHECK (field_type IN ('text', 'number', 'date', 'boolean')), -- Field type
    validation_rules TEXT, -- Optional validation rules for the field
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Custom_Field_Values Table
CREATE TABLE Custom_Field_Values (
    id SERIAL PRIMARY KEY,
    custom_field_id INT NOT NULL REFERENCES Custom_Fields(id) ON DELETE CASCADE, -- Links to the custom field definition
    inventory_id INT NOT NULL REFERENCES Inventory(id) ON DELETE CASCADE, -- Links to the inventory item
    value TEXT NOT NULL, -- Value of the custom field (stored as text)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Logs Table
CREATE TABLE Logs (
    id SERIAL PRIMARY KEY,
    action TEXT NOT NULL,
    user_id INT REFERENCES Users(id) ON DELETE SET NULL,
    entity VARCHAR(50) NOT NULL,
    entity_id INT NOT NULL,
    details TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Backups Table
CREATE TABLE Backups (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL REFERENCES Companies(id) ON DELETE CASCADE,
    backup_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    backup_path TEXT NOT NULL,
    restored BOOLEAN DEFAULT FALSE
);

-- Recycle Bin Table
CREATE TABLE Recycle_Bins (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL REFERENCES Companies(id) ON DELETE CASCADE,
    table_name VARCHAR(50) NOT NULL,
    record_id INT NOT NULL,
    deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    retention_period INT DEFAULT 30 -- Retention period in days
);

-- Step 4: Add triggers for updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply updated_at triggers
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON Users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_companies_updated_at
BEFORE UPDATE ON Companies
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_inventory_updated_at
BEFORE UPDATE ON Inventory
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_categories_updated_at
BEFORE UPDATE ON Categories
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_stock_updated_at
BEFORE UPDATE ON Stock
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_serialnumbers_updated_at
BEFORE UPDATE ON SerialNumbers
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Logs Trigger (Example for Inventory)
CREATE OR REPLACE FUNCTION log_inventory_changes()
RETURNS TRIGGER AS $$
BEGIN
   INSERT INTO Logs (action, user_id, entity, entity_id, details, created_at)
   VALUES (
       TG_OP,
       NULL, -- Replace with actual user_id if available
       'Inventory',
       NEW.id,
       'Change detected in Inventory table',
       CURRENT_TIMESTAMP
   );
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER log_inventory_changes
AFTER INSERT OR UPDATE OR DELETE ON Inventory
FOR EACH ROW
EXECUTE FUNCTION log_inventory_changes();
