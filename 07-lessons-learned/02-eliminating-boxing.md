# Eliminating boxing

The other huge benefit generics bring is the elimination of boxing. This repository's section on [boxing benchmarks](../06-benchmarks/01-boxing.md) revealed that eliminating boxing can speed up execution by a factor of 10 while consuming half the memory!

However, before you go wanting to rewrite all of those marshal/unmarshal functions that take empty interfaces, well...

---

Next: [Marshal/unmarshal](./03-marshal-unmarshal.md)
