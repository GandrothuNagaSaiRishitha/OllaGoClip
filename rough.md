New things i Learnt while building this: 
--> toml package
what this package does is that it loooks into the system configuration files and without needing to refactor our code according to different computers, its enables settings directly into a struct.

When your application starts up, it needs to read your settings. The package takes a human-readable text file (like config.toml) and automatically maps those settings directly into a structured format (a struct) in your Go code so your application can use them.

If your application has a settings menu or generates configuration files programmatically, go-toml lets you take variables and data structures inside your Go code and save them back into a neatly formatted .toml text file.


