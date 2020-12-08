package db

const AddNewATM = `insert into ATMs(address) values($1)`

const AddTransaction = `insert into archive(date, time, amount, sender_number, receiver_number, available_limit) 
values(($1),($2),($3),($4),($5),($6))`

const UpdateSenderAmount = `update accaunts set amount = ($1) where number = ($2)`
const UpdateReceiverAmount = `update accaunts set amount = ($1) where number = ($2)`
