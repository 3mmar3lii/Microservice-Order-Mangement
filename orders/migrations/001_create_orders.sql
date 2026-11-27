--type Order struct {
--	ID         string
--	CustomerID string
--	Status     string
--	Items      []Item
--}

-- Write table for order to save data if order table not exist
CREATE TABLE IF NOT EXISTS public.orders (
  id TEXT PRIMARY KEY,
  CustomerID TEXT
)