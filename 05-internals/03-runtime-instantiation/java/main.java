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

import java.lang.Class;
import java.util.ArrayList;
import java.lang.reflect.Constructor;

class Main {
    public static void printLen(ArrayList<?> list)  {
        System.out.println(list.size());
    }

    public static void main(String[] args) throws Exception {
        // We can get the Class based on a string value -- a useful feature for
        // some meta programming.
        Class<?> intClass = Class.forName("java.lang.Integer");

        // Get the constructor for instantiating an ArrayList.
        Constructor<ArrayList> listCtor = ArrayList.class.getDeclaredConstructor();

        // But there is no way to instantiate ArrayList using the intClass!
        // Instead all we can do is create the ArrayList the way it would look
        // at runtime anyway.
        ArrayList ints = listCtor.newInstance();

        // Add some integers to the list.
        ints.add(1);
        ints.add(2);
        ints.add(3);

        // And because there is no type safety for lists created at runtime,
        // and we used reflection to mimic that here, there's no compile-time
        // type safety either!
        ints.add("Hello");

        // That said, we *can* assert some compile-time type safety back into
        // the mix by randomly declaring "strs" to be an ArrayList<String>.
        //
        // But remember, we did *nothing* to ensure the underlying storage of
        // strs is bound to the String class. Again, this is purely a compile-
        // time conveience feature.
        ArrayList<String> strs = listCtor.newInstance();
        strs.add("Hello");
        strs.add("world");

        // By asserting that strs is "ArrayList<String>", the following line
        // would cause a compile error as expected.
        // strs.add(1);

        printLen(ints);
        printLen(strs);
    }
}