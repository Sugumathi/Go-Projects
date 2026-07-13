High Level Architecture
                 +----------------+
                 |    main.go     |
                 +-------+--------+
                         |
                         |
                Parse CLI Flags
                         |
                         v
               +-----------------+
               |      config     |
               +-----------------+
                         |
                         |
               Build Search Config
                         |
                         |
                         v
              +-------------------+
              |    SearchEngine   |
              +---------+---------+
                        |
        -----------------------------------
        |                |                |
        |                |                |
 File Walker       Worker Pool       stdin
        |                |
        |                |
        +--------+-------+
                 |
          Read Files
                 |
                 v
          Matcher Interface
                 |
        --------------------
        |                  |
 Literal Matcher     Regex Matcher
        |
        |
   Match Results
        |
        |
 Output Formatter
        |
        |
 Terminal

 // Main function
    Input: substring, filepath