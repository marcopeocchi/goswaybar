package metrics

//awk -F"[][]" '/Left:/ { print $2 }' <(amixer sget Master)