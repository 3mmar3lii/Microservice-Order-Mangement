INSERT INTO public.products (id, name, quantity, unit_price_minor, currency_code)
VALUES
    ('SKU-001', 'Wireless Mouse', 25, 2499, 'USD'),
    ('SKU-002', 'Mechanical Keyboard', 15, 7999, 'USD'),
    ('SKU-003', 'USB-C Cable', 50, 999, 'USD'),
    ('SKU-004', 'Laptop Stand', 18, 3499, 'USD'),
    ('SKU-005', '1080p Webcam', 12, 4999, 'USD'),
    ('SKU-006', 'USB Hub', 20, 2799, 'USD'),
    ('SKU-007', 'Wireless Headphones', 16, 5999, 'USD'),
    ('SKU-008', 'Desk Lamp', 22, 3199, 'USD'),
    ('SKU-009', '24-inch Monitor', 8, 12999, 'USD'),
    ('SKU-010', 'Mouse Pad', 30, 1499, 'USD'),
    ('SKU-011', 'Ethernet Adapter', 14, 2199, 'USD'),
    ('SKU-012', 'Portable SSD 1TB', 9, 10999, 'USD'),
    ('SKU-013', 'Bluetooth Speaker', 13, 4599, 'USD'),
    ('SKU-014', 'USB-C Charger', 40, 1899, 'USD'),
    ('SKU-015', 'Laptop Sleeve', 17, 2699, 'USD'),
    ('SKU-016', 'HDMI Cable', 35, 1299, 'USD'),
    ('SKU-017', 'USB Microphone', 10, 6999, 'USD'),
    ('SKU-018', 'Smart Plug', 24, 1799, 'USD'),
    ('SKU-019', 'Power Bank', 11, 3999, 'USD'),
    ('SKU-020', 'Monitor Arm', 7, 8999, 'USD')
ON CONFLICT (id) DO NOTHING;
