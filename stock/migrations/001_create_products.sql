CREATE TABLE IF NOT EXISTS public.products (
    id TEXT PRIMARY KEY,
    quantity INTEGER NOT NULL CHECK (quantity >= 0)
);
