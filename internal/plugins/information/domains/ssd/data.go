// Copyright (C) 2024 Rangertaha
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package ssd

// https://github.com/glaslos/ssdeep

var Data1 string = `
LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.
LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)


       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns

       --color[=WHEN]
              color the output WHEN; more info below

       -d, --directory
              list directories themselves, not their contents

       -D, --dired
              generate output designed for Emacs' dired mode

       -f     list all entries in directory order

       -F, --classify[=WHEN]
              append indicator (one of */=>@|) to entries WHEN

       --file-type
              likewise, except do not append '*'

       --format=WORD
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C

       --full-time
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

`

var Data2 string = `
WHOIS(1)                                          Debian GNU/Linux                                         WHOIS(1)

NAME
       whois - client for the whois directory service

SYNOPSIS
       whois  [ { -h | --host } HOST ] [ { -p | --port } PORT ] [ -abBcdGHIKlLmMrRx ] [ -g SOURCE:FIRST-LAST ] [ -i
       ATTR[,ATTR]... ] [ -s SOURCE[,SOURCE]... ] [ -T TYPE[,TYPE]... ] [ --verbose ] [ --no-recursion ] OBJECT

       whois -q KEYWORD

       whois -t TYPE

       whois -v TYPE

       whois --help

       whois --version

DESCRIPTION
       whois searches for an object in a RFC 3912 database.

       This version of the whois client tries to guess the right server to ask for  the  specified  object.  If  no
       guess  can  be made it will connect to whois.networksolutions.com for NIC handles or whois.arin.net for IPv4
       addresses and network names.

OPTIONS
       -h      HOST, --host=HOST Connect to HOST.

       -H      Do not display the legal disclaimers that some registries like to show you.

       -p      PORT, --port=PORT Connect to PORT.

       -I      First query whois.iana.org and then follow its referral to the whois server authoritative  for  that
               request.  This  works  for IP addresses, AS numbers and domains.  BEWARE: this implies that the IANA
               server will receive your complete query.

       --no-recursion
               Disable recursion from registry to registrar servers.

       --verbose
               Be verbose.

       --help  Display online help.

       --version
               Display the program version.

       Other options are flags understood by whois.ripe.net and some other RIPE-like servers:

       -a      Also search all the mirrored databases.

       -b      Return brief IP address ranges with abuse contact.

       -B      Disable objects filtering.  (Show the e-mail addresses.)

       -c      Return the smallest IP address range with a reference to an irt object.

       -d      Return the reverse DNS delegation object too.

       -g      SOURCE:FIRST-LAST Search updates from SOURCE database between FIRST and LAST update  serial  number.
               It is useful to obtain Near Real Time Mirroring stream.
 Manual page whois(1) line 1 (press h for help or q to quit)
LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)

LS(1)                                              User Commands                                              LS(1)

NAME
       ls - list directory contents

SYNOPSIS
       ls [OPTION]... [FILE]...

DESCRIPTION
       List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of
       -cftuvSUX nor --sort is specified.

       Mandatory arguments to long options are mandatory for short options too.

       -a, --all
              do not ignore entries starting with .

       -A, --almost-all
              do not list implied . and ..

       --author
              with -l, print the author of each file

       -b, --escape
              print C-style escapes for nongraphic characters

       --block-size=SIZE
              with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below

       -B, --ignore-backups
              do not list implied entries ending with ~

       -c     with  -lt:  sort  by, and show, ctime (time of last change of file status information); with -l: show
              ctime and sort by name; otherwise: sort by ctime, newest first

       -C     list entries by columns                                                                               
                                                                                                                    
       --color[=WHEN]                                                                                               
              color the output WHEN; more info below                                                                
                                                                                                                    
       -d, --directory                                                                                              
              list directories themselves, not their contents                                                       
                                                                                                                    
       -D, --dired                                                                                                  
              generate output designed for Emacs' dired mode                                                        
                                                                                                                    
       -f     list all entries in directory order                                                                   
                                                                                                                    
       -F, --classify[=WHEN]                                                                                        
              append indicator (one of */=>@|) to entries WHEN                                                      
                                                                                                                    
       --file-type                                                                                                  
              likewise, except do not append '*'                                                                    
                                                                                                                    
       --format=WORD                                                                                                
              across -x, commas -m, horizontal -x, long -l, single-column -1, verbose -l, vertical -C               
                                                                                                                    
       --full-time                                                                                                  
              like -l --time-style=full-iso

       -g     like -l, but do not list owner

       --group-directories-first
              group directories before files; can be augmented with a --sort option, but  any  use  of  --sort=none
              (-U) disables grouping

       -G, --no-group
              in a long listing, don't print group names

       -h, --human-readable
              with -l and -s, print sizes like 1K 234M 2G etc.

       --si   likewise, but use powers of 1000 not 1024

       -H, --dereference-command-line
              follow symbolic links listed on the command line

       --dereference-command-line-symlink-to-dir
              follow each command line symbolic link that points to a directory

       --hide=PATTERN
              do not list implied entries matching shell PATTERN (overridden by -a or -A)

       --hyperlink[=WHEN]
              hyperlink file names WHEN

       --indicator-style=WORD
              append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type),
              classify (-F)

       -i, --inode
              print the index number of each file

       -I, --ignore=PATTERN
              do not list implied entries matching shell PATTERN

       -k, --kibibytes
              default to 1024-byte blocks for file system usage; used only with -s and per directory totals

       -l     use a long listing format

       -L, --dereference
              when  showing file information for a symbolic link, show information for the file the link references
              rather than for the link itself

       -m     fill width with a comma separated list of entries

       -n, --numeric-uid-gid
              like -l, but list numeric user and group IDs

       -N, --literal
              print entry names without quoting

       -o     like -l, but do not list group information

       -p, --indicator-style=slash
              append / indicator to directories

       -q, --hide-control-chars
              print ? instead of nongraphic characters

       --show-control-chars
              show nongraphic characters as-is (the default, unless program is 'ls' and output is a terminal)

       -Q, --quote-name
              enclose entry names in double quotes

       --quoting-style=WORD
              use quoting style WORD for entry names: literal, locale, shell, shell-always, shell-escape, shell-es‐
              cape-always, c, escape (overrides QUOTING_STYLE environment variable)

       -r, --reverse
              reverse order while sorting

       -R, --recursive
              list subdirectories recursively

       -s, --size
              print the allocated size of each file, in blocks

       -S     sort by file size, largest first

       --sort=WORD
              sort by WORD instead of name: none (-U), size (-S), time (-t), version (-v), extension (-X), width

       --time=WORD
              select which timestamp used to display or sort; access time (-u): atime, access, use; metadata change
              time (-c): ctime, status; modified time (default): mtime, modification; birth time: birth, creation;

              with -l, WORD determines which time to show; with --sort=time, sort by WORD (newest first)

       --time-style=TIME_STYLE
              time/date format with -l; see TIME_STYLE below

       -t     sort by time, newest first; see --time

       -T, --tabsize=COLS
              assume tab stops at each COLS instead of 8

       -u     with -lt: sort by, and show, access time; with -l: show access time and sort by name; otherwise: sort
              by access time, newest first

       -U     do not sort; list entries in directory order

       -v     natural sort of (version) numbers within text

       -w, --width=COLS
              set output width to COLS.  0 means no limit

       -x     list entries by lines instead of by columns

       -X     sort alphabetically by entry extension

       -Z, --context
              print any security context of each file

       --zero end each output line with NUL, not newline

       -1     list one file per line

       --help display this help and exit

       --version
              output version information and exit

       The SIZE argument is an integer and optional unit (example: 10K is 10*1024).  Units are  K,M,G,T,P,E,Z,Y,R,Q
       (powers of 1024) or KB,MB,... (powers of 1000).  Binary prefixes can be used, too: KiB=K, MiB=M, and so on.

       The  TIME_STYLE  argument can be full-iso, long-iso, iso, locale, or +FORMAT.  FORMAT is interpreted like in
       date(1).  If FORMAT is FORMAT1<newline>FORMAT2, then FORMAT1 applies to non-recent files and FORMAT2 to  re‐
       cent  files.   TIME_STYLE  prefixed  with  'posix-'  takes  effect  only outside the POSIX locale.  Also the
       TIME_STYLE environment variable sets the default style to use.

       The WHEN argument defaults to 'always' and can also be 'auto' or 'never'.

       Using color  to  distinguish  file  types  is  disabled  both  by  default  and  with  --color=never.   With
       --color=auto,  ls emits color codes only when standard output is connected to a terminal.  The LS_COLORS en‐
       vironment variable can change the settings.  Use the dircolors(1) command to set it.

   Exit status:
       0      if OK,

       1      if minor problems (e.g., cannot access subdirectory),

       2      if serious trouble (e.g., cannot access command-line argument).

AUTHOR
       Written by Richard M. Stallman and David MacKenzie.

REPORTING BUGS
       GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
       Report any translation bugs to <https://translationproject.org/team/>

COPYRIGHT
       Copyright  ©  2023  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version   3   or   later
       <https://gnu.org/licenses/gpl.html>.
       This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent per‐
       mitted by law.

SEE ALSO
       dircolors(1)

       Full documentation <https://www.gnu.org/software/coreutils/ls>
       or available locally via: info '(coreutils) ls invocation'

GNU coreutils 9.4                                    March 2024                                               LS(1)


`
