*************
* CARD GAME *
*************

This is a Golang backend to provide APIs required for card games

1. Building the solution:
   cd to projects root directory (card-game)
   Execute "go build . example.com/card-game"
   It will create an exe file in projects root directory

2. Executing the solution:
   From cmd, execute "start card-game.exe"
   It will start the server at port 8080 on your machine

3. Sample cURL Requests:
    a. Create A Deck: curl --location --request GET 'http://localhost:8000/deck/new?shuffled=true&cards=AS,KD,2D'
    b. Open A Deck: curl --location --request GET 'http://localhost:8000/deck/6779c280-87fd-445a-8350-4a7be16b122e'
    c. Draw A Card: curl --location --request GET 'http://localhost:8000/deck/6779c280-87fd-445a-8350-4a7be16b122e/draw?count=1'

4. Execute Test Cases:
   a. cd to card-game/tests
   b. Execute "go test ."
   It will execute all test cases inside tests directory and show out put