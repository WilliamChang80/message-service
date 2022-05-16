CREATE TABLE messageservice.messages (
    sender_id bigint PRIMARY KEY,
    message text,
    receiver_id bigint,
    group_id bigint
)