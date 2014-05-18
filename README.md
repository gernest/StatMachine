StatMachine
===========

Go library for analyzing football results.

# Required Data #
The library requires the following data for each result:
- Goals scored.
- Goals scored at half time.
- Opponents goals scored.
- Opponents goals socred at half time.
- Home or away game.
- Season identifier 
- date
- round

# Supported Stats #

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

# League Tables #
The library supports creating league tables on a given set of results. 

```go
league := statmachine.NewLeague("English PL")
//... Add results to league
table := statmachine.GenerateLeagueTable(league, 2014)
```

A league table is a collection of league positions which each has the following properties:
- Place
- Team name
- Points
- Wins
- Draws
- Losses
- Goals for
- Goals against

# Grouping Results #
Its possible to group a set of results using the GroupBy method:

```go
func GroupBy(res []Result, f func(Result) string) map[string][]Result
```

An example of how you would group a set of results by season:

```go
groupedResults := statmachine.GroupBy(results, func(r statmachine.Result)string{return strconv.Itoa(r.SeasonId())})
```