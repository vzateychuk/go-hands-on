# Black-box тестирование ретейлового сервиса, использующего микросервисы для расчета цены и налога

The current PriceCalculator is implementation of the _Facade_ pattern.

For the purpose of writing our black-box tests, we will start by examining
the public interface of the retail package. A quick browse of
the retail.go file reveals a NewPriceCalculator function that receives the
URLs to the price and vat services as arguments and returns
a PriceCalculator instance. The calculator instance can be used to obtain an
item's VAT-inclusive price by invoking the PriceForItem method and
passing the item's UUID as an argument. On the other hand, if we are
interested in obtaining a VAT-inclusive item price for a particular date in
the past, we can invoke the PriceForItemAtDate method, which also accepts a
time period argument.

## Running test with coverage
```shell
go test -v -tags all_tests -coverprofile=coverage.txt -covermode=atomic ./...
```