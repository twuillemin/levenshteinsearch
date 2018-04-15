# What is this project for? #

This project offers a fast and efficient and efficient fuzzy search on large dictionary. To (try to) offer a good level
of performances, the two following approaches are used:
* The dictionary is stored as a Trie. So, if the search does not match, it is easy to cut directly a full branch.
* The matching is done by using a Levenshtein automaton which allow to easily propagate the state of the automaton as
the Trie is walked

The automaton is fully described by [Jules Jacobs](https://julesjacobs.github.io/2015/06/17/disqus-levenshtein-simple-and-fast.html)

# How do I get set up? #

`got get bitbucket.org/ThomasWuillemin/levenshteinsearch` should do the trick

# How do I use it? #

I tried to keep the interface as simple as possible.

* Create the dictionary
  `dict := CreateDictionary()`
* Fill the dictionary with the data
  `dict.Put("banana")`
  `dict.Put("monkey")`
* Search the data
  `result := dict.SearchAll("bana", 2)`

The result are returned as a map having:
* as key: a string which is the word
* as value: a `*WordInformation` structure which for now just have a counter representing how many times the word was
added to the dictionary

# Is it really efficient? #

The result was benched against
 * a very naive: Dictionary is stored as a simple list of strings. For each query, all string of the list is tested
 * s simple map: Dictionary is stored as a map having as key the word and as value the `*WordInformation` structure. Due
  to the map redundant word are not tested multiple times

 The search is done with text of *Alice's Adventures In Wonderland*. For the tests the sets of words <"rabbit"> and
 <"rabbit", "eart", "the"> are used. Results as follow

```
goos: windows
goarch: amd64
pkg: bitbucket.org/ThomasWuillemin/levenshteinsearch
BenchmarkNaive1Word-8                 30          53038766 ns/op        18400028 B/op     278689 allocs/op
BenchmarkNaive3Word-8                 10         143324980 ns/op        43202090 B/op     842243 allocs/op
BenchmarkMap1Word-8                  100          11824675 ns/op         3600898 B/op      36130 allocs/op
BenchmarkMap3Word-8                   50          32226752 ns/op         9775576 B/op     108589 allocs/op
BenchmarkOptimized1Word-8         500000              2608 ns/op            2368 B/op         50 allocs/op
BenchmarkOptimized3Word-8         200000              7769 ns/op            7104 B/op        150 allocs/op
```

# Next steps #

More optimizations

# Version #

* 0.1.0 - Initial version

# Contact #

Me, I suppose: Thomas Wuillemin (thomas.wuillemin _at_ gmail.com)