/*
Copyright 2022

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

using System;
using System.Runtime;
using System.Reflection;
using System.Collections;

namespace console
{
    class Program
    {
        static void printLen(IList list) {
            Console.Out.WriteLine(list.Count);
        }

        static IList newGenericListFor(String className) {
            // We can get the Class based on a string value -- a useful feature
            // for some meta programming.
            var clazz = TypeInfo.GetType(className);

            // Build a new list constructor using the class information.
            var listCtor = typeof(List<>).MakeGenericType(clazz);

            // Construct a new List<clazz> like we did before, only this time
            // using reflection at runtime.
            //
            // Please note the result is cast to a List. This is to ensure
            // the standard List methods such as Add, and fields such as Count,
            // are accessible via the ints variable.
            return (IList) Activator.CreateInstance(listCtor);
        }

        static void Main(string[] args) {

            // Create a List<Int32>.
            var ints = newGenericListFor("System.Int32");

            // Add some numbers to the list.
            ints.Add(1);
            ints.Add(2);
            ints.Add(3);

            if (args.Length > 0 && args[0] == "fail") {
                // If this program is executed with a single argument "fail",
                // then we try to add "Hello" to the ints list.
                //
                // This will compile because ints is not known to the compiler
                // yet as a List<Int32>, even though that is what it is.
                //
                // At runtime,however, this will fail because you cannot add a
                // String to a List<Int32>.
                ints.Add("Hello");
            }

            // Create a List<String>.
            var strs = newGenericListFor("System.String");

            // Add some strings to the list.
            strs.Add("Hello");
            strs.Add("world");

            printLen(ints);
            printLen(strs);
        }
    }
}
