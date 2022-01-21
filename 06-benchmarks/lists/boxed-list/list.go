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

package list

// List is a new type definition for []interface{}.
type List []interface{}

// Add a new element to the list.
//
// Please note that val will be boxed in order to pass it into the method using
// the empty interface.
func (a *List) Add(val interface{}) {
	*a = append(*a, val)
}
