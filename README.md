# Command-Line-Interface

I built a Command Line Interface Application in Golang with Cobra to perform various commands such as creating a new folder, listing all files in a directory, deleting files, 
initializing a git repository in a specified folder, add, subtracting, dividing, multiplying integers and floats, concatenating strings together, counting the amount of indexes
in a string, and parsing a string 

How to use the Command Line Interface Application
Go into the file directory and type the file name in the terminal "CLI-Tool". This will give access to all the available commands

The commands and how to use them:
- To add a new folder in the directory run "CLI-Tool createFolder -n folderName" Where folderName is the name of the folder input by the user
- To delete a file or a folder run "CLI-Tool deleteFile -r fileName" Where fileName is the name of the folder or file to be deleted
- To list all the files in a directory run CLI-Tool listFiles"
- To initialize a git repository inside a given folder run "CLI-Tool gitInit -d folderName" Where folderName is the name of the folder to initialize a git repository in.
- To split or parse a string with a delimiter run "CLI-Tool parse string delimiter" Where string is the input string and delimiter in the input delimiter provided by the user. For
  example "CLI-Tool parse My,name,is,maaz ," will output "My name is maaz"
- To count the number of indexes in a given string run "CLI-Tool String count string" Where string is the input string the user will type out
- To concatenate strings together separated by a space run "CLI-Tool ConcatenateStrings string" Where string is the input string the user will type out. This will 
  get rid of any space in between the string provided.
- To add integers run "CLI-Tool add x x" or to add floats run "CLI-Tool add x x -f". Where x is just any integer or float number.
- To subtract integers run "CLI-Tool subtract x x" or to subtract floats run "CLI-Tool subtract x x -f". Where x is just any integer or float number.
- To divide integers run "CLI-Tool division x x" or to divide floats run "CLI-Tool division x x -f". Where x is just any integer or float number.
- To multiply integers run "CLI-Tool mulitply x x" or to multiply floats run "CLI-Tool multiply x x -f". Where x is just any integer or float number.
