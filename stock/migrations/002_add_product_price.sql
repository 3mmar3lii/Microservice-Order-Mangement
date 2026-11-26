ALTER TABLE public.products
    ADD COLUMN unit_price_minor BIGINT NOT NULL CHECK (unit_price_minor >= 0),
    ADD COLUMN currency_code TEXT NOT NULL CHECK (currency_code ~ '^[A-Z]{3}$');
