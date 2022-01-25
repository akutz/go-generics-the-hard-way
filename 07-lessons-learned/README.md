# Lessons learned

The number one question I, and almost certainly many others, had/have on their mind regading Go generics is _When should I, if ever, consider adopting this new language feature, in both greenfield and brownfield projects?_

The answer, of course, is not a simple one. That is why this repository exists -- it represents the journey one engineer took to try and understand generics in Go and where and how they can help developers, development, and deliveriables.

* [**Generics are suited for container patterns**](./01-container-patterns.md): container patterns are finally sorted
* [**Generics _can_ improve performance by eliminating boxing**](./02-eliminating-boxing.md): get rid of the empty interface
* [**Marshal/unmarshal patterns will not immediately benefit**](./03-marshal-unmarshal.md): boxing anywhere is still boxing
* [**Impact to build times & file sizes is negligible (so far)**](./04-builds.md): watch that compiler go!

---

Next: [Container patterns](./01-container-patterns.md)
