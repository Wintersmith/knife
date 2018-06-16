# knife

A Work In Progress, mainly something to write whilst I get to grips with Go.  knife currently provides the following sub-modules:

 - aws - A module providing helper functions for AWS ( currently, very little )
 - files - Helper functions based on various os.path Python functions, for example, OSRename, OSDelete, OSPathExists
 - config - Eventually will provide similar functionality to Python’s configparser module.
 - web - Helper functions when dealing with things on the web.
 
     `GetGoVersion() string`
     
 - rss - Currently doesn't do anything but load OPML files.
 - dateutil - The beginnings of a library to make dealing with dates easier.

    `Now()`
    
    `NextDay( string, SUNDAY...SATURDAY ) time.Time`

    `GregorianEaster( int ) time.Time` 

