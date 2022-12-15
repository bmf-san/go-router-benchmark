package main

var staticRoutesRoot = route{
	"/", "/",
}

var staticRoutes1 = route{
	"/foo", "/foo",
}

var staticRoutes5 = route{
	"/foo/bar/baz/qux/quux", "/foo/bar/baz/qux/quux",
}

var staticRoutes10 = route{
	"/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred",
}
