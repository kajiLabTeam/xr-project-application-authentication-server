\c general_purpose_server;

-- applicationsテーブルにサンプルデータを手動で挿入
INSERT INTO applications (id, secret_key, name)
VALUES
    ('01F8VYXK67BGC1F9RP1E4S9YTV', '38c91691-dbe7-6ed8-a306-fa49ef16a71c', 'ABC Corp'),
    ('01F8VYXK67BGC1F9RP1E4S9YTW', '90004c51-ddd5-20c6-9e52-c78adfd32cf2', 'XYZ Corp');

-- usersテーブルにサンプルデータを手動で挿入
INSERT INTO users (id, name, mail, gender, age, height, weight, address, occupation, application_id)
VALUES
    ('01F8VYXK67BGC1F9RP1E4S9YTX', 'Alice', 'alice@example.com', 'Female', 25, 160.5, 55.0, '789 Elm St, Village', 'Engineer', '01F8VYXK67BGC1F9RP1E4S9YTV'),
    ('01F8VYXK67BGC1F9RP1E4S9YTZ', 'Bob', 'bob@example.com', 'Male', 30, 175.0, 70.0, '567 Pine St, Town', 'Teacher', '01F8VYXK67BGC1F9RP1E4S9YTW');
