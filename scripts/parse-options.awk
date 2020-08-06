#!/usr/bin/awk -f

# Parses the options.txt file and generates structs for each line based on the
# Opt type from /pkg/vimoptions/vimoptions.go

BEGIN {
  FS="'( |\t)+";
} {
  sQuotes = gensub(/"/, "'", "g", $3)
  print "\tOpt{"
  print "\t\t\"Long\": \"" substr($1, 2)"\","
  print "\t\t\"Short\": \"" substr(((NF>2) ? $2 : sQuotes), 2) "\","
  print "\t\t\"Description\": \"" ((NF>2) ? sQuotes : $2) "\","
  print "\t},"
}
