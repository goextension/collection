package contacts

import `github.com/goextension/contacts/collection`

type Collection[T any] interface {
	Filter(func(value T, key int) bool) Collection[T]

	Reduce(func(carry []any, value T, key int) []any) []T

	Map(func(value T, key int) T) Collection[T]

	Pluck(value string, column string) collection.Arrayable

	Merge(items []T) Collection[T]

	Push() Collection[T]

	Reverse() Collection[T]

	All() []T

	Implode(separator string) string

	Max() any

	Min() any

	collection.Nullable

	collection.Countable

	collection.Arrayable

	collection.Flusher
}
