# Python
## Language convention


 * Function names should be lowercase, with words separated by underscores as necessary to improve readability.
 * Variables names: use the function naming rules.
 * Use 4 spaces per indentation level.
 * Limit all lines to a maximum of 79 characters.
 * Break line before binary operator.

## tips
### slices
Syntax : `sequence[start:stop:step]`

Example :
```python
L = range(10)
L[5:9]
# [5,6,7,8,9]
L[::2]
# [0, 2, 4, 6, 8]
L[::-1]
#[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]

a = range(3) #[0, 1, 2]
a[1:3] = [4, 5, 6]
a
# [0, 4, 5, 6]
```



### argparse

```python
import argparse
parser = argparse.ArgumentParser()
parser.add_argument('square', type=int, help='display a square of a given number')
parser.add_argument('move', choices=['rock', 'paper', 'scissors'])
parser.add_argument('-v', '--verbose', action='store_true', help='increase output verbosity')
args = parser.parse_args()
answer = args.square**2
if args.verbose:
    print 'the square of {} equals {}'.format(args.square, answer)
else:
    print answer
```
`type=` can take any callable that takes a single string argument and returns the converted value.
`choices` keyword : check against a range of values.



### files

```python
with open(newfile, 'w', encoding='utf-8') as outfile:
    outfile.write(stuff)
```
| Mode      | Signification |
| ----------|---------------|
| r | Read only (default) |
| w | Write (file overwrite) |
| a | Append |
| r+ | Read + Write  |
| rb, wb, r+b | binary mode (important for win) |

`f.write(string)` writes the contents of string to the file, returning None.

To write something other than a string, it needs to be converted to a string first: `f.write(str(value))`
