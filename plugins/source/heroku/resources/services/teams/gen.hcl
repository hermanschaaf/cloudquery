//check-for-changes
service          = "heroku"
output_directory = "."
add_generate     = true

resource "github" "" "teams" {
  path = "github.com/heroku/heroku-go/v5.Team"

  ignoreError "IgnoreError" {
    path = "github.com/cloudquery/cloudquery/plugins/source/heroku/client.IgnoreError"
  }
}