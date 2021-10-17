## Week 2 Assignment

I have created a previous PR that added Github actions to allow for testing of each PR on each branch. This ensures that all code committed passes the correct tests. I added these in a previous PR to ensure that the tests ran against this specific PR. 

This PR provides the first test go file in the repo to run a suite of tests against three different collections. These are namely, arrays, slices and maps. The purpose of these tests is to provide adequate code coverage and ensure that each piece of code in this submission has a valid set of tests that can ensure the code passed to it will pass or fail depending upon the required tests. This however does not guarantee that all failures will be caught, simply the tests we have written. Mode will be discussed on this point further in the documentation.

# Tests

As discussed, the tests aim to provide adequate testing for all 3 collection types.

Some of the difficulties I ran into with this PR and testing specifically, was how to ensure that my tests were adequate and also that I did not simply run other built-in functions that while they may do the work, do not actually help me understand what is going on, or I should say writing the code to ensure I can prove, disprove what is actually happening. An example of this is the comparison of 2 maps. I could have used the code as follows :

```	// Perform a DeepEqual map function comparison
	mapCompare := reflect.DeepEqual(exp, act)

	// Test to ensure the DeepEqual comparison is true, otherwise fail
	if !mapCompare {
	    t.Errorf("%s is not equal to %s", act, exp)
	}
```

The code above uses a function called DeepEqual from the reflect module. This compares two maps and checks thst the maps are entirely equal. While it does a good job of ensuring that the maps are indeed equal, or not as the case may be, it kind of takes care of the issue without me either proving or disproving (to myself at least) that the maps are indeed correct. While the Go language of course takes care of this, I feel while it does provide the answer that the maps are correct, it does not answer the specific question of "Iterate through the act variable and assert that its contents match that of the exp variable, at least from a visual perspective.".

For this reason, I think it is more sensible to prove that the values are the same by using the following code :

```
	for key, _ := range act {
		_, ok := exp[key]
		if !ok {
			t.Errorf("Key %s not found in map 'exp'.", key)
		}
	}
```


# Array check

There is a single test for arrays which is the assertion that all values in the act array match those in the exp array. There is no length test in the array as arrays are fixed length, therefore this becomes a moot point. The array is defined initially with one of the mandatory attributes being the size.

The comparison of array contents is taken care of by two different tests that provide the same outcome. The first check is as follows :

```
	// Iterate through act slice and check for presence of each value in exp slice
	for i, v := range act {
		if v != exp[i] {
			t.Errorf("Value %s not found in array 'exp'.", v)
		}
	}
```
The code loops through each entry in the 'act' array and checks for the presence of the value in the 'exp' array. If a value is found that does not, exist the error will be returned by the test.

I have provided another test which provides the same outcome but uses the comparison operator on the arrays :

```
	if act != exp {
		t.Errorf("Returned value %s,  expect %s", act, exp)
	}
```



# Slice check

There are two tests for slices in TestSlice, length check and value comparison.

The first test for length uses the built-in length function of the slice collection. One of the attributes for a slice is the length of the slice. For this, I add the length of the existing slice, 'exp' by calling the len(exp) function. There is no need to define the capacity of the slice as it defaults to 4, the size that we want it to be and will not add to it therefore it is not required,

This code uses the for and range methods to loop through the 'act' slice and check the value of each index in the 'exp' slice. While this test works, there are some issues with it. 

The first issue is that if the slices are different sizes e.g. the slice we are ranging through 'act', has a length greater than 'exp', when we go to compare the slices, we can end up with a panic runtime error that shows the index being out of range. 

```
panic: runtime error: index out of range [4] with length 4 [recovered]
        panic: runtime error: index out of range [4] with length 4
```

There is a secondary issue in that the ordering also matters. If I generate a slice with the same values, but in a different order, then none of the entries match. 

This of course is due to the fact that slices are ordered lists, but just to highlight the test does not take this into account, merely that the values are the same.

The second test compares the length of the array, which works but on its own, does not guarantee that the values are the same, merely that the count of entries in the slices are equal.

This is why both tests are needed.

One other item just for note is the Copy command for copying slices. This is confusing as its destination followed by source, This works of course but is not intuitive, at least to me.

# Map check

The two tests in TestMap are similar to TestSlice, a map length comparison and also the values of each key.


# Conclusion

The testing module in this course is very interesting and shows the value of testing for different scenarios. The fact that you can add as many tests as you like is extremely useful for testing all sorts of scenarios while the benefit of an inbuilt test function makes it easy to adopt and consume.

As mentioned earlier, it is important to understand the limitations of tests and ensure adequate scenarios are covered to ensure tests are comprehensive.

# Tests output

```
=== RUN   TestArray
--- PASS: TestArray (0.00s)
=== RUN   TestSlice
--- PASS: TestSlice (0.00s)
=== RUN   TestMap
--- PASS: TestMap (0.00s)
PASS
ok      main/week02     0.554s
```

