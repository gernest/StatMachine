StatMachine
===========

Go library for analyzing football results. This library is used to generate everything from simple stats (league tables) to
more complex stats (#number of first halfs scored in in a row). 

# Required Data #
The library requires that results are provided as a input from a 3d party library and the following data for each result:
- Goals scored.
- Goals scored at half time.
- Opponents goals scored.
- Opponents goals socred at half time.
- Home or away game.
- Season identifier 
- date
- round

# League Tables #
When you have a collection of results you can easily generate a league tables.

```go
league := statmachine.NewLeague("English PL")
league = getResults(league)//get results from an external provider
table := statmachine.GenerateLeagueTable(league, 2014)
```
It is also easy to generate league tables for results against teams in the top half of the league
and bottom half of the league.

```go
league := statmachine.NewLeague("English PL")
league = getResults(league)//get results from an external provider
againstTopHalfTable := statmachine.GenerateLeagueTableAgainstTopHalf(league, 2014)
againstBottomHalfTable := statmachine.GenerateLeagueTableBottomTopHalf(league, 2014)
```

All these methods return a LeagueTable struct that is a collection of league positions which each has the following properties:
- Place
- Team name
- Points
- Wins
- Draws
- Losses
- Goals for
- Goals against

# Example of usage #
Lets say that you have all the results for a given season.

```go
league := statmachine.NewLeague("English PL")
league = getResultsForCurrentSeason(league)//get results from an external provider
```

Now lets say that you want to find the most games won in a row for each of the teams in the league:

```go
for _,team := range league.Teams {
  gamesWonInARow[team.Name] = statmachine.gamesWonInARow(team.Results)
}
```

# Other Supported Stats #

The following stats can be calculated:

- Total number of games won/drawn/lost.
- Games won/drawn/lost in a row.
- Games scored in/conceded in a row.
- Total number of goals scored an conceded.
- Find the last game won/drawn/lost.
- Total number of games scored in/clean sheets.
- Most games won/lost in a row.
- Most games without a win/loss in a row.
- Most halves scored in in a row.
- Most first halves/second halves scored in in a row.


# Filtering Results #
The library supports filtering a set of results by the following properties:
- Home/away results
- leading/trailing at half time
- By season
- By opponent
- By rounds

If results contains all results for a given team, the following snippet
would give you their results in the first 3 rounds of each season:

```go
filteredResults := statmachine.ResultsByRounds(results, []int{1,2, 3})
```
Combining these filters can be quite powerful. If results
contains all results for a given team, you could get all results against a
certain team at home with the following snippet:

```go
atHome := statmachine.HomeResults(results)
atHomeVsTeamA := statmachine.ResultsByOpponent(atHome, opponentId)
```

# Grouping Results #
Its possible to group a set of results using the GroupBy method:

```go
func GroupBy(res []Result, f func(Result) string) map[string][]Result
```

An example of how you would group a set of results by season:

```go
groupedResults := statmachine.GroupBy(results, func(r statmachine.Result)string{return strconv.Itoa(r.SeasonId())})
```
