package db

const AddNewATM = `insert into ATMs(address) values($1)`
const AddTransaction = `insert into archive(date, time, amount, sender_number, receiver_number, available_limit) values(($1),($2),($3),($4),($5),($6))`
const AddUser = `insert into users(name, surname,age, gender, login, password) values (($1),($2),($3),($4),($5),($6))`

const UpdateSenderAmount = `update accaunts set amount = ($1) where number = ($2)`
const UpdateReceiverAmount = `update accaunts set amount = ($1) where number = ($2)`

const CheckBalance = `select amount from accaunts where user_id = ($1)`
const SenderAmount = `select amount, number from accaunts where user_id =($1)`
const ReceiverAmount = `select amount from accaunts where number = ($2)`
const Archive = `select id, date,time,amount, sender_number, receiver_number, available_limit from archive  where sender_number = ($1)`
const ATMs = `select id, address from ATMs`
const ShowUsers = `select id, age, name, surname from users`
