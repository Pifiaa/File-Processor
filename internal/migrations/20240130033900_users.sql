-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    Email VARCHAR(255) UNIQUE,
    Password VARCHAR(255)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF ExISTS Users
-- +goose StatementEnd
