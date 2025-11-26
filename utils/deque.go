package utils

// Error type for Deque struct
type DequeError struct {
	msg string
}

// Error implementation for DequeError
func (de *DequeError) Error() string {
	return de.msg
}

// Double-ended queue implementation
type Deque[T any] []T

// Push value onto back of deque
func (d *Deque[T]) PushBack(val T) {
	*d = append(*d, val)
}

// Push value onto front of deque
func (d *Deque[T]) PushFront(val T) {
	*d = append([]T{val}, *d...)
}

// Pop value off back of deque
//
// Will return zero value and non-nil error if deque is empty, otherwise returns value and nil error
func (d *Deque[T]) PopBack() (T, error) {
	if len(*d) == 0 {
		var zero T
		return zero, &DequeError{msg: "Cannot pop empty deque"}
	}

	backIndex := len(*d) - 1
	back := (*d)[backIndex]
	*d = (*d)[:backIndex]

	return back, nil
}

// Pop value off front of deque
//
// Will return zero value and non-nil error if deque is empty, otherwise returns value and nil error
func (d *Deque[T]) PopFront() (T, error) {
	if len(*d) == 0 {
		var zero T
		return zero, &DequeError{msg: "Cannot pop empty deque"}
	}

	front := (*d)[0]
	*d = (*d)[1:]

	return front, nil
}

// Size of the deque
func (d Deque[T]) Size() int {
	return len(d)
}
