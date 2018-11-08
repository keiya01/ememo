# ememo
This tool makes it easy to manage TODO from the command line.

## Command

- `--text || -t`
    - Set contents to text file.Please put the file name in the first argument.
    - After entering the content you will be asked for the TODO you want to register, please enter the TODO.

- `--read || -r`
    - show text file contents.Please put the file name in the first argument.

- `--comp || -c`
    - Check the completed TODO.Please enter the file name as an argument.
    - When entering the file name, the TODO list will be displayed. Please enter the completed TODO by the number.

- `--mark || -m`
    - show markdown contents.Please do not enter anything in the argument.

## Markdown
- ` - ` : This character is converted to ` ● ` so you can easily create a list.

-  ` = ` : This character is converted to ` ◎ `, please use it for important items.

- ` ; ` : Please use this character at the end of the line as it is recognized as ` newline `.
