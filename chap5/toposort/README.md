// consider the problem of computing a sequence of computer science courses that satisfies the
// prerequisite requirements of each one. The prerequisites are given in the prereqs table below,
// which is a mapping from each course to the list of courses that must be completed before it.

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
    "algorithms": {"data structures"},
    "calculus": {"linear algebra"},
    "compilers": {
        "data structures",
        "formal languages",
        "computer organization",
    }, 
    "data structures": {"discrete math"},
    "databases": {"data structures"},
    "discrete math": {"intro to programming"},
    "formal languages": {"discrete math"},
    "networks": {"operating systems"},
    "operating systems": {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}