# Undercover
##Undercover is a multiplayer game. The game procedure is as following:

1. A judge will randomly distribute 2 words to all players without them knowing each other's word. Eg. there are 5 players, the judge give out 4 "Holmes" and 1 "Conan" to them randomly. 

2. Then the one holding "Conan" will become the "Undercover", while the others are normal people. The 2 words should be different but have similarities (detailed rules on words will be introduced later).

3. Each player need to describe the word they get, it is better not too obvious for the undercover to find out the word, but also should give enough hints for teammates (to let them know you belong to same group).

4. After a round of description, everyone need to vote the most suspicious person they think, the one gets highest votes will be out.

5. After a few rounds (depends on different game configuration), if the undercover is still alive then undercover wins, otherwise normal people wins

###Different game configuration based on players number.

Minimum players number will be 3 (including at least 1 undercover). If any undercover survived to the last round (there are only 3 people left), then undercover wins; otherwise normal people wins.

Recommended:

| Total Players | Normal People | Undercover |
|---------------|---------------|------------|
|       5       |       4       |      1     |
|       9       |       6       |      2     |
|       10      |       7       |      3     |
|       14      |       10      |      4     |

###Rules on Words:
1. Given words should not exceed players' knowledge
2. The 2 words cannot be including each other, eg, "furniture" and "table"
3. The 2 words should be same part of speech, eg. both noun or both verb
4. The 2 words cannot be synonyms of each other, eg. "black" and "dark", or "beautiful" and "pretty"
5. When player describe his given word, it must be correct and related, eg. when the 2 words are "Conan" and "Holmes" and one player's word is "Holmes", he cannot use "he is handsome" to describe him as it's not so related.

##Get Started
start the server:
go run main.go

web-based clients:
localhost:8000/admin to access admin interface: game configuration and control
localhost:8000/player to access player interface: join game


