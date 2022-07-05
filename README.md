bobbymrdr
=========

Bobby Might Remove the DiRectory.

> NOTE: This is a very dangerous program. ⚠️

Don't run it if you don't know what you are doing. You have been warned.

Now, here is what it does.

Removes a directory from your filesystem after a specified duration. It may create the directory if one doesn't exist.
The expressions for the time are the ones supported by Go's [time.ParseDuration](https://pkg.go.dev/time#ParseDuration) such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". 

## Usage

Again, don't use this if you don't know what you are doing or aren't sure. There are no safeguards. I am not responsible for any data loss or corruption. ⚠️

```sh
$ bobbymrdr -d "new-project-directory" -r "5m" 
```


## Why?

I wanted to write some code that does something dangerous. 😎

# LICENSE

                DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                        Version 2, December 2004
    
     Copyright (C) 2018 Zikani Nyirenda Mwase
    
     Everyone is permitted to copy and distribute verbatim or modified
     copies of this code, and changing it is allowed as long
     as the name is changed.
    
                DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
       TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
    
      0. You just DO WHAT THE FUCK YOU WANT TO.

---

(c) Zikani Nyirenda Mwase