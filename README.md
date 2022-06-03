# Brainfuck. Reimagined

I welcome you to Brainfuck 2.

It's been far too long since the Brainfuck programming language has been updated, so i've decided to take matters into my own hands, creating a new programming language, based on the very experience that Brainfuck offered, but with a different sort of syntax.

### Same old Brainfuckery

- incredibly difficult to print anything

- numbers can only be integers

- no variables!

- requires knowledge regarding binary code!

### Even more Brainfuckery, even more power

- Every command takes approximately half a second to print, keeping you at the edge of your desk, hooked until the very end of printing your tiny ass string! (feature added for extra Brainfuckery)

- Syntax like you've never seen before (it'll make you wanna die!)

Anyways, let's skip to the good part and get started learning this fabulous language!

## Building

Building this project is pretty easy, simply `git clone https://github.com/notmax6677/Brainfuck2` and run `go build` to build the project natively.

## How to use

There is an example script that you are welcome to check out at any time, as it's essentially a template for your Brainfuck2 project!

The syntax is pretty simple, the script is split into individual commands, here's the layout of one:
`[t][input];`

`t` represents the token, which defines the action, there are currently 4 types of tokens

`!` - you need to input a whole byte that will convert to a single character, you may add `l` at the end to instantly make the character lowercase, instead of writing its lowercase counterpart in binary

`"` - here you can input a number (integers only) and it will output the integer in binary form, sorry lol (quote key for even more brainfuckery)

`+` - addition, syntax looks something like `+num1%num2;`

`-` - subtraction, syntax is similar to addition, except switch out the `+` with a `-`

the `%` character separates the two numbers (also integers only lol) in a math equation

thankfully the interpreter was feeling nice and has decided to output results of math equations in normal form, rather than binary :)

### Running Brainfuck2

Once you have built it, simply Brainfuck(2) your script by running `bf2 script.bf2`

You can get a rather surprising help message from Brainfuck(2) itself by writing `bf2 help` in your terminal!

note: it likes to curse..

anyway, enjoy your new Brainfuckery :)

