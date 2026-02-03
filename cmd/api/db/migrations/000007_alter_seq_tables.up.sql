alter sequence hr.regions_region_id_seq
restart with 10 increment by 1;

alter sequence hr.locations_location_id_seq
restart with 10 increment by 1;

alter sequence hr.jobs_job_id_seq
restart with 10 increment by 1;

alter sequence hr.dependents_dependent_id_seq
restart with 50 increment by 1;

alter sequence HR.departments_department_id_seq
restart with 50 increment by 1;

alter sequence hr.employees_employee_id_seq
restart with 300 increment by 1;