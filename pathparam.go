package main

var pathParamRoutes1Colon = route{
	"/foo/:bar", "/foo/bar",
}

var pathParamRoutes5Colon = route{
	"/foo/:bar/:baz/:qux/:quux/:corge", "/foo/bar/baz/qux/quux/corge",
}

var pathParamRoutes10Colon = route{
	"/foo/:bar/:baz/:qux/:quux/:corge/:grault/:garply/:waldo/:fred/:plugh", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}

var pathParamRoutes1Bracket = route{
	"/foo/{bar}", "/foo/bar",
}

var pathParamRoutes5Bracket = route{
	"/foo/{bar}/{baz}/{qux}/{quux}/{corge}", "/foo/bar/baz/qux/quux/corge",
}

var pathParamRoutes10Bracket = route{
	"/foo/{bar}/{baz}/{qux}/{quux}/{corge}/{grault}/{garply}/{waldo}/{fred}/{plugh}", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}

var pathParamRoutes1InequalitySign = route{
	"/foo/<bar>", "/foo/bar",
}

var pathParamRoutes5InequalitySign = route{
	"/foo/<bar>/<baz>/<qux>/<quux>/<corge>", "/foo/bar/baz/qux/quux/corge",
}

var pathParamRoutes10InequalitySign = route{
	"/foo/<bar>/<baz>/<qux>/<quux>/<corge>/<grault>/<garply>/<waldo>/<fred>/<plugh>", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}
