Example GraphQL Query:

query {
getLeapYears(input: {startYear: 2023, endYear: 2043, yearDuration: 365.26}) {
leapYear
}
}

Result:

{
"data": {
"getLeapYears": {
"leapYear": [
2024,
2028,
2032,
2036,
2040
]
}
}
}