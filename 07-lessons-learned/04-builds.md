# Impact to build times & file sizes

The benchmarks in this repository for [build times](../06-benchmarks/02-build-times.md) and [file sizes](../06-benchmarks/03-file-sizes.md) indicate that generics do not have a huge impact on either, except for the size of package archives. For some reason there is an issue where a package archive built from generic code is 2x the size of its non-generic variant. This is a known issue ([golang/go#50438](https://github.com/golang/go/issues/50438)), but one that does not appear to actually affect the final, executable binary.

---

Next: _That's it, thank you for reading!_
