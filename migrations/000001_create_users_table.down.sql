-- Drop the trigger from the users table first
DROP TRIGGER IF EXISTS set_timestamp ON users;

-- Drop the trigger function
DROP FUNCTION IF EXISTS trigger_set_timestamp();

-- Drop the users table
DROP TABLE IF EXISTS users;
