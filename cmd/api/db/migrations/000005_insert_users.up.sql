INSERT INTO hr.roles (role_name) VALUES 
('admin'),
('hr'),
('finance');

INSERT INTO hr.users (user_name, user_email, user_password, user_handphone, created_on) VALUES 
('admin', 'admin@code.id', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92LDg/x//q9Iy5XJU4lDq', '+628123456789', CURRENT_TIMESTAMP),  -- password: password123
('wini', 'hr@code.id', '$2a$10$W1x8z5f3K7m2P9q4R6t8Y0a2b4c6d8e0f2g4h6i8j0k2l4m6n8o0p', '+628987654321', CURRENT_TIMESTAMP),  -- password: adminpass
('widi', 'emp@code.id', '$2a$10$X3y7z9a1b5c7d9e1f3g5h7i9j1k3l5m7n9o1p3q5r7s9t1u3v5w7x', '+628111222333', CURRENT_TIMESTAMP);  -- password: hrpass


INSERT INTO hr.user_roles (user_id, role_id) VALUES 
(1, 1),  -- admin user punya role admin
(1, 2),  -- admin user juga punya role hr
(2, 2),  -- wini punya role hr
(3, 3);  -- widi punya role finance