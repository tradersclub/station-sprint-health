# Challenge 2

Implement a function that receive a word from the channel called `chanNewWord`, check if this word doesn't have any repeated letter, and store this info in slice called `wordsWithAlphabet`.

You can insert a new word just sending a request like this:

``[GET] http://localhost:8080/set?word=myword``

And you can check all the slice doing a request like this:

``[GET] http://localhost:8080/get``

Ps.: You can use just the browser to run this requests