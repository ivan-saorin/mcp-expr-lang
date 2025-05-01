# Expr Evaluation Tool Guide

The `eval` tool allows you to evaluate expressions using the powerful expr language. This guide demonstrates various examples to showcase its capabilities.

## Basic Operations

### Arithmetic Operations
```
eval("10 + 5")         // Result: 15
eval("10 - 5")         // Result: 5
eval("10 * 5")         // Result: 50
eval("10 / 5")         // Result: 2
eval("10 % 3")         // Result: 1
eval("2 ^ 3")          // Result: 8 (exponentiation)
```

### String Operations
```
eval("'Hello' + ' ' + 'World'")  // Result: "Hello World"
eval("len('Hello')")             // Result: 5
```

### Boolean Operations
```
eval("true and false")           // Result: false
eval("true or false")            // Result: true
eval("!true")                    // Result: false
eval("10 > 5")                   // Result: true
eval("10 == 10")                 // Result: true
eval("'a' in ['a', 'b', 'c']")   // Result: true
```

## Working with Arrays

### Array Creation and Access
```
eval("[1, 2, 3, 4, 5]")          // Result: [1, 2, 3, 4, 5]
eval("[1, 2, 3][1]")             // Result: 2 (zero-based indexing)
```

### Array Functions
```
eval("len([1, 2, 3, 4, 5])")     // Result: 5
eval("map([1, 2, 3], # * 2)")    // Result: [2, 4, 6]
eval("filter([1, 2, 3, 4], # > 2)") // Result: [3, 4]
eval("all([1, 2, 3], # > 0)")    // Result: true
eval("any([1, 2, 3], # > 2)")    // Result: true
eval("sum([1, 2, 3, 4])")        // Result: 10
```

### Sorting Arrays
```
eval("sort([3, 1, 4, 2])")       // Result: [1, 2, 3, 4]
```

## Working with Maps/Objects

### Map Creation and Access
```
eval("{a: 1, b: 2, c: 3}")       // Result: {"a": 1, "b": 2, "c": 3}
eval("{a: 1, b: 2}.a")           // Result: 1
```

## Complex Examples

### Filtering and Transforming Data
```
users = [{"Name": "John", "Age": 30}, {"Name": "Ivan", "Age": 51}, {"Name": "Eve", "Age": 15}]
eval("filter(users, .Age >= 18)")  
// Result: [{"Name": "John", "Age": 30}, {"Name": "Ivan", "Age": 51}]

eval("map(users, {name: .Name, isAdult: .Age >= 18})")
// Result: [{"name": "John", "isAdult": true}, {"name": "Ivan", "isAdult": true}, {"name": "Eve", "isAdult": false}]
```

### Sorting Complex Data
```
users = [{"Name": "John", "Age": 30}, {"Name": "Ivan", "Age": 51}, {"Name": "Eve", "Age": 15}]
eval("sortBy(users, .Age, 'desc')")
// Result: [{"Name": "Ivan", "Age": 51}, {"Name": "John", "Age": 30}, {"Name": "Eve", "Age": 15}]

eval("sortBy(users, .Name)")
// Result: [{"Name": "Eve", "Age": 15}, {"Name": "Ivan", "Age": 51}, {"Name": "John", "Age": 30}]
```

### Aggregation Functions
```
products = [{"Name": "Apple", "Price": 0.5, "Quantity": 10}, {"Name": "Banana", "Price": 0.3, "Quantity": 15}, {"Name": "Orange", "Price": 0.6, "Quantity": 8}]
eval("sum(map(products, .Price * .Quantity))")
// Result: 13.3 (total value of inventory)

eval("reduce(products, 0, acc + .Price * .Quantity)")
// Result: 13.3 (same calculation using reduce)
```

### Conditional Logic
```
users = [{"Name": "John", "Age": 30}, {"Name": "Ivan", "Age": 51}, {"Name": "Eve", "Age": 15}]
eval("map(users, {name: .Name, status: .Age < 18 ? 'minor' : (.Age >= 65 ? 'senior' : 'adult')})")
// Result: [{"name": "John", "status": "adult"}, {"name": "Ivan", "status": "adult"}, {"name": "Eve", "status": "minor"}]
```

### String Manipulation
```
eval("upper('hello')")           // Result: "HELLO"
eval("lower('HELLO')")           // Result: "hello"
eval("trim(' hello ')")          // Result: "hello"
eval("split('a,b,c', ',')")      // Result: ["a", "b", "c"]
```

### Date Functions
```
eval("now()")                    // Result: current timestamp
eval("now().Year")               // Result: current year
eval("now().Format('2006-01-02')") // Result: current date formatted as YYYY-MM-DD
```

## Advanced Features

### Function Composition
```
numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
eval("sum(filter(numbers, # % 2 == 0))")
// Result: 30 (sum of even numbers)
```

### Custom Environment Variables
You can provide custom variables in the environment:
```
// With env: {"x": 10, "y": 20}
eval("x + y")                    // Result: 30
```

### Regular Expressions
```
eval("match('hello world', 'hello')")  // Result: true
eval("replace('hello world', 'hello', 'hi')") // Result: "hi world"
```

## Tips for Using the Eval Tool

1. Use single quotes for strings within the expression
2. Use dot notation to access object properties
3. Use `#` as a placeholder in map/filter functions
4. Remember that array indexing is zero-based
5. For complex expressions, break them down into smaller parts

This guide covers many of the powerful features of the expr language. For more details and advanced usage, refer to the [official documentation](https://expr-lang.org/docs/language-definition).