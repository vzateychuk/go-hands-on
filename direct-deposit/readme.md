# Task description

1. You will need to create a struct called directDeposit.
2. The directDeposit struct will have three string fields: `lastName`, `firstName`, and
`bankName`. It will also have two int fields called routingNumber and accountNumber.
3. The directDeposit struct will have a validateRoutingNumber method. The method
will return ErrInvalidRoutingNum when the routing number is less than 100.
4. The directDeposit struct will have a validateLastName method. It will return
ErrInvalidLastName when the lastName is an empty string.
5. The directDeposit struct will have a method report. It will print out each of the
fields' values.
6. In the main() function, assign values to the directDeposit struct's fields and call
each of the directDeposit struct's methods.