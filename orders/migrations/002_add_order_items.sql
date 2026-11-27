ALTER TABLE public.orders
    ALTER COLUMN id SET DEFAULT gen_random_uuid()::text;

ALTER TABLE public.orders
    ADD COLUMN IF NOT EXISTS status TEXT NOT NULL DEFAULT 'pending';

CREATE TABLE IF NOT EXISTS public.order_items (
    order_id TEXT NOT NULL REFERENCES public.orders(id) ON DELETE CASCADE,
    item_id TEXT NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (order_id, item_id)
);
