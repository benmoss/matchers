# matchers

Provides right now only a single [Gomega](https://onsi.github.io/gomega/) matcher for deep equality that has the benefit 
of providing an explanation for the inequality when it fails. 

Uses the DeepEqual function from [juju/testing](https://github.com/juju/testing) so has all the same 
caveats as [their implementation](https://godoc.org/github.com/juju/testing/checkers#DeepEqual).
