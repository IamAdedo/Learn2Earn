## Instructions

Create the files and directories so that when you use the command ls below, the output will look like this :

$ TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'

dr-------x 1986-01-05 00:00 0
-r------w- 1986-11-13 00:01 1
-rw----r-- 1988-03-05 00:10 2
lrwxrwxrwx 1990-02-16 00:11 3 -> 0
-r-x--x--- 1990-10-07 01:00 4
-r--rw---- 1990-11-07 01:01 5
-r--rw---- 1991-02-08 01:10 6
-r-x--x--- 1991-03-08 01:11 7
-rw----r-- 1994-05-20 10:00 8
-r------w- 1994-06-10 10:01 9
dr-------x 1995-04-10 10:10 A
$


Note: In order to run the command given in the example above, on Mac operating systems, it is necessary to install the GNU Core Utilities. After installing the command will be slightly different like so:

$ TZ=utc gls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
The mode from line 4 may look different, for example like:

lrwxr-xr-x 1990-02-16 00:11 3 -> 0
Once it is done, use the command below to create the file done.tar to be submitted.

$ tar -cf done.tar *
$ ls
0  1  2  3  4  5  6  7  8  9  A  done.tar
$
Only done.tar should be submitted.

Test panel
You can test your project here!

Hey ☺

Here will be displayed your test output.
But first, complete the exercise and submit it on Gitea!








# File and Directory Creation Challenge

This project requires creating a specific set of files and directories with precise permissions, timestamps, and symlinks. When configured correctly, listing the directory contents with the specified `ls` command will produce the exact target output.

## Target Output

Your directory structure and file attributes must match this exact layout:

```text
dr-------x 1986-01-05 00:00 0
-r------w- 1986-11-13 00:01 1
-rw----r-- 1988-03-05 00:10 2
lrwxrwxrwx 1990-02-16 00:11 3 -> 0
-r-x--x--- 1990-10-07 01:00 4
-r--rw---- 1990-11-07 01:01 5
-r--rw---- 1991-02-08 01:10 6
-r-x--x--- 1991-03-08 01:11 7
-rw----r-- 1994-05-20 10:00 8
-r------w- 1994-06-10 10:01 9
dr-------x 1995-04-10 10:10 A
```

> **Note on Symlinks:** The permission mode for line `3` (the symlink) may vary depending on your operating system (e.g., `lrwxr-xr-x`). This is acceptable.

## Verification Commands

To test and verify your setup, run the appropriate command for your operating system.

### Linux
```bash
TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
```

### macOS
macOS users must first install GNU Core Utilities (`brew install coreutils`) to use `gls`:
```bash
TZ=utc gls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
```

## Submission Instructions

Once your directory matches the target output exactly, bundle your files into a tar archive for submission.

1. **Create the archive:**
   ```bash
   tar -cf done.tar *
   ```

2. **Verify your directory contents:**
   ```bash
   ls
   ```
   *Expected output:*
   ```text
   0 1 2 3 4 5 6 7 8 9 A done.tar
   ```

3. **Submit:** Upload **only** the `done.tar` file to Gitea.





---

## Instructions
Create the files and directories so that when you use the command ls below, the output will look like this :

$ TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
dr-------x 1986-01-05 00:00 0
-r------w- 1986-11-13 00:01 1
-rw----r-- 1988-03-05 00:10 2
lrwxrwxrwx 1990-02-16 00:11 3 -> 0
-r-x--x--- 1990-10-07 01:00 4
-r--rw---- 1990-11-07 01:01 5
-r--rw---- 1991-02-08 01:10 6
-r-x--x--- 1991-03-08 01:11 7
-rw----r-- 1994-05-20 10:00 8
-r------w- 1994-06-10 10:01 9
dr-------x 1995-04-10 10:10 A
$

Note: In order to run the command given in the example above, on Mac operating systems, it is necessary to install the GNU Core Utilities. After installing the command will be slightly different like so:
$ TZ=utc gls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'
The mode from line 4 may look different, for example like:
lrwxr-xr-x 1990-02-16 00:11 3 -> 0

Once it is done, use the command below to create the file done.tar to be submitted.

$ tar -cf done.tar *
$ ls0 1 2 3 4 5 6 7 8 9 A done.tar
$

Only done.tar should be submitted.
------------------------------
## Test panel
You can test your project here!

Hey ☺ Here will be displayed your test output. But first, complete the exercise and submit it on Gitea!

------------------------------
If you'd like, I can write a shell script to automate creating these exact files, permissions, and timestamps. Would you like me to provide the commands or help with a specific part of the setup?

