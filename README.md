# gotree
Generate a txt filetree from a txt file.

The tool takes a txt-file with contains a simple structure:

```
folder1
    folder2
        file1
        file2
    folder3
```
 
Then it generates a tree structure.

```
folder1
    └── folder2
        └── file1
            ├── file2
            └── folder3
```

