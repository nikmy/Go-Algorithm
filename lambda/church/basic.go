package church

// Term is an abstraction of a function
// that can be applied any number of
// times to itself. It is the base
// concept in lambda calculus. There
// will be aliases for this type for
// better readability.
type Term func(Term) Term
