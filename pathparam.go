package main

var pathParamColonRoutes1 = route{
	"/foo/:bar", "/foo/bar",
}

var pathParamColonRoutes5 = route{
	"/foo/:bar/:baz/:qux/:quux/:corge", "/foo/bar/baz/qux/quux/corge",
}

var pathParamColonRoutes10 = route{
	"/foo/:bar/:baz/:qux/:quux/:corge/:grault/:garply/:waldo/:fred/:plugh", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}

var pathParamBracketRoutes1 = route{
	"/foo/{bar}", "/foo/bar",
}

var pathParamBracketRoutes5 = route{
	"/foo/{bar}/{baz}/{qux}/{quux}/{corge}", "/foo/bar/baz/qux/quux/corge",
}

var pathParamBracketRoutes10 = route{
	"/foo/{bar}/{baz}/{qux}/{quux}/{corge}/{grault}/{garply}/{waldo}/{fred}/{plugh}", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}

var pathParamInequalitySignRoutes1 = route{
	"/foo/<bar>", "/foo/bar",
}

var pathParamInequalitySignRoutes5 = route{
	"/foo/<bar>/<baz>/<qux>/<quux>/<corge>", "/foo/bar/baz/qux/quux/corge",
}

var pathParamInequalitySignRoutes10 = route{
	"/foo/<bar>/<baz>/<qux>/<quux>/<corge>/<grault>/<garply>/<waldo>/<fred>/<plugh>", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}
