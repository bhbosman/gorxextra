module github.com/bhbosman/gorxextra

go 1.15

require (
	github.com/bhbosman/goerrors v0.0.0-20200918064252-e47717b09c4f
	github.com/cenkalti/backoff/v4 v4.0.2 // indirect
	github.com/reactivex/rxgo/v2 v2.1.0
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/reactivex/rxgo/v2 v2.1.0 => github.com/bhbosman/rxgo/v2 v2.1.1-0.20200918073619-601ce01f6d06
