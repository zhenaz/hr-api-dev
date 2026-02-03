CREATE TABLE IF NOT EXISTS hr.employee_documents (
    edoc_id serial primary key,
    file_name text,
    file_size float,
    file_type varchar(15),
    file_url text,
    employee_id int,
    created_date  timestamp,
    modified_date timestamp,
    CONSTRAINT fk_employee_edoc FOREIGN KEY (employee_id) REFERENCES hr.employees (employee_id) 
);