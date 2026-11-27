ALTER TABLE public.order_items
    ADD COLUMN IF NOT EXISTS unit_price_minor BIGINT
        CHECK (unit_price_minor >= 0);

ALTER TABLE public.order_items
    ADD COLUMN IF NOT EXISTS currency_code TEXT
        CHECK (currency_code ~ '^[A-Z]{3}$');
